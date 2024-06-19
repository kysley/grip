package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := os.Args[1]

	match := flag.String("match", "*", "Glob pattern to match files")
	out := flag.String("out", "./grip_out.txt", "Path to output file")
	headFile := flag.String("head-file", "", "Path to file containing prompt to add to start of output (optional)")
	head := flag.String("head", "", "Prompt to add to the start of output (optional)")

	flag.CommandLine.Parse(os.Args[2:])

	// Ensure a directory argument is provided
	if dir == "" {
		fmt.Println("You must specify a directory.")
		os.Exit(1)
	}

	// Ensure --head and --head-file are not both set
	if *head != "" && *headFile != "" {
		fmt.Println("You cannot specify both --head and --head-file.")
		os.Exit(1)
	}

	// Read header content
	var headerContent string
	if *head != "" {
		headerContent = *head
	} else if *headFile != "" {
		data, err := ioutil.ReadFile(*headFile)
		if err != nil {
			fmt.Println("Error reading head file:", err)
			os.Exit(1)
		}
		headerContent = string(data)
	}
	print(*match)

	// Collect contents
	var contents []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && matchFile(path, *match) {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			relativePath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			contents = append(contents, "// "+relativePath+"\n"+string(data))
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the directory:", err)
		os.Exit(1)
	}

	// Write to output file
	err = ioutil.WriteFile(*out, []byte(headerContent+"\n"+strings.Join(contents, "\n\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}
}

// matchFile checks if the file matches the given pattern
func matchFile(path, pattern string) bool {
	matched, err := filepath.Match(pattern, filepath.Base(path))
	if err != nil {
		return false
	}
	return matched
}
