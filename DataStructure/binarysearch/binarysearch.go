package binarysearch

func binarysearch(array []int, targetNum int, start int, end int) int {

	if start > end {
		return -1
	}

	midIdx := (start + end) / 2
	midVal := array[midIdx]

	if targetNum < midVal {
		return binarysearch(array, targetNum, start, midIdx)
	} else if targetNum > midVal {
		return binarysearch(array, targetNum, midIdx+1, end)
	} else {
		return midIdx
	}

}
