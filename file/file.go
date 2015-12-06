package file

import (
	"bufio"
	"log"
	"os"
)

func ReadAll(path string) *[]string {
	file, _err := os.Open(path)
	if _err != nil {
		log.Fatal("File cannot be opened : " + path)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return &lines
}
