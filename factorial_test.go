package math

import (
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
			name: "Factorial 4! = 24",
			args: args{
				n: 4,
			},
			want: 24,
		},
		{
			name: "Factorial 6! = 720",
			args: args{
				n: 6,
			},
			want: 720,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialInt(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
