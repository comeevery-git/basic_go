package experiment

import (
	"bytes"
	"log"
	"testing"
)

func TestRunMemoryTest(t *testing.T) {
	// Capture the standard output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Run the memory test function
	RunMemoryTest()

	// Check the expected log output
	expectedOutput := "== Memory test started ==\n" +
		"Initial memory usage:\n" +
		"<memory usage>\n" +
		"#####  Creating large slice\n" +
		"Memory usage after creating large slice:\n" +
		"<memory usage>\n" +
		"#####  Releasing slice reference\n" +
		"#####  Running GC\n" +
		"Memory usage after GC:\n" +
		"<memory usage>\n" +
		"== Memory test finished ==\n"
	actualOutput := buf.String()
	if actualOutput != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}