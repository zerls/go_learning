package main

import (
	"./github"
	"fmt"
	"log"
)

func main() {

	ss := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(ss)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.Items)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}
