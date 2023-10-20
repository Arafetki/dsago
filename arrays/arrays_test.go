package arrays

import "testing"

func TestLinearSearch(t *testing.T) {
	t.Parallel()
	arr := []int{1, 8, 2, 6, 3, 25}
	type testCase struct {
		el   int
		want bool
	}
	testCases := []testCase{
		{el: 6, want: true},
		{el: 3, want: true},
		{el: 50, want: false},
	}

	for _, tc := range testCases {
		got := LinearSearch(arr, tc.el)
		if got != tc.want {
			t.Errorf("want %v got %v", tc.want, got)
		}
	}
}

func TestBinarySearch(T *testing.T) {
	// Test Binary Search
}
