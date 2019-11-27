package number

import "math"

type Variance interface {
	Len() int
	Get(i int) float64
}
type Avg Variance

//计算方差
func CalcVariance(data Variance) float64 {
	n := data.Len()
	if n == 0 {
		return 0.0
	}
	avg := CalcAverage(data.(Avg))
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += math.Pow((data.Get(i) - avg), 2)
	}
	return sum / float64(n)
}

//计算平均数
func CalcAverage(data Avg) float64 {
	n := data.Len()
	if n == 0 {
		return 0.0
	}
	x := 0.0
	for i := 0; i < n; i++ {
		x += data.Get(i)
	}
	return x / float64(n)
}
