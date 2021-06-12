package sort

/*Permutations []int*/
func Merge(input []int) []int {
	length := len(input)
	if length < 2 {
		return input
	}
	mid := length / 2
	return internalMerge(Merge(input[:mid]), Merge(input[mid:]))
}

func internalMerge(left, right []int) []int {
	result := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}
