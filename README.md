## Advent of Code Go Workspace

This repository hosts all of my Advent of Code solutions using Go. Each competition year lives on its own git branch (`2025`, `2026`, ...), but every branch shares the same module name (`github.com/SimonOneNineEight/aoc`) so imports stay stable.

### Layout

```
.
├── 2025/
│   └── day01/         # Daily solutions (one folder per day)
├── internal/
│   └── aoc/           # Shared helper packages
├── scripts/           # Tooling for scaffolding & input downloads
├── go.mod
└── README.md
```

### Prerequisites

- Go 1.25+
- Puzzle input download requires a valid Advent of Code session cookie stored as `AOC_SESSION` in `.env` (never commit this file).

### Workflow

1. **Checkout the correct branch**
   ```sh
   git checkout 2025
   ```
2. **Scaffold the new day**
   ```sh
   ./scripts/new-day.sh 2025 03
   ```
   This creates `2025/day03` with starter `main.go`, `main_test.go`, `README.md`, and placeholder inputs.
3. **Fetch the real puzzle input**
   ```sh
   ./scripts/fetch-input.sh 2025 03
   ```
4. Paste the sample input into `sample.txt`, solve the puzzle in `main.go`, and document notes in the day README.
5. Run and test:
   ```sh
   go run ./2025/day03
   go test ./2025/day03
   ```
6. Commit when the day is complete.

### Scripts

- `scripts/new-day.sh YEAR DAY`
  - Creates the day folder (`YEAR/dayDD`), starter Go files, README, and placeholder inputs.
  - Detects the module path via `go list -m` so imports remain correct even if the module name changes later.
  - Skips files that already exist to avoid overwriting work.

- `scripts/fetch-input.sh YEAR DAY`
  - Loads `AOC_SESSION` from `.env` (if present) or the environment and downloads the puzzle input via `curl`.
  - Saves the result to `YEAR/dayDD/input.txt`.

### Notes

- All helper code belongs under `internal/aoc/` where it can be imported from any day (e.g., `github.com/SimonOneNineEight/aoc/internal/aoc`).
- Keep sensitive files out of git—`.env`, `session.txt`, and downloaded inputs are already ignored.
- Feel free to adjust the templates or scripts as new patterns emerge during Advent of Code.
