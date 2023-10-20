package arrays

import (
	"golang.org/x/exp/constraints"
)

// Search algorithms

// @ Linear Search -- O(N) Time Complexity (Worst Case) -- N=len(arr)

func LinearSearch[T constraints.Ordered](arr []T, el T) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
}

// @ Binary Search -- O(logN) Time Complexity (Worst Case) -- N=len(arr)

func BinarySearch[T constraints.Ordered](arr []T, l, r int, el T) bool {
	for l <= r {
		mid := l + (r-l)/2
		if el == arr[mid] {
			return true
		}
		if el < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return false
}
