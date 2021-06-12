package combinatorics

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Factorial 4! = 24",
			args: args{
				n: 399168,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != nil {
				t.Errorf("Factorial() = want %v", tt.want)
			}
		})
	}
}
