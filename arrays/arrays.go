package arrays

// Search algorithms

// @ Linear Search -- O(N) Time Complexity (Worst Case) -- N=len(arr)

func LinearSearch[T comparable](arr []T, el T) (T, bool) {
	var res T
	for _, val := range arr {
		if val == el {
			return val, true
		}
	}
	return res, false
}
