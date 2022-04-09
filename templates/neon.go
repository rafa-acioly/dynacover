package templates

import (
	"fmt"
	"os"
)

// NewNeonTemplate returns an instance of Template for the basic cover image
func NewNeonTemplate() Template {
	return NeonTemplate{
		Height:   500,
		Width:    1500,
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
				Width: 132,
				Height: 132,
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

// NeonTemplate represents all information about the basic cover image
type NeonTemplate struct {
	Height int
	Width int
	Elements []ElementProperty
}

func (c NeonTemplate) GetHeight() int {
	return c.Height
}

func (c NeonTemplate) GetWidth() int {
	return c.Width
}

func (c NeonTemplate) GetSourcePath() string {
	curDir, _ := os.Getwd()

	return fmt.Sprintf("%s/%s", curDir, "templates/stubs/neon.png")
}

func (c NeonTemplate) GetElementsSettings() []ElementProperty {
	return c.Elements
}
