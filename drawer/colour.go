package drawer

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

func pickColour(ctx *gg.Context, number int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(2)

	red, green, blue := 0, 0, 0
	if randomNumber == 1 {
		red, green, blue = 152, 251, 152
	} else {
		red, green, blue = 32, 178, 170
	}

	ctx.SetRGB255(red, green, blue)
}
