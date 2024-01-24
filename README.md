# thc: Temperature Humidity Calculator
Temperature Humidity Calculator 

# Quickly start

## install

    go get -u github.com/ka1hung/thc

## example
```go
package main

import (
	"fmt"
	"github.com/ka1hung/thc"
)

func main() {
	fmt.Printf("°F to °C: %v °C\n", thc.FtoC(100)) 
	//°F to °C: 37.78 °C

	fmt.Printf("°C to °F: %v °F\n", thc.CtoF(30)) 
	//°C to °F: 86 °F

	fmt.Printf("DewPoint: %v °C\n", thc.DewPoint(35, 85)) 
	//DewPoint: 32.1 °C

	fmt.Printf("WetBulb: %v(°C)\n", thc.WetBulb(35, 85)) 
	//WetBulb: 35.37(°C)

	fmt.Printf("Absolute Humidity: %v(g/kg(a))\n", thc.AH(35, 85)) 
	//Absolute Humidity: 30.79(g/kg(a))

	fmt.Printf("Enthalpy: %v(kJ/kg)\n", thc.Enth(35, 85))
	//Enthalpy: 114.13(kJ/kg)

	thi, msg := thc.THI(35, 85)
	fmt.Printf("THI: %v(%s)", thi, msg) 
	//THI: 33(Muggy)
}
```
Hope you like it.

ref: https://www.buildenvi.com/gongju/psychrometrics  