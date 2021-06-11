package sort

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Merge sort",
			args: args{
				size: 100,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputArray := randomArray(tt.args.size)
			got := Merge(inputArray)
			log.Printf(`Result: %v`, got)
			if got[0] < tt.want {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
