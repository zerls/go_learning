package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		//1.8
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//	b,err :=ioutil.ReadAll(resp.Body)
		//	resp.Body.Close()
		//	if err !=nil{
		//		fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
		//		os.Exit(1)
		//	}
		//	fmt.Printf("%s",b)
		//}

		//1.7
		//if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		//	fmt.Fprintf(os.Stderr, "ioError: %v\n", err)
		//}

		//1.9
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
	}
}
