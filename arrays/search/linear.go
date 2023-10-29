package search

// Linear Search - N = len(array)
// Time Complexity : O(N) -- Worst Case

func Linear(arr []int, el int) (int, bool) {
	for i, v := range arr {
		if v == el {
			return i, true
		}
	}
	return -1, false
}
