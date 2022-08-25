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
	fmt.Printf("°C to °F: %v °F\n", thc.CtoF(30))
	fmt.Printf("DewPoint: %v °C\n", thc.DewPoint(35, 85))
	fmt.Printf("WetBulb: %v(°C)\n", thc.WetBulb(35, 85))
	fmt.Printf("Absolute Humidity: %v(KG_Water/KG_DryAir)\n", thc.AH(35, 85))
	fmt.Printf("Enthalpy: %v(kJ/kg):\n", thc.Enth(35, 85))
	thi, msg := thc.THI(35, 85)
	fmt.Printf("THI: %v(%s)", thi, msg)
}
```
Hope you like it.