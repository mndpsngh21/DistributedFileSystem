package main

import (
	"distributedFileSystem/filejoiner"
	"distributedFileSystem/filesplitter"
	"fmt"
)

func main() {
	fmt.Println("Starting system")
	two_KB := 1024
	filename := "sample.pdf"  // Replace with your file name
	chunkSize := two_KB * 100 // Chunk size in bytes
	destinationDir := "splitfiles"
	//split file
	err := filesplitter.SplitFile(filename, destinationDir, chunkSize)
	if err != nil {
		fmt.Errorf("Failed to split file %v", filename)
	}

	chunkFolder := destinationDir
	mergedFile := "merged.pdf"
	// merge same file
	err = filejoiner.JoinFiles(chunkFolder, mergedFile)
	if err != nil {
		fmt.Errorf("Failed to merge file %v", filename)
	}
}
