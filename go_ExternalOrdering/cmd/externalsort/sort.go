package main

import (
	"../../pipline"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	p := createNetworkPipeline("small.in",
		512, 4)

	//time.Sleep(time.Hour)
	writeToFile(p, "small.out")
	printFile("small.out")
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipline.WriterSink(writer, p)
}

func createPipeline(
	filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipline.Init()
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipline.ReaderSource(
			bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults,
			pipline.InMenSort(source))

	}

	return pipline.MergeN(sortResults...)
}

func createNetworkPipeline(
	filename string,
	fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	pipline.Init()
	sortAddr := []string{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipline.ReaderSource(
			bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		//推给网络服务器
		pipline.NetworkSink(addr, pipline.InMenSort(source))
		sortAddr = append(sortAddr, addr)

	}
	//从网络服务器获取
	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults,
			pipline.NetworkSource(addr))
	}

	return pipline.MergeN(sortResults...)
}
