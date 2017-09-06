package save

import (
	"image"
	"os"
)

// Save struct
type Save struct{}

// SaveImage returns a created file and an error
func (s *Save) SaveImage(img image.Image) (*os.File, error) {
	file, err := os.Create("map.jpeg")
	if err != nil {
		panic("error creating file")
	}

	return file, err
}
