package merge

func merge(store []int, t1 []int, t2 []int) {

	arrLeftPt := 0
	arrRightPt := 0

	pt := 0

	for {

		if arrLeftPt < len(t1) || arrRightPt < len(t2) {

			if t1[arrLeftPt] < t2[arrRightPt] {
				store[pt] = t1[arrLeftPt]
				arrLeftPt += 1
				pt += 1
			} else {	
				store[pt] = t2[arrRightPt]
				arrRightPt += 1
				pt += 1
			}

			if arrLeftPt == len(t1) {
				for i := arrRightPt; i < len(t2); i++ { 
					store[pt] = t2[i]
					pt +=1					
				}
				break
			}

			if arrRightPt == len(t2) {

				for i := arrLeftPt; i < len(t1); i++ {
					store[pt] = t1[i] 
					pt += 1					
				}
				break
			}
		} else {
			break
		}

	}

}

func MergeSort(A []int, left int, right int) {

	if left == right {
		return
	} 
	
	var mid int = (right - left) / 2
	rightPt := mid + 1

	temp1 := make([]int, rightPt)
	for i := 0; i < len(temp1); i++ {
		temp1[i] = A[i]
	}
	
	MergeSort(temp1, 0, len(temp1) - 1)

	temp2 := make([]int, len(A) - rightPt)
	for i := 0; i < len(temp2); i++ {
		temp2[i] = A[rightPt + i]
	}
	
	MergeSort(temp2, 0, len(temp2) - 1)

	temp := make([]int, len(A))

	merge(temp, temp1, temp2)

	for i := 0; i < len(temp); i++ {
		A[i] = temp[i]
	}

}