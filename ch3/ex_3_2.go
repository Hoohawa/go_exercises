package main

import (
	"flag"
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

var function = flag.Int("f", 1, "-f [1,2,3] // to chose a different graph function")

func main() {
	flag.Parse()
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

			fmt.Printf("<polygon points='%g %g %g %g %g %g %g %g'/>",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
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
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	var z float64
	switch {
	case *function == 1:
		z = Bumps(x, y)
	case *function == 2:
		z = Dome(x, y)
	case *function == 3:
		z = Crazy(x, y)
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func Bumps(x, y float64) float64 {
	return math.Sin(x) * math.Cos(y) / 5
}

func Dome(x, y float64) float64 {
	var r = math.Hypot(x, y)
	return -(x*x+y*y)*r*0.001 + 1
}

func Crazy(x, y float64) float64 {
	var r = math.Hypot(x, y)
	return (math.Sin(math.Sqrt(x * 2 * y))) / r
}
