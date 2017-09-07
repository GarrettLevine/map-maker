package drawer

import (
	"math"

	"github.com/fogleman/gg"
)

const squareSize int = 25

// Drawer struct
type Drawer struct{}

type dimensions struct {
	x float64
	y float64
	w float64
	h float64
}

// DrawMap draws an a square based on the context size
func (d *Drawer) DrawMap(ctx *gg.Context) error {

	sz := ctx.Height()
	yCounter := 0
	xCounter := 0
	drawCounter := 0
	drawLimit := math.Pow(float64(sz/squareSize), float64(2))

	for i := 0; drawCounter <= int(drawLimit); i += squareSize {
		drawCounter++
		pickColour(ctx, i)

		drawRectangle(ctx, dimensions{
			x: float64(xCounter),
			y: float64(yCounter),
			w: float64(squareSize),
			h: float64(squareSize),
		})

		xCounter += squareSize

		if drawCounter%(sz/squareSize) == 0 {
			yCounter += squareSize
			xCounter = 0
		}
	}

	err := ctx.SavePNG("out.png")
	return err
}

func drawRectangle(ctx *gg.Context, dim dimensions) {
	ctx.DrawRectangle(dim.x, dim.y, dim.w, dim.h)
	ctx.Fill()
}
