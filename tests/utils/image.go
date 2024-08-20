package utils

import (
	"github.com/disintegration/imaging"
	"image/color"
)

// CreateDummyImage creates a simple 1024x1024 image with a solid color and saves it as a JPEG at the specified path.
func CreateDummyImage(path string) error {
	// Create a 1024x1024 image with a solid color
	img := imaging.New(1024, 1024, color.NRGBA{R: 200, G: 100, B: 100, A: 255})

	// Save the image as a JPEG to the specified path
	if err := imaging.Save(img, path); err != nil {
		return err
	}
	return nil
}
