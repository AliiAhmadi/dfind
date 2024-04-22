package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var path string
	current, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&path, "path", current, "the root path that want to start scan and find duplicated files")
	flag.Parse()

	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Printf("%v %v\n", path, info.Size())
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
