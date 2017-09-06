package mapmaker

import (
	"github.com/fogleman/gg"
)

// Drawer interface draws a map
type Drawer interface {
	DrawMap(*gg.Context) error
}
