package main


func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		pick := arr[i]
		lst := i
		for j:= i-1; j > -1; j-- {
			if arr[j] > pick {
				arr[j + 1] = arr[j]
				lst = j
			} else {
				break
			}
		}
		arr[lst] = pick
	}
}