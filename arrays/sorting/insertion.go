package sorting

func InsertionSort(arr []int) {
	n := len(arr)
	var (
		key int
		j   int
	)
	for i := 1; i < n; i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
