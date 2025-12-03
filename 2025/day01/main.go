package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/SimonOneNineEight/aoc/internal/aoc"
)

func main() {
	lines := aoc.MustReadLines("./input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	current := 50
	count := 0
	for _, turn := range lines {
		direction := turn[:1]
		rotationStr := turn[1:]

		rotation, rotationErr := strconv.Atoi(rotationStr)
		if rotationErr != nil {
			log.Fatalf("parse: %v", rotationErr)
			return 0
		}

		switch direction {
		case "L":
			current = (current - rotation) % 100
		case "R":
			current = (current + rotation) % 100
		}

		if current < 0 {
			current += 100
		}
		if current == 0 {
			count += 1
		}

	}

	return count
}

func part2(lines []string) int {
	current := 50
	count := 0
	for _, turn := range lines {
		direction := turn[:1]
		rotationStr := turn[1:]

		rotation, rotationErr := strconv.Atoi(rotationStr)
		if rotationErr != nil {
			log.Fatalf("parse: %v", rotationErr)
			return 0
		}

		switch direction {
		case "L":
			count += rotation / 100
			remain := rotation % 100
			distance := current
			current = (current - remain) % 100
			if distance != 0 && current <= 0 {
				count += 1
			}

			if current < 0 {
				current += 100
			}
		case "R":
			count += rotation / 100
			remain := rotation % 100
			if current+remain > 99 {
				count += 1
			}
			current = (current + remain) % 100
		}

	}

	return count
}
