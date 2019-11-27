package number

import "testing"

type D struct {
	Num float64
}

type demo []D

func (s demo) Len() int {
	return len(s)
}

func (s demo) Get(i int) float64 {
	return s[i].Num
}

func TestCalcAverage(t *testing.T) {

	s := demo{D{Num: 1.0}, D{Num: 2.0}, D{Num: 3.0}, D{Num: 4.0}, D{Num: 5.0}}

	type args struct {
		data Avg
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "计算平均数",
			args: args{
				data: s,
			},
			want: 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcAverage(tt.args.data); got != tt.want {
				t.Errorf("CalcAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcVariance(t *testing.T) {
	s := demo{D{Num: 1.0}, D{Num: 2.0}, D{Num: 3.0}, D{Num: 4.0}, D{Num: 5.0}}
	type args struct {
		data Variance
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "计算方差",
			args: args{
				data: s,
			},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcVariance(tt.args.data); got != tt.want {
				t.Errorf("CalcVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
