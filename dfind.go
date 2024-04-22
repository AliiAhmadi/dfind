package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	info fs.FileInfo
	path string
}

func main() {
	var path string
	current, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&path, "path", current, "the root path that want to start scan and find duplicated files")
	flag.Parse()

	files := make([]File, 0)

	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, File{
				info: info,
				path: path,
			})
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
}
