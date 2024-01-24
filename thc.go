package thc

import (
	"math"

	"github.com/shopspring/decimal"
)

// Psychrometric Chart
// https://upload.wikimedia.org/wikipedia/commons/9/9d/PsychrometricChart.SeaLevel.SI.svg

// setting for DecimalPlaces of each calculate result, deflaut is 2
var DecimalPlaces int32 = 2

var pvs []float64 = []float64{
	611.9, 657.7, 706.6, 758.6, 814.0,
	873.0, 935.7, 1002.4, 1073.2, 1148.4,
	1228.2, 1312.9, 1402.7, 1497.8, 1598.5,
	1705.2, 1818, 1937.4, 2063.6, 2196.9,
	2337.7, 2486.4, 2643.3, 2808.7, 2983.2,
	3167, 3360.6, 3564.5, 3779.1, 4004.8,
	4242.2, 4491.8, 4754, 5029.4, 5318.5,
	5621.9, 5940.3, 6274.1, 6624, 6990.8,
	7374.9, 7777.1, 8198.1, 8638.7, 9099.5,
	9581.3, 10085, 10611.2, 11160.9, 11734.8,
	12334, 12959.1, 13611.3, 14291.4, 15000.3,
	15739.1, 16508.9, 17310.5, 18145.2, 19013.9,
	19917.9, 20858.2, 21836, 22852.6, 23909.1,
	25006.8, 26147.1, 27331.1, 28560.3, 29836,
	31159.6, 32532.5, 33956.3, 35432.3, 36962.1,
	38547.2, 40189.3, 41889.9, 43650.7, 45473.3,
	47359.5, 49310.9, 51329.4, 53416.8, 55574.8,
	57805.4, 60110.4, 62491.9, 64951.6, 67491.7,
	70114.2, 72821.2, 75614.6, 78496.8, 81469.8,
	84535.8, 87697.2, 90956.1, 94314.9, 97775.9,
	101341.5}

// CtoF convert °C to °F
func CtoF(t float64) float64 {
	v := t*9/5 + 32
	f, _ := decimal.NewFromFloat(v).Round(DecimalPlaces).Float64()
	return f
}

// FtoC convert °F to °C
func FtoC(t float64) float64 {
	v := (t - 32) / 9 * 5
	f, _ := decimal.NewFromFloat(v).Round(DecimalPlaces).Float64()
	return f
}

// Dew Point Temperature calculate
// in:  t = Temperature(°C); h = Relative Humidity(%);
// out: DewPoint(°C)
func DewPoint(t, h float64) float64 {
	k := (math.Log10(h)-2)/0.4343 + (17.62*t)/(243.12+t)
	v := 243.12 * k / (17.62 - k)
	f, _ := decimal.NewFromFloat(v).Round(DecimalPlaces).Float64()
	return f
}

// Wet Bulb Temperature calculate
// in: t = Temperature(°C); h = Relative Humidity(%);
// out: WetBulb(°C)
func WetBulb(t, h float64) float64 {
	wet_bulb := (-5.806 + 0.672*t - 0.006*t*t) + (0.061+0.004*t+0.000099*t*t)*h + (-0.000033-0.000005*t-0.0000001*t*t)*h*h
	f, _ := decimal.NewFromFloat(wet_bulb).Round(DecimalPlaces).Float64()
	return f
}

// Absolute Humidity calculate
// in: t = Temperature(°C); h = Relative Humidity(%);
// out: AH (KG_Water/KG_DryAir)
func AH(t, h float64) float64 {
	pv := h * pvs[int(t)] / 100
	w := (0.622 * pv) / (101325 - pv)
	f, _ := decimal.NewFromFloat(w * 1000).Round(DecimalPlaces).Float64()
	return f
}

// Enthalpy calculate
// in: t = Temperature(°C); h = Relative Humidity(%);
// out: Enth(kJ/kg)
func Enth(t, h float64) float64 {
	Enth := (1.005 * t) + ((AH(t, h) * 0.001) * (1.805*t + 2501))
	f, _ := decimal.NewFromFloat(Enth).Round(DecimalPlaces).Float64()
	return f
}

// Temperature Humidity Index calculate (Comfortable Index)
// in: t = Temperature(°C); h = Relative Humidity(%);
// out: index, msg
//
//	THI<=10		-> very cold
//
// 11<=THI<=15		-> cold
// 16<=THI<=19 	-> Slightly cold
// 20<=THI<=26 	-> comfortable
// 27<=THI<=30 	-> muggy
// 31<=THI 		-> danger
func THI(t, h float64) (int, string) {
	dp := DewPoint(t, h)
	thi := t - 0.55*(1-(math.Exp((17.269*dp)/(dp+237.3)))/math.Exp(((17.269*t)/(t+237.3))))*(t-14)
	index := int(math.Round(thi))
	msg := ""
	if index <= 10 {
		msg = "Very cold"
	} else if index >= 11 && index <= 15 {
		msg = "Cold"
	} else if index >= 16 && index <= 19 {
		msg = "Slightly cold"
	} else if index >= 20 && index <= 26 {
		msg = "Comfortable"
	} else if index > 27 && index <= 30 {
		msg = "Muggy"
	} else if index >= 31 {
		msg = "Muggy"
	}
	return index, msg
}
