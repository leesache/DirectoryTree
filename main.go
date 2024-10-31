package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, op bool) error {
	op = false
	out = nil
	return dirTreeWithPrefix(path, "")
}

func dirTreeWithPrefix(path string, prefix string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, file := range files {
		isLast := i == len(files)-1

		if isLast {
			fmt.Print(prefix + "└───")
		} else {
			fmt.Print(prefix + "├───")
		}

		fmt.Print(file.Name())

		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				return err
			}
			if info.Size() == 0 {
				fmt.Print(" (empty)")
			} else {
				fmt.Printf(" (%db)", info.Size())
			}
		}
		fmt.Println()

		if file.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "\t"
			} else {
				newPrefix += "│\t"
			}
			dirTreeWithPrefix(filepath.Join(path, file.Name()), newPrefix)
		}
	}
	return nil
}
