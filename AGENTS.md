# Repository Guidelines

## Project Structure & Module Organization
This workspace targets Advent of Code solutions in Go. Each challenge lives in `YEAR/dayDD/` (for example `2025/day01/`) with `main.go`, `main_test.go`, puzzle inputs, and a short README so every day stays self-contained. Shared helpers such as parsers or math utilities belong in `internal/aoc/`; keep them generic so imports like `github.com/SimonOneNineEight/aoc/internal/aoc` remain stable. Tooling for scaffolding and downloads resides in `scripts/`. Generated inputs, `.env`, and other secrets must stay untracked to avoid leaking Advent of Code credentials.

## Build, Test, and Development Commands
- `./scripts/new-day.sh 2025 03` scaffolds `2025/day03/` with starter Go files and placeholder inputs.  
- `./scripts/fetch-input.sh 2025 03` loads `AOC_SESSION` from `.env` and pulls the real puzzle input into `input.txt`.  
- `go run ./2025/day03` executes the solution for fast validation.  
- `go test ./2025/day03` runs unit tests for a single day; use `go test ./...` for all packages.

## Coding Style & Naming Conventions
Follow idiomatic Go: tabs, `gofmt` on save, and small, composable functions. Name packages lowercase and files after their role (`solver.go`, `parser.go`). Keep exported helper names descriptive (`aoc.ParseGrid`), while day-local helpers can stay unexported. Inputs go in `sample.txt` and `input.txt`, and include a brief approach note in the day README.

## Testing Guidelines
Use the Go `testing` package and favor table-driven tests with clear case labels. Mirror puzzle parts with tests such as `TestPartOne` and `TestPartTwo`. When asserting against AoC samples, load `sample.txt` for the fixture. Target full passing runs of `go test ./YEAR/dayDD ./internal/...` before opening a PR; add regression tests whenever a bug is fixed.

## Commit & Pull Request Guidelines
Write imperative commit subjects that highlight the day, e.g., `feat(2025/day03): add part two solver`. Structure the body with context and mention input quirks or helper updates. Pull requests should describe the approach, reference any related AoC discussion threads or issues, and include output screenshots only when behavior is visual. Confirm scripts were run (`new-day`, `fetch-input`, `go test`) and call out new helper APIs so reviewers can double-check other days for side effects.

## Security & Configuration Tips
Store your Advent of Code session cookie in `.env` as `AOC_SESSION` and never commit that file. When running scripts, ensure the environment exports the variable or the `.env` file exists in the repo root. Rotate the cookie each season and remove obsolete inputs from shared backups.
