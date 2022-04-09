package templates

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"math"
)

type ElementProperty struct {
	PosX   int
	PosY   int
	Width  int
	Height int
}

type Template interface {
	GetHeight() int
	GetWidth() int
	GetSourcePath() string
	GetElementsSettings() []ElementProperty
}

// GetTemplateSetting is a factory that will return the template corresponding to the template name
func GetTemplateSetting(templateName string) Template {
	templates := map[string]func() Template{
		"basic":    NewBasicTemplate,
		"colorful": NewColorfulTemplate,
		"neon":     NewNeonTemplate,
		"sponsors": NewSponsorsTemplate,
	}

	if _, exist := templates[templateName]; !exist {
		log.Fatalf("could not load template from string: %s", templateName)
	}

	return templates[templateName]()
}

// CropImageCenter will crop the image on the center
func CropImageCenter(img image.Image, width, height int) *image.NRGBA {
	return imaging.CropCenter(img, width, height)
}

func ResizeImage(img image.Image, width, height int) *image.NRGBA {
	return imaging.Resize(img, width, height, imaging.ResampleFilter{})
}

// OpenImage will load the image from the disk
func OpenImage(imagePath string, opts ...imaging.DecodeOption) (image.Image, error) {
	return imaging.Open(imagePath, opts...)
}

// PasteImage will place an image on top of other on the given position
func PasteImage(bg image.Image, img image.Image, position image.Point) *image.NRGBA {
	return imaging.Paste(bg, img, position)
}

// SaveImage writes the given image on the disk
func SaveImage(img image.Image, location string) error {
	return imaging.Save(img, location)
}

func CropIntoCircle(src image.Image, factor float64) image.Image {
	d := src.Bounds().Dx()
	if src.Bounds().Dy() < d {
		d = src.Bounds().Dy()
	}
	dst := imaging.CropCenter(src, d, d)
	radius := float64(d) / 2
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			xf := float64(x)
			yf := float64(y)
			delta := math.Sqrt((xf-radius)*(xf-radius)+(yf-radius)*(yf-radius)) + factor - radius

			switch {
			case delta > factor:
				dst.Set(x, y, color.Transparent)
			case delta > 0:
				m := 1 - delta/factor
				c := dst.NRGBAAt(x, y)
				c.A = uint8(float64(c.A) * m)
				//dst.SetNRGBA(x, y, c)
				dst.Set(x, y, color.Transparent)
			}
		}
	}
	return dst
}
