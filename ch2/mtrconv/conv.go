package mtrconv

import "fmt"

// Temperature types
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// Length types
type Meter float64
type Mile float64
type Inch float64

//Weight types
type Pound float64
type Gram float64
type Ounce float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (mi Mile) String() string { return fmt.Sprintf("%g mi", mi) }
func (in Inch) String() string { return fmt.Sprintf("%g in", in) }

func (p Pound) String() string { return fmt.Sprintf("%g lbs", p) }
func (g Gram) String() string  { return fmt.Sprintf("%g g", g) }
func (o Ounce) String() string { return fmt.Sprintf("%g oz", o) }
