package mapmaker

import (
	"github.com/fogleman/gg"
)

// Drawer interface draws a map
type Drawer interface {
	DrawMap(*gg.Context) error
}

// Save interface saves the image
// type Save interface {
// 	SaveImage(image.Image) (os.File, error)
// }
