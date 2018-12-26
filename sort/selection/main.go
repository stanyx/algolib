package main

func SelectionSort(A []int) {
	for i := 0; i < len(A); i++ {

		min := i

		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min] {
				min = j
			}
		}

		if A[min] < A[i] {
			temp := A[i]
			A[i] = A[min]
			A[min] = temp
		}
	}
}
