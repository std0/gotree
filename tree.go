package main

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"io/ioutil"
)

const (
	emptySpace   = "\t"
	middleItem   = "├───"
	continueLine = "│\t"
	lastItem     = "└───"
)

func getFormattedFileInfo(file os.FileInfo) string {
	var size string
	if file.Size() == 0 {
		size = "empty"
	} else {
		size = fmt.Sprint(file.Size()) + "b"
	}
	return fmt.Sprint(file.Name() + " (" + size + ")")
}

func tree(out io.Writer, path, indent string, withFiles bool) error {
	file, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !file.IsDir() {
		if withFiles {
			fmt.Fprintln(out, getFormattedFileInfo(file))
		}
		return nil
	}
	fmt.Fprintln(out, file.Name())

	subFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if !withFiles {
		dirs := make([]os.FileInfo, 0)
		for _, subFile := range subFiles {
			if subFile.IsDir() {
				dirs = append(dirs, subFile)
			}
		}
		subFiles = dirs
	}

	for i, subFile := range subFiles {
		newIndent := continueLine
		if i == len(subFiles) - 1 {
			fmt.Fprintf(out, indent + lastItem)
			newIndent = emptySpace
		} else {
			fmt.Fprintf(out, indent + middleItem)
		}

		err := tree(out, path + "/" + subFile.Name(), indent + newIndent, withFiles)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run tree.go PATH [-f]")
		os.Exit(1)
	}
	out := os.Stdout
	path, err := filepath.Abs(os.Args[1])
	withFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err = tree(out, path, "", withFiles)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
