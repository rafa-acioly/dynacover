package templates

import "log"

type ElementProperty struct {
	PosX int
	PosY int
	Width int
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
	templates := map[string]Template{
		"basic": NewBasicTemplate(),
		"colorful": NewColorfulTemplate(),
		"neon": NewNeonTemplate(),
		"sponsors": NewSponsorsTemplate(),
	}

	if _, exist := templates[templateName]; !exist {
		log.Fatalf("could not load template from string: %s", templateName)
	}

	return templates[templateName]
}
