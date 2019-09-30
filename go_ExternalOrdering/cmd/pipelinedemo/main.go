package main

import (
	"../../pipline"
	"bufio"
	"fmt"
	"os"
)

func main() {
	const (
		filename = "small.in"
		n        = 64
	)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	defer file.Close()
	reader := bufio.NewReader(file)
	p = pipline.ReaderSource(reader, -1)

	count := 0
	for v := range p {
		count++
		fmt.Println(v)
		if count >= 100 {
			break
		}

	}
	//mergeDemo()

}

func mergeDemo() {
	p := pipline.Merge(
		pipline.InMenSort(
			pipline.ArrySource(3, 2, 6, 7, 4)),
		pipline.InMenSort(
			pipline.ArrySource(7, 4, 0, 3, 2, 13, 8)))

	for v := range p {
		fmt.Println(v)
	}
}
