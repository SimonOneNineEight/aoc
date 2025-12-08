package main

import (
	"fmt"

	"github.com/SimonOneNineEight/aoc/internal/aoc"
)

func main() {
	lines := aoc.MustReadLines("input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	count := 0
	rows := len(lines)
	cols := len(lines[0])
	directionCol := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	directionRow := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if lines[row][col] != '@' {
				continue
			}

			rolls := 0
			for i := 0; i < 8; i++ {
				newRow, newCol := row+directionRow[i], col+directionCol[i]
				if newRow < 0 || newCol < 0 || newRow >= rows || newCol >= cols {
					continue
				}
				if lines[newRow][newCol] == '@' {
					rolls++
				}
			}

			if rolls < 4 {
				count++
			}
		}
	}

	return count
}

func part2(lines []string) int {
	sum := 0
	deleted := -1

	for deleted != 0 {
		count := 0
		rows := len(lines)
		cols := len(lines[0])
		directionCol := []int{-1, 0, 1, -1, 1, -1, 0, 1}
		directionRow := []int{-1, -1, -1, 0, 0, 1, 1, 1}
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				if lines[row][col] != '@' {
					continue
				}

				rolls := 0
				for i := 0; i < 8; i++ {
					newRow, newCol := row+directionRow[i], col+directionCol[i]
					if newRow < 0 || newCol < 0 || newRow >= rows || newCol >= cols {
						continue
					}
					if lines[newRow][newCol] == '@' {
						rolls++
					}
				}

				if rolls < 4 {
					count++
					line := []rune(lines[row])
					line[col] = ','
					lines[row] = string(line)
				}
			}
		}

		deleted = count

		if deleted > 0 {
			sum += deleted
		}
	}

	return sum
}

