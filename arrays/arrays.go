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

// @ Binary Search -- O(logN) Time Complexity (Worst Case) -- N=len(arr)

func BinarySearch(arr []int, l, r int, el int) (int, bool) {
	for l <= r {
		mid := l + (r-l)/2
		if el == arr[mid] {
			return mid, true
		}
		if el < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return -1, false
}
