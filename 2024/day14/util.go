package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func saveImage(robots robotList, filename string) {
	im := image.NewGray(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			count := 0
			for _, r := range robots {
				if r.p.x == x && r.p.y == y {
					count++
					break
				}
			}
			if count == 0 {
				im.Set(x, y, color.White)
			} else {
				im.Set(x, y, color.Black)
			}
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, im); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
