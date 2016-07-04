package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill:white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			// Skip invalid polygons
			if isNotPlottable(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Fprintf(os.Stderr,
					"ERROR:'%g %g %g %g %g %g %g %g'\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				continue
			}
			red, green := calculateRedGreen(i, j)
			fmt.Printf("<polygon points='%g %g %g %g %g %g %g %g' style='stroke:rgb(%d, %d, 0)'/>",
				ax, ay, bx, by, cx, cy, dx, dy, red, green)
		}
	}
	fmt.Println("</svg>")
}

func calculateRedGreen(i, j int) (int, int) {
	_, _, z := calculateXYZ(i, j)
	positiveZ := z + 0.2    // Z is approx. [-0.2, 1]
	colorScale := 0.5 / 256 // Scale to make it look nice
	red := positiveZ / colorScale
	green := (0.5 - positiveZ) / colorScale
	return int(red), int(green)
}

func isNotPlottable(coords ...float64) bool {
	for _, c := range coords {
		if math.IsNaN(c) || math.IsInf(c, -1) || math.IsInf(c, +1) {
			return true
		}
	}
	return false
}

func corner(i, j int) (float64, float64) {
	x, y, z := calculateXYZ(i, j)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func calculateXYZ(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	return x, y, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
