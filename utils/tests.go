package utils

import (
	"bytes"
	"os"
)

func CaptureOutput(f func()) string {
	// Save the original stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Restore stdout and capture the output
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old

	return buf.String()
}
