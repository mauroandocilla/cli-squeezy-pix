package optimizer

import "github.com/disintegration/imaging"

func GenerateResponsiveImages(inputPath string, outputPaths []string, widths []int) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	for i, width := range widths {
		resizedImg := imaging.Resize(img, width, 0, imaging.Lanczos)
		err := imaging.Save(resizedImg, outputPaths[i])
		if err != nil {
			return err
		}
	}

	return nil
}
