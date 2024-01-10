package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// check if there is a command line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command line argument")
		return
	}

	// get command line argument
	file := os.Args[1]
	// get PATH environment variable
	path := os.Getenv("PATH")

	for _, directory := range strings.Split(path, ":") {
		fullPath := directory + "/" + file
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			//Is is a regural file
			if fileInfo.Mode().IsRegular() {
				// Is it executable
				if fileInfo.Mode()&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}

}
