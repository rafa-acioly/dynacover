package templates

import (
	"fmt"
	"os"
)

// NewBasicTemplate returns an instance of Template for the basic cover image
func NewBasicTemplate() Template {
	return BasicTemplate{
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

// BasicTemplate represents all information about the basic cover image
type BasicTemplate struct {
	Height int
	Width int
	Elements []ElementProperty
}

func (b BasicTemplate) GetElementsSettings() []ElementProperty {
	return b.Elements
}

func (b BasicTemplate) GetHeight() int {
	return b.Height
}

func (b BasicTemplate) GetWidth() int {
	return b.Width
}

func (b BasicTemplate) GetSourcePath() string {
	curDir, _ := os.Getwd()

	return fmt.Sprintf("%s/%s", curDir, "/stubs/basic.png")
}
