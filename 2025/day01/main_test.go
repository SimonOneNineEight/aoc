package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// TestMainCallsMain demonstrates how to capture stdout when calling main().
func TestMainCallsMain(t *testing.T) {
	const inputPath = "input.txt"

	originalInput, err := os.ReadFile(inputPath)
	if err != nil {
		t.Fatalf("read %s: %v", inputPath, err)
	}
	t.Cleanup(func() {
		if err := os.WriteFile(inputPath, originalInput, 0o644); err != nil {
			t.Fatalf("restore %s: %v", inputPath, err)
		}
	})

	sample := []byte("line1\nline2\n")
	if err := os.WriteFile(inputPath, sample, 0o644); err != nil {
		t.Fatalf("write sample: %v", err)
	}

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = origStdout })

	main()

	if err := w.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("copy output: %v", err)
	}

	want := "Part 1: 2\nPart 2: 4\n"
	if buf.String() != want {
		t.Fatalf("unexpected output\nwant:\n%s\ngot:\n%s", want, buf.String())
	}
}
