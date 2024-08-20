package tests

import (
	"os"
	"testing"

	"github.com/mauroandocilla/cli-squeezy-pix/internal/optimizer"
	"github.com/mauroandocilla/cli-squeezy-pix/tests/utils"
)

func TestGenerateResponsiveImages(t *testing.T) {
	inputPath := "testdata/input.jpg"
	widths := []int{320, 640, 1024}
	outputPaths := []string{
		"testdata/output-320px.jpg",
		"testdata/output-640px.jpg",
		"testdata/output-1024px.jpg",
	}

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

	// Run the GenerateResponsiveImages function
	err = optimizer.GenerateResponsiveImages(inputPath, outputPaths, widths)
	if err != nil {
		t.Errorf("GenerateResponsiveImages() failed: %v", err)
	}

	// Verify that the output image files were created and clean them up after the loop
	for _, path := range outputPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Expected output image file not found: %v", path)
		} else {
			// Remove the output file after checking its existence
			if err := os.Remove(path); err != nil {
				t.Errorf("Failed to remove output image file: %v", err)
			}
		}
	}
}
