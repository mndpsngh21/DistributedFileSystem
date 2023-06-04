package filesplitter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SplitFile splits a file into smaller chunks.
func SplitFile(filename string, destinationDir string, chunkSize int) error {

	// Create the destination directory if it doesn't exist
	err := os.MkdirAll(destinationDir, 0755)
	if err != nil {
		fmt.Errorf("Error creating destination directory:", err)
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	chunkNum := 0

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		chunkFilename := fmt.Sprintf("%s.%d", filename, chunkNum)
		chunkPath := filepath.Join(destinationDir, chunkFilename)
		chunkFile, err := os.Create(chunkPath)
		if err != nil {
			return err
		}

		_, err = chunkFile.Write(buffer[:n])
		if err != nil {
			chunkFile.Close()
			return err
		}

		chunkFile.Close()
		chunkNum++
	}

	return nil
}
