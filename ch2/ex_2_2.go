// To run the example first build the tempconv package with "$ go install .../ch2/tempconv"
// Then do: $ run ex_2_1.go
package main

import (
	"flag"
	"fmt"
	"github.com/hoohawa/go_exercises/ch2/mtrconv"
	"os"
)

var from = flag.String("f", "", "-f [m|mi|in],[lbs,oz,g],[c,f,k]")
var to = flag.String("t", "", "-t [m|mi|in],[lbs,oz,g],[c,f,k]")
var value = flag.Float64("v", 0.0, "-v [float64]")

func main() {
	if len(os.Args) > 1 {
		flag.Parse()
	} else {
		fmt.Printf("Please input from metric, to metric, value to convet\n: ")
		fmt.Scanf("%s %s %f", from, to, value)
	}

	switch {
	// Length conversions
	case *from == "m" && *to == "mi":
		fmt.Print(mtrconv.MToMI(mtrconv.Meter(*value)))
	case *from == "m" && *to == "in":
		fmt.Print(mtrconv.MToIN(mtrconv.Meter(*value)))

	case *from == "mi" && *to == "m":
		fmt.Print(mtrconv.MIToM(mtrconv.Mile(*value)))
	case *from == "mi" && *to == "in":
		fmt.Print(mtrconv.MIToIN(mtrconv.Mile(*value)))

	case *from == "in" && *to == "m":
		fmt.Print(mtrconv.INToM(mtrconv.Inch(*value)))
	case *from == "in" && *to == "in":
		fmt.Print(mtrconv.INToMI(mtrconv.Inch(*value)))

	// Weight conversions
	case *from == "g" && *to == "lbs":
		fmt.Print(mtrconv.GToP(mtrconv.Gram(*value)))
	case *from == "g" && *to == "oz":
		fmt.Print(mtrconv.GToOZ(mtrconv.Gram(*value)))

	case *from == "lbs" && *to == "g":
		fmt.Print(mtrconv.PToG(mtrconv.Pound(*value)))
	case *from == "lbs" && *to == "oz":
		fmt.Print(mtrconv.PToOZ(mtrconv.Pound(*value)))

	case *from == "oz" && *to == "g":
		fmt.Print(mtrconv.OZToG(mtrconv.Ounce(*value)))
	case *from == "oz" && *to == "lbs":
		fmt.Print(mtrconv.OZToP(mtrconv.Ounce(*value)))

	// Temperature conversions
	case *from == "c" && *to == "k":
		fmt.Print(mtrconv.CToK(mtrconv.Celsius(*value)))
	case *from == "c" && *to == "f":
		fmt.Print(mtrconv.CToF(mtrconv.Celsius(*value)))

	case *from == "k" && *to == "f":
		fmt.Print(mtrconv.KToC(mtrconv.Kelvin(*value)))
	case *from == "k" && *to == "c":
		fmt.Print(mtrconv.KToF(mtrconv.Kelvin(*value)))

	case *from == "f" && *to == "c":
		fmt.Print(mtrconv.FToC(mtrconv.Fahrenheit(*value)))
	case *from == "f" && *to == "k":
		fmt.Print(mtrconv.FToK(mtrconv.Fahrenheit(*value)))
	}
}
