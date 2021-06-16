package utils

/*SliceToChunk [][]int*/
func SliceToChunk(data []int, chunkSize int) [][]int {
	if len(data) == 0 {
		return nil
	}
	divided := make([][]int, (len(data)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(data) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = data[prev:next]
		prev = next
		i++
	}
	divided[i] = data[prev:]
	return divided
}
