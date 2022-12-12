package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name            string
	filesSize       int
	subDirectories  []*Directory
	parentDirectory *Directory
}

// calculateFilesSize returns the total size of files at the root of the directory (excluding subdirectories)
func calculateFilesSize(content []string) int {
	totalSize := 0
	for _, line := range content {
		size, err := strconv.Atoi(line[:strings.Index(line, " ")])
		if err == nil {
			totalSize += size
		}
	}
	return totalSize
}

func parseSubDirectories(content []string, parent *Directory) []*Directory {
	subDirs := []*Directory(nil)
	for _, line := range content {
		if strings.HasPrefix(line, "dir") {
			subDirs = append(subDirs, &Directory{
				name:            line[4:],
				parentDirectory: parent,
			})
		}
	}
	return subDirs
}

func buildFileTree(terminalOutput []string) *Directory {
	fileTree := Directory{
		name: "/",
	}
	dir := &fileTree

	for i, line := range terminalOutput {

		if strings.HasPrefix(line, "$ cd ") {
			directoryToReach := line[5:]
			if directoryToReach == ".." {
				dir = dir.parentDirectory
			} else {
				for _, subDir := range dir.subDirectories {
					if subDir.name == line[5:] {
						dir = subDir
					}
				}
			}

		} else if line == "$ ls" {
			nextCommandIndex := i + 1
			for ; nextCommandIndex < len(terminalOutput); nextCommandIndex++ {
				if strings.HasPrefix(terminalOutput[nextCommandIndex], "$") {
					break
				}
			}
			dirName := terminalOutput[i-1][5:]
			lsOutput := terminalOutput[i+1 : nextCommandIndex]
			for _, subDir := range dir.subDirectories {
				if subDir.name == dirName {
					dir = subDir
				}
			}
			dir.filesSize = calculateFilesSize(lsOutput)
			dir.subDirectories = parseSubDirectories(lsOutput, dir)
		}
	}

	return &fileTree
}

func addSubDirectoriesSizes(directories []*Directory) int {
	sum := 0
	for _, dir := range directories {
		if len(dir.subDirectories) != 0 {
			dir.filesSize += addSubDirectoriesSizes(dir.subDirectories)
		} else {
			sum += dir.filesSize
		}
	}
	return sum
}

func calculateTotalFilesSize(data string) int {
	fileTree := buildFileTree(strings.Split(data, "\n"))
	sum := 0 //countStuff(fileTree, 0, 0)
	for _, dir := range fileTree.subDirectories {
		if len(dir.subDirectories) != 0 {
			dir.filesSize += addSubDirectoriesSizes(dir.subDirectories)
		}
		if dir.filesSize < 100000 {
			sum += dir.filesSize
		}
	}
	return sum
}

func main() {
	buffer, err := os.ReadFile("commands_list.txt")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	size := calculateTotalFilesSize(string(buffer))

	fmt.Printf("Total size of directories over 100.000: %d", size)
}
