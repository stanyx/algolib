package main

func OddEvenSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		a := 0 
		if i % 2 != 0 {
			a = 1
		}
		for j := a; j < len(arr) - 1; j+=2 {
			if arr[j] > arr[j + 1] {
				temp := arr[j+1]
				arr[j + 1] = arr[j]
				arr[j] = temp
			}
		}
	}
}