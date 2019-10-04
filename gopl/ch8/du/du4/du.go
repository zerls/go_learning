package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", true, "show verbose progress messages")

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)
var done = make(chan bool)

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{`C:\`}
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	ok1 := true
	for ok1 {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				ok1 = false
				break
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				//Do nothing.
			}
			return
			//panic("bye")
		}

	}

	fmt.Printf("final totals：")
	printDiskUsage(nfiles, nbytes) // final totals
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	if cancelled() {
		return
	}

	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() {
		<-sema // release token
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
