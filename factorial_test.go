package math

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want big.Int
	}{
		{
			name: "Factorial 4! = 24",
			args: args{
				n: 4,
			},
			want: *big.NewInt(24),
		},
		{
			name: "Factorial 6! = 720",
			args: args{
				n: 6,
			},
			want: *big.NewInt(720),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
