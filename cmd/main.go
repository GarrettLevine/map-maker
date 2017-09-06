package main

import (
	"map-maker/drawer"

	"github.com/fogleman/gg"
)

func main() {
	d := drawer.Drawer{}
	// s := save.Save{}
	ctx := gg.NewContext(400, 400)

	err := d.DrawMap(ctx)
	if err != nil {
		panic("error creating map")
	}

	// file, err := s.SaveImage(mapImage)
	// if err != nil {
	// 	panic("error creating image file")
	// }

	// jpeg.Encode(file, mapImage, nil)
}
