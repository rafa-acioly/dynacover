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
	Use:   "cover",
	Short: "generate cover image",
	Run:   generateCover,
}

func generateCover(cmd *cobra.Command, args []string) {
	templateName, _ := cmd.Flags().GetString("template")
	template := templates.GetTemplateSetting(templateName)

	coverImage, err := templates.OpenImage(template.GetSourcePath())
	if err != nil {
		log.Fatalf("could not open template cover image: %v", err)
	}

	// TODO: get a list of the template.PlaceHolderQuantity followers from twitter
	cwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", cwd, "templates/stubs/account.png")
	fmt.Println("Path to placeholder: " + path)
	placeholder, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("could not open stub file: %v", err)
	}

	// TODO: zip the elements from template along with the followers list, this way we can write a single for loop
	for _, element := range template.GetElementsSettings() {

		croppedImage := templates.CropIntoCircle(placeholder, 1)
		croppedImage = templates.ResizeImage(croppedImage, element.Width, element.Height)

		coverImage = templates.PasteImage(coverImage, croppedImage, image.Pt(element.PosX, element.PosY))
	}

	if err := templates.SaveImage(coverImage, fmt.Sprintf("%s/%s", cwd, "templates/final.png")); err != nil {
		log.Fatalf("could not save final image: %v", err)
	}
}
