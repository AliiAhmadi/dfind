package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type File struct {
	info fs.FileInfo
	path string
	hash string
}

func ForceExit(f *os.File, msg string, status int) {
	fmt.Fprint(f, msg)
	os.Exit(status)
}

func Unique[T comparable](values []T) []T {
	res := make([]T, 0, len(values))
	seen := make(map[T]struct{}, len(values))

	for _, item := range values {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		res = append(res, item)
	}

	return res
}

func GetMD5(location string) (string, error) {
	f, err := os.Open(location)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	dirs := make([]string, 0)
	files := make([]File, 0)

	args := Unique(os.Args[1:])

	for _, dir := range args {
		_, err := os.Open(dir)
		if err != nil {
			ForceExit(os.Stderr, fmt.Sprintf("%v does not exist\n", dir), 1)
		}

		dirs = append(dirs, dir)
	}

	if len(args) == 0 {
		wd, _ := os.Getwd()
		dirs = append(dirs, wd)
	}

	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if !info.IsDir() {
				h, err := GetMD5(path)
				if err != nil {
					ForceExit(os.Stderr, err.Error(), 1)
				}

				files = append(files, File{
					info: info,
					path: path,
					hash: h,
				})
			}
			return nil
		})

		if err != nil {
			ForceExit(os.Stderr, fmt.Sprintf("%v\n", err.Error()), 2)
		}
	}

	fmt.Println(files)
}
