package drawer

import (
	"github.com/fogleman/gg"
)

// Drawer struct
type Drawer struct{}

// DrawMap draws an 16 by 16 square map
func (d *Drawer) DrawMap(ctx *gg.Context) error {
	ctx.SetRGB(255, 0, 0)

	var x, y, w, h float64
	x, y, w, h = 0, 0, 25, 25

	ctx.DrawRectangle(x, y, w, h)
	ctx.Fill()
	err := ctx.SavePNG("out.png")
	return err
}
