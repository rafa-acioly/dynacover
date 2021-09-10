package templates


import (
	"fmt"
	"os"
)

// NewColorfulTemplate returns an instance of Template for the basic cover image
func NewColorfulTemplate() Template {
	return ColorFulTemplate{
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

// ColorFulTemplate represents all information about the basic cover image
type ColorFulTemplate struct {
	Height int
	Width int
	Elements []ElementProperty
}

func (c ColorFulTemplate) GetHeight() int {
	return c.Height
}

func (c ColorFulTemplate) GetWidth() int {
	return c.Width
}

func (c ColorFulTemplate) GetSourcePath() string {
	curDir, _ := os.Getwd()

	return fmt.Sprintf("%s/%s", curDir, "stub/colorfull.png")
}

func (c ColorFulTemplate) GetElementsSettings() []ElementProperty {
	return c.Elements
}