package utils_test

import (
	"os"
	"testing"

	"github.com/mauroandocilla/cli-squeezy-pix/tests/utils"
)

func TestCreateDummyImage(t *testing.T) {
	// Define the path where the dummy image will be created
	imagePath := "testdata/dummy_image.jpg"

	// Call CreateDummyImage to create the image
	err := utils.CreateDummyImage(imagePath)
	if err != nil {
		t.Fatalf("Failed to create dummy image: %v", err)
	}

	// Check that the image file was created successfully
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		t.Errorf("Expected image file not found: %v", imagePath)
	}

	// Clean up the created image file after the test
	defer func() {
		if err := os.Remove(imagePath); err != nil {
			t.Errorf("Failed to remove dummy image file: %v", err)
		}
	}()
}
