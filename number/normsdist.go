package number_helper

import (
	"math"
)

//标准正态分布的累积密度函数
func CalcNormsdist(a float64) float64 {
	var p float64 = 0.2316419
	var b1 float64 = 0.31938153
	var b2 float64 = -0.356563782
	var b3 float64 = 1.781477937
	var b4 float64 = -1.821255978
	var b5 float64 = 1.330274429
	var x = math.Abs(a)

	var t = 1 / (1 + p*x)
	var val = 1 - (1/(math.Sqrt(2*math.Pi))*math.Exp(-1*math.Pow(a, 2)/2))*(b1*t+b2*math.Pow(t, 2)+b3*math.Pow(t, 3)+b4*math.Pow(t, 4)+b5*math.Pow(t, 5))
	if a < 0 {
		val = 1 - val
	}
	return val
}
