#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 YEAR DAY" >&2
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

day="$(printf "%d" "$day_input")"
day_dir="${year}/day$(printf "%02d" "$day_input")"
input_path="${day_dir}/input.txt"

if [[ -f ".env" ]]; then
  set -a
  # shellcheck disable=SC1091
  source ".env"
  set +a
fi

session="${AOC_SESSION:-}"
if [[ -z "$session" ]]; then
  echo "AOC_SESSION is not set. Export it or place it in .env" >&2
  exit 1
fi

mkdir -p "$day_dir"

echo "downloading input to ${input_path}"
curl -sSf \
  --cookie "session=${session}" \
  "https://adventofcode.com/${year}/day/${day}/input" \
  -o "${input_path}"
