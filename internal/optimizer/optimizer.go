package optimizer

import (
	"github.com/disintegration/imaging"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// OptimizeImage resizes and compresses the input image
func OptimizeImage(inputPath, outputPath string, quality int) error {
	// Open the image
	img, err := imaging.Open(inputPath)
	if err != nil {
		log.Printf("Error opening image: %v", err)
		return err
	}

	// Get the file extension to determine the format
	fileExtension := strings.ToLower(filepath.Ext(outputPath))

	// If quality is provided, use it directly; otherwise, determine it automatically
	if quality > 0 {
		// Save the image with the provided quality based on the file format
		switch fileExtension {
		case ".jpeg", ".jpg":
			err = imaging.Save(img, outputPath, imaging.JPEGQuality(quality))
		case ".png":
			// PNG uses lossless compression, the quality parameter is not directly used, so we can just save the image
			err = imaging.Save(img, outputPath, imaging.PNGCompressionLevel(pngCompressionLevel(quality)))
		case ".gif":
			// GIF doesn't support direct quality settings, we save it without additional compression settings
			err = imaging.Save(img, outputPath)
		default:
			log.Printf("Unsupported image format: %s", fileExtension)
			return err
		}
		if err != nil {
			log.Fatalf("Failed to save image: %v", err)
			return err
		}
	} else {
		// Automatically determine compression rate (only for JPEG)
		if fileExtension == ".jpeg" || fileExtension == ".jpg" {
			quality = 100
			maxFileSizeReduction := 0.3 // Allow up to 30% file size reduction
			initialFileSize, err := getFileSize(inputPath)
			if err != nil {
				return err
			}

			for quality > 10 {
				// Save the image with current quality
				tempOutput := "temp_output.jpg"
				err = imaging.Save(img, tempOutput, imaging.JPEGQuality(quality))
				if err != nil {
					log.Fatalf("Failed to save image: %v", err)
					return err
				}

				// Check the new file size
				compressedFileSize, err := getFileSize(tempOutput)
				if err != nil {
					return err
				}

				// Calculate the reduction ratio
				reductionRatio := float64(initialFileSize-compressedFileSize) / float64(initialFileSize)

				// If the reduction is within acceptable limits, break the loop
				if reductionRatio >= maxFileSizeReduction {
					break
				}

				// Decrease the quality and try again
				quality -= 5
			}

			// Save the final optimized image
			err = imaging.Save(img, outputPath, imaging.JPEGQuality(quality))
			if err != nil {
				log.Fatalf("Failed to save optimized image: %v", err)
				return err
			}

			// Clean up the temporary file
			if err := os.Remove("temp_output.jpg"); err != nil {
				log.Printf("Failed to remove temporary file: %v", err)
			}
		} else {
			// For PNG and GIF, no automatic quality adjustment is applied, simply save the image
			switch fileExtension {
			case ".png":
				err = imaging.Save(img, outputPath, imaging.PNGCompressionLevel(png.BestCompression))
			case ".gif":
				err = imaging.Save(img, outputPath)
			}
			if err != nil {
				log.Fatalf("Failed to save image: %v", err)
				return err
			}
		}
	}

	return nil
}

// getFileSize gets the file size in bytes
func getFileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// pngCompressionLevel maps quality percentage to png.CompressionLevel
func pngCompressionLevel(quality int) png.CompressionLevel {
	switch {
	case quality >= 75:
		return png.BestSpeed
	case quality >= 50:
		return png.DefaultCompression
	case quality >= 25:
		return png.BestCompression
	default:
		return png.BestCompression
	}
}
