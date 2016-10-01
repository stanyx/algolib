package main

func CocktailSort(A []int) []int {

	left := 0
	right := len(A) - 1
	swaps := 0
	iters := 0
	
	for {

		j := left

		for {

			if j >= right {
				break
			}

			if (A[j] > A[j + 1]) {
				t := A[j + 1]
				A[j + 1] = A[j]
				A[j] = t
				swaps += 1
			}

			j++
			iters += 1
		}

		i := right

		for {

			if i <= left {
				break
			}

			if (A[i] < A[i - 1]) {
				t := A[i - 1]
				A[i - 1] = A[i]
				A[i] = t
				swaps += 1
			}
			
			i--
			iters += 1
		}

		left += 1
		right -= 1

		if left > right {
			break
		}
	}

	return A
}