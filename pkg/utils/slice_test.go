package utils

import (
	"reflect"
	"testing"
)

func TestSliceToChunk(t *testing.T) {
	type args struct {
		data      []int
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Chunks",
			args: args{
				data:      []int{1, 2, 3, 4, 5, 6, 7},
				chunkSize: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToChunk(tt.args.data, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
