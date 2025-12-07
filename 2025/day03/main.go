package main

import (
	"fmt"
	"strconv"

	"github.com/SimonOneNineEight/aoc/internal/aoc"
)

func main() {
	lines := aoc.MustReadLines("input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		n := len(line)
		number := []rune(line)
		startIndex := 0
		endIndex := 1

		for index, digit := range number {
			if index == 0 {
				continue
			}

			if digit > number[startIndex] && index != n-1 {
				startIndex = index
				endIndex = index + 1
			} else if digit > number[endIndex] {
				endIndex = index
			}
		}

		sum += int(number[startIndex]-'0')*10 + int(number[endIndex]-'0')

	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		n := len(line)
		digitsToKeep := 12
		digitsToDrop := n - digitsToKeep

		stack := make([]byte, 0, len(line))

		for i := 0; i < len(line); i++ {
			digit := line[i]

			for digitsToDrop > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
				stack = stack[:len(stack)-1]
				digitsToDrop--
			}

			stack = append(stack, digit)
		}
		jolt, err := strconv.Atoi(string(stack[:digitsToKeep]))
		if err != nil {
			fmt.Println("err")
		}
		sum += jolt
	}
	return sum
}

