package main

import (
	"map-maker/drawer"

	"github.com/fogleman/gg"
)

func main() {
	d := drawer.Drawer{}
	// s := save.Save{}
	ctx := gg.NewContext(600, 600)

	err := d.DrawMap(ctx)
	if err != nil {
		panic("error creating map")
	}
}
