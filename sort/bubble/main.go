package bubble

func BubbleSortV1(A []int) {
	for i := 0; i < len(A) - 1; i++ {
		for j := 0; j < len(A) - 1 - i; j++ {
			if A[j] > A[j + 1] {
				tmp := A[j]
				A[j] = A[j+1]
				A[j+1] = tmp
			}
		}
	}
}

func BubbleSortV2(arr []int) {

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j + 1] {
				temp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}
}
