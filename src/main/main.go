package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func printTree(out io.Writer, folder string, verbose bool, offset int, delimiters []string) {
	file, _ := os.Open(folder)
	filesInfo, _ := file.Readdir(0)

	printLevel := func(n int, isLast bool) {
		for i := 0; i < n; i++ {
			fmt.Fprintf(out, "%s\t", delimiters[i])
		}
		if isLast {
			fmt.Fprint(out, "└───")
		} else {
			fmt.Fprint(out, "├───")
		}
	}

	if !verbose {
		filteredFilesInfo := filesInfo[:0]
		for _, fileInfo := range filesInfo {
			if fileInfo.IsDir() {
				filteredFilesInfo = append(filteredFilesInfo, fileInfo)
			}
		}
		filesInfo = filteredFilesInfo[:]
	}
	sort.Slice(filesInfo, func(i, j int) bool { return filesInfo[i].Name() < filesInfo[j].Name() })
	lastFileIndex := len(filesInfo) - 1
	var delimiter string

	for i, fileInfo := range filesInfo {
		printLevel(offset, i == lastFileIndex)
		if !fileInfo.IsDir() {
			if fileInfo.Size() == 0 {
				fmt.Fprintf(out, "%s (empty)\n", fileInfo.Name())
			} else {
				fmt.Fprintf(out, "%s (%db)\n", fileInfo.Name(), fileInfo.Size())
			}
		} else {
			fmt.Fprintln(out, fileInfo.Name())
			if i == lastFileIndex {
				delimiter = ""
			} else {
				delimiter = "│"
			}
			offset++
			delimiters = append(delimiters, delimiter)
			printTree(out, filepath.Join(folder, fileInfo.Name()), verbose, offset, delimiters)
			delimiters = delimiters[:len(delimiters)-1]
			offset--
		}
	}
}

func dirTree(out io.Writer, folder string, verbose bool) error {
	printTree(out, folder, verbose, 0, []string{})
	return nil
}

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
