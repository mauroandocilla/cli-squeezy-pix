package utils

import (
	"os"
	"testing"
)

// SetupAndRunTests handles the creation of the testdata directory, running tests, and cleaning up.
func SetupAndRunTests(m *testing.M) int {
	// Setup: Create the testdata directory before running the tests
	err := os.Mkdir("testdata", 0755)
	if err != nil && !os.IsExist(err) {
		panic("Failed to create testdata directory")
	}

	// Run the tests
	exitCode := m.Run()

	// Teardown: Remove the testdata directory after running the tests
	err = os.RemoveAll("testdata")
	if err != nil {
		panic("Failed to remove testdata directory")
	}

	return exitCode
}
