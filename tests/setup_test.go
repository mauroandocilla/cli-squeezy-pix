package tests

import (
	"os"
	"testing"

	"github.com/mauroandocilla/cli-squeezy-pix/tests/utils"
)

func TestMain(m *testing.M) {
	// Use the SetupAndRunTests function to run the tests
	exitCode := utils.SetupAndRunTests(m)

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
