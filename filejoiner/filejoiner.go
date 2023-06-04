package filejoiner

import (
	"io"
	"os"
	"path/filepath"
	"sort"
)

func JoinFiles(splitDir, outputFile string) error {
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	files, err := os.ReadDir(splitDir)
	if err != nil {
		return err
	}

	// Sort the files by name to ensure correct order
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		splitFile := filepath.Join(splitDir, file.Name())
		// if !strings.HasPrefix(file.Name(), filepath.Base(outputFile)) {
		// 	continue
		// }

		input, err := os.Open(splitFile)
		if err != nil {
			return err
		}
		defer input.Close()

		_, err = io.Copy(output, input)
		if err != nil {
			return err
		}
	}

	return nil
}
