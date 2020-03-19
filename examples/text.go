package main

import (
	"fmt"

	"github.com/zaklaus/gg"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("slkscr.ttf", 96); err != nil {
		panic(err)
	}
	w, h, asc, desc, fasc, fdesc := dc.MeasureTextMetrics("Hello, world!")
	fmt.Printf("w %f h %f asc %f desc %f fasc %f fdesc %f\n", w, h, asc, desc, fasc, fdesc)
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}
