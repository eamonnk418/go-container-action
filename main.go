package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func listFiles(root string) ([]string, error) {
	var fileList []string

	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err // Return the error if encountered
		}
		if !info.IsDir() {
			fileList = append(fileList, path) // Append the file path to the slice
		}
		return nil
	}); err != nil {
		return nil, err // Return the error if there was an issue with Walk
	}

	return fileList, nil
}


func main() {
	files, err := listFiles(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("%s\n", file)
	}
}
