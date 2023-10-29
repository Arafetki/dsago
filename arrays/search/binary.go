package search

// Binary Search - N = len(array)
// Time Complexity : O(logN) -- Worst Case
// Array Must Be Sorted

func Binary(arr []int, low int, high int, el int) (int, bool) {

	for low <= high {
		middle := low + (high-low)/2
		if el == arr[middle] {
			return middle, true
		} else if el < arr[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1, false
}
