package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	var dir string

	flag.StringVar(&dir, "dir", "", "directory that has the files needing organization")
	flag.Parse()

	if dir == "" {
		fmt.Println("Specify a directory")
		return
	}

	fmt.Println("Directory:", dir)

	err := organizeFiles(dir)
	if err != nil {
		log.Fatal(err)
	}
}
