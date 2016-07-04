// +build ignore

// To run the example first build the tempconv package with "$ go install .../ch2/tempconv"
// Then do: $ run ex_2_1.go
package main

import (
	"fmt"
	"github.com/hoohawa/go_exercises/ch2/tempconv"
)

func main() {
	var tempC tempconv.Celsius
	fmt.Print("Input temperature in Celsius: ")
	fmt.Scanf("%g", &tempC)
	fmt.Printf("Converted: %.2f F, %.2fK\n", tempconv.CToF(tempC), tempconv.CToK(tempC))

	var tempF tempconv.Fahrenheit
	fmt.Print("Input temperature in Fahrenheit: ")
	fmt.Scanf("%g", &tempF)
	fmt.Printf("Converted: %.2f C, %.2fK\n", tempconv.FToC(tempF), tempconv.FToK(tempF))

	var tempK tempconv.Kelvin
	fmt.Print("Input temperature in Kelvin: ")
	fmt.Scanf("%g", &tempK)
	fmt.Printf("Converted: %.2fC, %.2fF\n", tempconv.KToC(tempK), tempconv.KToF(tempK))
}
