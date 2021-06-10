package combinatorics

import (
	"reflect"
	"testing"
)

func TestPlacement(t *testing.T) {
	type args struct {
		n int64
		k int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Placement N=10 K=10",
			args: args{
				n: 10,
				k: 10,
			},
			want: `3628800`,
		},
		{
			name: "Placement N=25 K=25",
			args: args{
				n: 25,
				k: 25,
			},
			want: `15511210043330985984000000`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlacementStr(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Placement() = %v, want %v", got, tt.want)
			}
		})
	}
}
