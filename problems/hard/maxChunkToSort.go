// https://leetcode.com/problems/max-chunks-to-make-sorted-ii/
// if there could be chunk, the max number in last chunk should be
// smaller than the min number in next chunk
// so we could use the thought to resolve the question
// Basiclly, it's a DP
func maxChunksToSorted(arr []int) int {
	var chunkNumber int = 0
	var length int = len(arr)
	//DP state
	// the last chunk slice, lastChunk[i] means the max number in first i numbers
	lastChunk := make([]int, length)
	// the next chunk slice,nextChunk[i] means the min number in next i numbers
	nextChunk := make([]int, length)
	//DP init
	lastChunk[0] = arr[0]
	nextChunk[length-1] = arr[length-1]
	// DP function
	// init the lastChunk
	for i := 1; i < length; i++ {
		lastChunk[i] = max(arr[i], lastChunk[i-1])
	}
	// init the nextChunk
	for i := length - 2; i >= 0; i-- {
		nextChunk[i] = min(arr[i], nextChunk[i+1])
	}
	// result
	for i := 0; i < length-1; i++ {
		if lastChunk[i] <= nextChunk[i+1] {
			chunkNumber++
		}
	}
	return chunkNumber + 1
}
func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
