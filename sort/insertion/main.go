package insertion


func InsertionSort(A []int) {

	for i := 1; i < len(A); i++ {
		key := A[i]
		lst := i
		for j := i - 1; j > -1; j-- {
			if A[j] > A[i] {
				A[j + 1] = A[j]
				lst = j
			} else {
				break
			}
		}
		A[lst] = A[i]
	}
}
