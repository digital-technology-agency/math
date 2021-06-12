package sort

/*Insertion []int*/
func Insertion(input []int) []int {
	for j := 1; j < len(input); j++ {
		key := input[j]
		i := j - 1
		for input[i] > key {
			input[i+1] = input[i]
			i -= 1
			if i < 0 {
				break
			}
		}
		input[i+1] = key

	}
	return input
}
