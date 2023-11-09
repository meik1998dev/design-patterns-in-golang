package helpers

import (
	"bytes"
	"io"
	"os"
)

func CaptureOutput(f func()) string {
	// Keep the original stdout to restore later.
	originalStdout := os.Stdout

	// Create a new pipe.
	r, w, _ := os.Pipe()

	// Set the current stdout to the write end of the pipe.
	os.Stdout = w

	// Run the function, capturing its output.
	f()

	// Close the write end of the pipe so we can read from it.
	w.Close()

	// Reset stdout back to the original value.
	os.Stdout = originalStdout

	// Read the captured output from the read end of the pipe.
	var buf bytes.Buffer
	io.Copy(&buf, r) // This should be io.Copy, not fmt.Fprint.
	r.Close()

	// Return the captured output as a string.
	return buf.String()
}
