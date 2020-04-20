package insertion

import "fmt"

// InsertionSort sorts the input slice of integers in ascending order
// Worst-case complexity is O(N^2)
// Best-case complexity is O(N)
func InsertionSort(A []int) {
	fmt.Println("INSERTION SORT ALGORITHM")
	for i := 1; i < len(A); i++ {
		key := A[i]
		insertionIndex := i
		for j := i - 1; j >= 0 && A[j] > key; j-- {
			A[j+1] = A[j]
			insertionIndex = j
		}
		A[insertionIndex] = key
	}
}
