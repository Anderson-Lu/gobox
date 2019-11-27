package application

import (
	"testing"
)

func BenchmarkTypeConvert1(b *testing.B) {
	var d interface{} = 1
	var total = 0
	b.StartTimer()
	for i := 0; i < 1000000; i++ {
		total += d.(int)
	}

	b.StopTimer()
}

func BenchmarkTypeConvert2(b *testing.B) {
	var d int = 1
	var total = 0
	b.StartTimer()
	for i := 0; i < 1000000; i++ {
		total += d
	}

	b.StopTimer()
}
