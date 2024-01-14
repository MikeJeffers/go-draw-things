package main

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/gg"
)

func drawPath(path Path, step float64, dc *gg.Context) {
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(2)
	nilPt := Point[int]{x: -1, y: -1}
	prev := Point[int]{x: -1, y: -1}

	for _, v := range path.points {
		if prev != nilPt {
			dc.DrawLine(scaleAndOffset(prev.x, step), scaleAndOffset(prev.y, step), scaleAndOffset(v.x, step), scaleAndOffset(v.y, step))
		}
		prev = v
	}
	dc.Stroke()
}

func drawToImage(width, height, step int, paths []Path) {
	dc := gg.NewContext(width*step, height*step)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, float64(width*step), float64(height*step))
	dc.Fill()

	for _, v := range paths {
		drawPath(v, float64(step), dc)
	}

	num := int(rand.Float32() * 100)
	err := dc.SavePNG(fmt.Sprint(num, "_out.jpg"))
	if err != nil {
		fmt.Printf("Output file failed to save: %s", err.Error())
	}
}
