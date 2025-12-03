#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 YEAR DAY" >&2
  exit 1
fi

if [[ ! -f "go.mod" ]]; then
  echo "run this script from the repository root so go.mod is available" >&2
  exit 1
fi

year="$1"
day_input="$2"

if [[ ! "$year" =~ ^[0-9]{4}$ ]]; then
  echo "YEAR must be a four-digit number" >&2
  exit 1
fi

if [[ ! "$day_input" =~ ^[0-9]{1,2}$ ]]; then
  echo "DAY must be 1-25 (e.g. 1 or 07)" >&2
  exit 1
fi

day="$(printf "%02d" "$day_input")"
day_dir="${year}/day${day}"
module_path="$(go list -m)"

mkdir -p "$day_dir"

create_file() {
  local path="$1"
  local contents="$2"
  if [[ -f "$path" ]]; then
    echo "skip ${path} (already exists)"
  else
    printf "%s" "$contents" >"$path"
    echo "create ${path}"
  fi
}

main_go=$(cat <<EOF
package main

import (
	"fmt"

	"${module_path}/internal/aoc"
)

func main() {
	lines := aoc.MustReadLines("input.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}
EOF
)

main_test_go=$(cat <<'EOF'
package main

import "testing"

func TestPart1(t *testing.T) {
	t.Skip("add part 1 tests")
}

func TestPart2(t *testing.T) {
	t.Skip("add part 2 tests")
}
EOF
)

day_readme=$(cat <<EOF
# ${year} Day ${day}

- **Puzzle:** https://adventofcode.com/${year}/day/$((10#$day))
- **Quick notes:** TODO

## Sample

Place the provided example input in \`sample.txt\` to use it for tests.
EOF
)

touch "${day_dir}/input.txt"
create_file "${day_dir}/main.go" "$main_go"
create_file "${day_dir}/main_test.go" "$main_test_go"
create_file "${day_dir}/README.md" "$day_readme"
create_file "${day_dir}/sample.txt" ""
