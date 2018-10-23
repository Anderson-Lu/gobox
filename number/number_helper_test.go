package number_helper

import (
	"fmt"
	"math"
	"testing"
)

func TestCalcDigist(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "测试精度获取1", args: args{n: 0.1}, want: 1},
		{name: "测试精度获取2", args: args{n: 0.01}, want: 2},
		{name: "测试精度获取2", args: args{n: 0.001}, want: 3},
		{name: "测试精度获取2", args: args{n: 0.0001}, want: 4},
		{name: "测试精度获取2", args: args{n: 0.00001}, want: 5},
		{name: "测试精度获取2", args: args{n: 0.000001}, want: 6},
		{name: "测试精度获取2", args: args{n: 0.00000000001}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcDigist(tt.args.n); got != tt.want {
				t.Errorf("CalcDigist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLn(t *testing.T) {
	if math.Abs(Ln(10)-2.302585) > 0.000001 {
		t.Errorf("clac error,result=%f,want=%f", Ln(10), 2.302585)
	}
}

func TestRound(t *testing.T) {
	type args struct {
		f float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "测试", args: args{f: 789.9999, n: 3}, want: 789.999},
		{name: "测试", args: args{f: 0.9999, n: 3}, want: 0.999},
		{name: "测试", args: args{f: 0.00019, n: 3}, want: 0.00019},
		{name: "测试", args: args{f: 188888, n: 3}, want: 188888},
		{name: "测试", args: args{f: 18.001, n: 3}, want: 18.001},
		{name: "测试", args: args{f: 18.0012, n: 3}, want: 18.001},
		{name: "测试", args: args{f: 18.0015, n: 3}, want: 18.001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.f, tt.args.n); got != tt.want {
				fmt.Println(got)
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
