package tests

import (
	"os"
	"testing"

	"github.com/mauroandocilla/cli-squeezy-pix/internal/optimizer"
	"github.com/mauroandocilla/cli-squeezy-pix/tests/utils"
)

func TestConvertToFormat(t *testing.T) {
	inputPath := "testdata/input.jpg"
	outputPath := "testdata/output.png"

	// Create a dummy image file for testing using the utility function
	err := utils.CreateDummyImage(inputPath)
	if err != nil {
		t.Fatalf("Failed to create dummy image file: %v", err)
	}

	// Defer the cleanup of the input file and handle the error
	defer func() {
		if err := os.Remove(inputPath); err != nil {
			t.Errorf("Failed to remove input image file: %v", err)
		}
	}()

	// Run the ConvertToFormat function
	err = optimizer.ConvertToFormat(inputPath, outputPath)
	if err != nil {
		t.Errorf("ConvertToFormat() failed: %v", err)
	}

	// Verify that the output image file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Expected output image file not found: %v", outputPath)
	} else {
		// Clean up the output image file and handle the error
		if err := os.Remove(outputPath); err != nil {
			t.Errorf("Failed to remove output image file: %v", err)
		}
	}
}
