package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.CMYK{C: 218, M: 6, Y: 218}}

const colorIndex = 1

func main() {
	outFile, err := os.Create("./out.data")
	if err == nil {
		err = lissajous(outFile)
	}
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

func lissajous(out io.Writer) error {
	const (
		cycles = 10
		res    = 0.001
		size   = 320
		frames = 24
		delay  = 8
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: frames}
	phase := 0.0
	for i := 0; i < frames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return gif.EncodeAll(out, &anim)
}
