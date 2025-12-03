package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/SimonOneNineEight/aoc/internal/aoc"
)

func main() {
	lines := aoc.MustReadLines("input.txt")
	fmt.Println("Part 1:", part1(lines[0]))
	fmt.Println("Part 2:", part2(lines[0]))
}

func part1(line string) int {
	line = strings.TrimSpace(line)
	if line == "" {
		return 0
	}

	ranges := strings.Split(line, ",")
	sum := 0
	pow10 := []int{1} // pow10[i] == 10^i

	for _, idRange := range ranges {
		idRange = strings.TrimSpace(idRange)
		if idRange == "" {
			continue
		}

		start, end := parseBounds(idRange)
		sum += sumRepeatedIDs(start, end, &pow10)
	}
	return sum
}

func parseBounds(raw string) (int, int) {
	startStr, endStr, ok := strings.Cut(raw, "-")
	if !ok {
		log.Fatalf("invalid range %q", raw)
	}

	start, err := strconv.Atoi(startStr)
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		log.Fatal(err)
	}
	return start, end
}

func sumRepeatedIDs(start, end int, pow10 *[]int) int {
	sum := 0
	for n := start; n <= end; n++ {
		digits := digitCount(n)
		if digits%2 == 1 {
			continue
		}

		half := digits / 2
		ensurePow10(pow10, half)
		p := (*pow10)[half]
		left, right := n/p, n%p
		if left == right {
			sum += n
		}
	}
	return sum
}

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func ensurePow10(pow10 *[]int, exp int) {
	for len(*pow10) <= exp {
		next := (*pow10)[len(*pow10)-1] * 10
		*pow10 = append(*pow10, next)
	}
}

func part2(line string) int {
	line = strings.TrimSpace(line)
	if line == "" {
		return 0
	}

	ranges := strings.Split(line, ",")
	sum := 0
	pow10 := []int{1} // pow10[i] == 10^i

	for _, idRange := range ranges {
		idRange = strings.TrimSpace(idRange)
		if idRange == "" {
			continue
		}

		start, end := parseBounds(idRange)
		sum += sumPatterenedIDs(start, end, &pow10)
	}
	return sum
}

func sumPatterenedIDs(start int, end int, pow10 *[]int) int {
	sum := 0

main:
	for n := start; n <= end; n++ {
		num := strconv.Itoa(n)
		digits := digitCount(n)
		halfDigits := digits / 2

	outer:
		for i := 1; i <= halfDigits; i++ {
			if digits%i != 0 {
				continue
			}
			chunks := sliceChunks(num, i)
			currentIndex := 0
			for currentIndex < len(chunks)-1 {
				if chunks[currentIndex] != chunks[currentIndex+1] {
					continue outer
				}
				currentIndex++
			}
			sum += n
			continue main

		}

		// ensurePow10(pow10, digits)
		//
		// for i := 1; i <= halfDigits; i++ {
		// 	if digits%i != 0 {
		// 		continue
		// 	}
		// 	target := 0
		// 	for j := 0; j <= i; j++ {
		// 		p := (*pow10)[j]
		// 		target += p * (n % p)
		// 	}
		//
		// 	current := 0
		// 	index := i + 1
		// 	for index < digits {
		// 		for k := index + 1; k <= index+i; index++ {
		// 			p := (*pow10)[k]
		// 			current += p * (n % p)
		// 		}
		// 		if current != target {
		// 			current = 0
		// 			break
		// 		} else {
		// 			index += i
		// 		}
		// 	}
		// 	sum += n
		// 	target = 0
		// 	current = 0
		// 	index = i + 1
		// }

	}

	return sum
}

func sliceChunks(s string, size int) []string {
	var out []string
	for i := 0; i < len(s); i += size {
		out = append(out, s[i:i+size])
	}
	return out
}
