package number_helper

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

/*
* Return random int value between 0 and max-1
 */
func RandomInt(max int) int {
	return rand.Intn(max)
}

/*
* Return the min value between a1 and a2
 */
func Min(a1, a2 float64) float64 {
	if a1 < a2 {
		return a1
	}
	return a2
}

/*
* Return the min value between a1 to aN
 */
func MinN(aN ...int) int {
	ret := math.MaxInt16
	for _, v := range aN {
		if v < ret {
			ret = v
		}
	}
	return ret
}

/*
* Return the min value between a1 to aN
 */
func MinFloat64N(aN ...float64) float64 {
	ret := math.MaxFloat64
	for _, v := range aN {
		if v < ret {
			ret = v
		}
	}
	return ret
}

//向上或者向下按照有效位数取整
//n为有效位数
//raw为原始数字
//isUp是否向上取整,否则为向下取整
func FloorOrCeil(n int, raw float64, isUp bool) float64 {

	if raw >= math.Pow(10.0, float64(n)) {
		return raw
	}

	t := strconv.FormatFloat(raw, 'G', n, 64)
	ret, _ := strconv.ParseFloat(t, 64)
	t1 := strconv.FormatFloat(raw, 'G', n+1, 64)
	ret1, _ := strconv.ParseFloat(t1, 64)
	pointIndex := strings.Index(t1, ".")

	lastCh, _ := strconv.ParseInt(t1[len(t1)-1:], 10, 64)
	isNotZero := (ret1-ret != 0)

	if !isNotZero {
		return ret
	}

	points := len(t1) - pointIndex - 1

	//整数位数大于小数位数
	if pointIndex > points && pointIndex == n && points != 1 {

		if isUp && lastCh < 5 {
			ret += 1.0
		}

		if isNotZero && lastCh >= 5 && !isUp && points != 0 {
			ret -= 1.0
		}
	} else if pointIndex > points && pointIndex == n && points == 1 {
		if lastCh <= 5 && isUp {
			ret += 1.0
		}
		if lastCh > 5 && !isUp {
			ret -= 1.0
		}
	} else {

		fator := math.Pow(10, float64(-points+1))

		if isUp {

			if isNotZero && lastCh < 5 {
				ret += fator
			} else if isNotZero && lastCh <= 5 && points == 2 {
				ret += fator
			}
		}

		if isNotZero && lastCh >= 5 && !isUp && points != 2 {
			ret -= fator
		}
	}

	tmp := strconv.FormatFloat(ret, 'G', n, 64) //3.3343
	ret, _ = strconv.ParseFloat(tmp, 64)
	return ret
}

//判断有多少位精度
func CalcDigist(n float64) int {
	switch n {
	case 0.1:
		return 1
	case 0.01:
		return 2
	case 0.001:
		return 3
	case 0.0001:
		return 4
	case 0.00001:
		return 5
	case 0.000001:
		return 6
	case 0.0000001:
		return 7
	case 0.00000001:
		return 8
	default:
		return -1
	}
}

//按照指定精度位数向上或者向下取整
func Round(f float64, n int) float64 {
	var i int = 1
	var count int = 0
	for {
		if int(f/float64(i)) > 0 {
			i *= 10
			count++
			continue
		}
		break
	}
	return FloorOrCeil(count+n, f, false)
}

func f(x, a float64) float64 {
	return math.Exp(x) - a
}

func Ln(n float64) float64 {
	var lo, hi, m float64
	if n <= 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	EPS := 0.00001
	lo = 0
	hi = n
	for math.Abs(lo-hi) >= EPS {
		m = float64((lo + hi) / 2.0)
		if f(m, n) < 0 {
			lo = m
		} else {
			hi = m
		}
	}
	return float64((lo + hi) / 2.0)
}
