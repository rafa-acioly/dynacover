package templates

import (
	"fmt"
	"os"
)

// NewSponsorsTemplate returns an instance of Template for the basic cover image
func NewSponsorsTemplate() Template {
	return SponsorsTemplate{
		Height:   1500,
		Width:    500,
		Elements: []ElementProperty{
			{
				PosX: 486,
				PosY: 272,
				Width: 130,
				Height: 130,
			},
			{
				PosX: 670,
				PosY: 272,
				Width: 130,
				Height: 130,
			},
			{
				PosX: 859,
				PosY: 270,
				Width: 132,
				Height: 132,
			},
			{
				PosX: 1049,
				PosY: 270,
				Width: 132,
				Height: 132,
			},
			{
				PosX: 1236,
				PosY: 270,
				Width: 130,
				Height: 130,
			},
		},
	}
}

// SponsorsTemplate represents all information about the basic cover image
type SponsorsTemplate struct {
	Height int
	Width int
	Elements []ElementProperty
}

func (c SponsorsTemplate) GetHeight() int {
	return c.Height
}

func (c SponsorsTemplate) GetWidth() int {
	return c.Width
}

func (c SponsorsTemplate) GetSourcePath() string {
	curDir, _ := os.Getwd()

	return fmt.Sprintf("%s/%s", curDir, "templates/stubs/sponsors.png")
}

func (c SponsorsTemplate) GetElementsSettings() []ElementProperty {
	return c.Elements
}
