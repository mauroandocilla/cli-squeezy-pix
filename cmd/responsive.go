package cmd

import (
	"fmt"
	"log"

	"github.com/mauroandocilla/cli-squeezy-pix/help"
	"github.com/mauroandocilla/cli-squeezy-pix/internal/optimizer"
	"github.com/spf13/cobra"
)

// Variables for the responsive command
var (
	widths []int
)

// ResponsiveCmd represents the responsive command
var ResponsiveCmd = &cobra.Command{
	Use:     "resp",
	Aliases: []string{"responsive"},
	Short:   help.ResponsiveShortHelpText,
	Long:    help.ResponsiveHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		// Call the responsive image generation function
		outputPaths := generateOutputPaths(outputPath, widths)
		err := optimizer.GenerateResponsiveImages(inputPath, outputPaths, widths)
		if err != nil {
			log.Fatalf("Error generating responsive images: %v", err)
		} else {
			fmt.Println("Responsive images generated successfully!")
		}
	},
}

func init() {
	// Config flags for the responsive command
	ResponsiveCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input image file")
	ResponsiveCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output image base path")
	ResponsiveCmd.Flags().IntSliceVarP(&widths, "widths", "w", []int{320, 640, 1024}, "Widths for responsive images")

	if err := ResponsiveCmd.MarkFlagRequired("input"); err != nil {
		log.Fatalf("Error marking flag 'input' as required: %v", err)
	}
	if err := ResponsiveCmd.MarkFlagRequired("output"); err != nil {
		log.Fatalf("Error marking flag 'output' as required: %v", err)
	}
}

// generateOutputPaths generates output paths based on the base name and widths
func generateOutputPaths(basePath string, widths []int) []string {
	outputPaths := make([]string, len(widths))
	for i, width := range widths {
		outputPaths[i] = fmt.Sprintf("%s-%dpx.jpg", basePath, width)
	}
	return outputPaths
}
