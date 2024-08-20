package optimizer

import (
	"github.com/disintegration/imaging"
)

// ConvertToFormat converts the input image to the specified format
func ConvertToFormat(inputPath, outputPath string) error {
	// Open the input image
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// Save the image to the output path in its current format
	err = imaging.Save(img, outputPath)
	if err != nil {
		return err
	}

	return nil
}
