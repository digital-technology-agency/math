package sort

import (
	"crypto/rand"
	"log"
	"math/big"
	"testing"
)

var (
	resultSorted = []int{1, 2, 3, 4, 5, 6}
)

func TestInsertion(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sort by insertion",
			args: args{
				size: 100,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputArray := randomArray(tt.args.size)
			got := Insertion(inputArray)
			log.Printf(`Result: %v`, got)
			if got[0] < tt.want {
				t.Errorf("Insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func randomArray(size int) []int {
	result := []int{}
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
	total := n.Int64()
	for total > 0 {
		value, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
		result = append(result, int(value.Int64()))
		total--
	}
	return result
}
