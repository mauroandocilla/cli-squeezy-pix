package cmd

import (
	"fmt"
	"log"

	"github.com/mauroandocilla/cli-squeezy-pix/help"
	"github.com/mauroandocilla/cli-squeezy-pix/internal/optimizer"
	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
var ConvertCmd = &cobra.Command{
	Use:     "conv",
	Aliases: []string{"convert"},
	Short:   help.ConvertShortHelpText,
	Long:    help.ConvertHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		// Call the image conversion function
		err := optimizer.ConvertToFormat(inputPath, outputPath)
		if err != nil {
			log.Fatalf("Error converting image: %v", err)
		} else {
			fmt.Printf("Image successfully converted to %s\n")
		}
	},
}

func init() {
	// Config flags for the convert command
	ConvertCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input image file")
	ConvertCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output image file")

	if err := ConvertCmd.MarkFlagRequired("input"); err != nil {
		log.Fatalf("Error marking flag 'input' as required: %v", err)
	}
	if err := ConvertCmd.MarkFlagRequired("output"); err != nil {
		log.Fatalf("Error marking flag 'output' as required: %v", err)
	}
}
