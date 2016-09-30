package main

func SelectionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {

		min := arr[i+1]
		min_idx := i + 1

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				min_idx = j
			}
		}

		if min < arr[i] {
			temp := arr[i]
			arr[min_idx] = temp
			arr[i] = min
		}
	}
}