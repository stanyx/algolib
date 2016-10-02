package merge

import "testing"

func TestMergeSort(t *testing.T) {
	E := []int{1, 2, 3, 4, 5}
	A := []int{5, 4, 3, 2, 1}
	MergeSort(A, 0, len(A))
	for i, a := range(A) {
		if a != E[i] {
			t.Error("Expected ", E[i], "got ", a)
		}
	}
}