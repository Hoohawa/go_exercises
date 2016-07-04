package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
	color.RGBA{0x0F, 0x30, 0x00, 0xFF},
	color.RGBA{0xAA, 0x00, 0x0F, 0xFF},
	color.RGBA{0x0F, 0x00, 0x0F, 0xFF},
	color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			var colorIndex = uint8(x*10) % uint8(len(pallete)-1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
