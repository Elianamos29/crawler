package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		if len(args) == 1 {
			fmt.Println("no website provided")
			os.Exit(1)
		} else {
			fmt.Println("too many arguments provided")
			os.Exit(1)
		}
	}

	base_url := args[1]

	fmt.Printf("starting crawl of: %s\n", base_url)
}