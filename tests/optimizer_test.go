package tests

import (
	"os"
	"testing"

	"github.com/mauroandocilla/cli-squeezy-pix/internal/optimizer"
	"github.com/mauroandocilla/cli-squeezy-pix/tests/utils"
)

func TestOptimizeImage(t *testing.T) {
	inputPath := "testdata/input.jpg"
	outputPath := "testdata/output.jpg"

	// Create a valid dummy image file for testing using the utility function
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

	err = optimizer.OptimizeImage(inputPath, outputPath, 85)
	if err != nil {
		t.Errorf("OptimizeImage() failed: %v", err)
	}

	// Verify that the output image file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Expected output image file not found: %v", outputPath)
	}

	// Defer the cleanup of the output file and handle the error
	defer func() {
		if err := os.Remove(outputPath); err != nil {
			t.Errorf("Failed to remove output image file: %v", err)
		}
	}()
}
