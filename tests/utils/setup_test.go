package utils

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Use the SetupAndRunTests function to run the tests
	exitCode := SetupAndRunTests(m)

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
