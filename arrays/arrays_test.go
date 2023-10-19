package arrays

import "testing"

func TestLinearSearch(t *testing.T) {
	t.Parallel()
	type testCase struct {
		arr  []int
		el   int
		want bool
	}

	testCases := []testCase{
		{arr: []int{1, 8, 2, 6, 3, 25}, el: 6, want: true},
		{arr: []int{3, 9, 2, 6, 16, 25}, el: 3, want: true},
		{arr: []int{1, 25, 2, 6, 3, 49}, el: 50, want: false},
	}

	for _, tc := range testCases {
		_, got := LinearSearch(tc.arr, tc.el)
		if got != tc.want {
			t.Errorf("want %v got %v", tc.want, got)
		}
	}
}

func TestBinarySearch(T *testing.T) {
	// Test Binary Search
}
