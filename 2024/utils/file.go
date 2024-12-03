package utils

import (
	"bufio"
	"os"
)

func getFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}

func GetFileContent(path string) []string {
	file := getFile(path)

	lines := []string{}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
