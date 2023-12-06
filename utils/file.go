package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadFile reads a file path passed as argument and returns the content of the file slice of strings, breaking the file by lines
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, nil
}

func PrintLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
