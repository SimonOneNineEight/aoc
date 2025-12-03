package aoc

import (
	"bufio"
	"os"
)

// ReadLines returns every line in the file located at path.
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// MustReadLines is the panic-on-error variant of ReadLines.
func MustReadLines(path string) []string {
	lines, err := ReadLines(path)
	if err != nil {
		panic(err)
	}
	return lines
}
