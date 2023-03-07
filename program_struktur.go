package main

/*import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF)) // "212°F = 100°C"
	}
	func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
	}
	*/

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC Celcius = 0
	BoilingC Celcius = 100

)

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 320) }
func FToc(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }


