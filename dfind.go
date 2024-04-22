package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	current, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&path, "path", current, "the root path that want to start scan and find duplicated files")
	flag.Parse()

	fmt.Println(path)
}
