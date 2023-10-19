package arrays

// Search algorithms

// @ Linear Search -- O(N) Time Complexity (Worst Case) -- N=len(arr)

func LinearSearch[T comparable](arr []T, el T) (int, bool) {
	for i, val := range arr {
		if val == el {
			return i, true
		}
	}
	return -1, false
}
