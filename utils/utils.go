package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileContents(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var contentBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentBuilder.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return contentBuilder.String(), nil
}
