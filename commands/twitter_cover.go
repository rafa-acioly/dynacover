package commands

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/rafa-acioly/dynacover/templates"
	"github.com/spf13/cobra"
	"image"
	"log"
	"os"
)

var CoverCmd = &cobra.Command{
  Use: "cover",
  Short: "generate cover image",
  Run: generateCover,
}

func generateCover(cmd *cobra.Command, args []string) {
	templateName, _ := cmd.Flags().GetString("template")
	template := templates.GetTemplateSetting(templateName)

	coverImage, err := imaging.Open(template.GetSourcePath())
	if err != nil {
		log.Fatalf("could not open template cover image: %v", err)
	}

	// TODO: get a list of the template.PlaceHolderQuantity followers from twitter
	cwd, _ := os.Getwd()
	placeholder, err := imaging.Open(fmt.Sprintf("%s/%s", cwd, "templates/stubs/400x400.png"))
	if err != nil {
		log.Fatalf("could not open stub file: %v", err)
	}

	for _, element := range template.GetElementsSettings() {

		// TODO: crop the follower image
		croppedImage := imaging.CropCenter(placeholder, element.Width, element.Height)

		// TODO: before pasting the image we need to turn it into a circle
		coverImage = imaging.Paste(coverImage, croppedImage, image.Pt(element.PosX, element.PosY))
	}

	if err := imaging.Save(coverImage, fmt.Sprintf("%s/%s", cwd, "templates/out/final.png")); err != nil {
		log.Fatalf("could not save final image: %v", err)
	}
}