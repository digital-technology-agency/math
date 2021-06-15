package combinatorics

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
		want int64
	}{
		{
			name: "Factorial 10! = 24",
			args: args{
				n: 10,
			},
			want: 3628800,
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

func TestFactorialTree(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "1000000!",
			args: args{
				n: 1000000,
			},
			want: big.NewInt(3628800),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialTree(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FactorialTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorialTreeString(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1000000!",
			args: args{
				n: 1000000,
			},
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialTreeString(tt.args.n); got != tt.want {
				t.Errorf("FactorialTreeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
