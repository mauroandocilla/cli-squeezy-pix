package cmd

import (
	"fmt"
	"log"

	"github.com/mauroandocilla/cli-squeezy-pix/help"
	"github.com/mauroandocilla/cli-squeezy-pix/internal/optimizer"
	"github.com/spf13/cobra"
)

// Variables for the optimize command
var (
	inputPath  string
	outputPath string
	quality    int
)

// Cmd represents the optimize command
var OptimizeCmd = &cobra.Command{
	Use:     "opt",
	Aliases: []string{"optimize"},
	Short:   help.OptimizeShortHelpText,
	Long:    help.OptimizeHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		// Call the image optimization function
		err := optimizer.OptimizeImage(inputPath, outputPath, quality)
		if err != nil {
			log.Fatalf("Error optimizing image: %v", err)
		} else {
			fmt.Println("Image successfully optimized!")
		}
	},
}

func init() {
	// Config flags for the optimize command
	OptimizeCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input image file")
	OptimizeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output image file")
	OptimizeCmd.Flags().IntVarP(&quality, "quality", "q", 85, "Quality of the image (1-100)")

	if err := OptimizeCmd.MarkFlagRequired("input"); err != nil {
		log.Fatalf("Error marking flag 'input' as required: %v", err)
	}
	if err := OptimizeCmd.MarkFlagRequired("output"); err != nil {
		log.Fatalf("Error marking flag 'output' as required: %v", err)
	}
}
