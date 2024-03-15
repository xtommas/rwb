package windows

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SaveImage(sourcePath, destPath string) error {
	if !strings.Contains(destPath, ".png") || !strings.Contains(destPath, ".jpeg") || !strings.Contains(destPath, ".jpg") {
		destPath = destPath + "wallpaper.png"
	}

	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("Couldn't open destination file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("Couldn't copy dest from source: %v", err)

	}

	inputFile.Close()

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't remove source file: %v", err)
	}
	return nil
}
