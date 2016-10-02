package insertion

import "testing"

func TestInsertionSort(t *testing.T) {
	A := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	E := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	InsertionSort(A)
	for i, a := range(A) {
		if a != E[i] {
			t.Error("Expected ", E[i], "got ", a)
		}
	}
}