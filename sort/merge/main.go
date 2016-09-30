package main

func merge(store []int64, t1 []int64, t2 []int64) {

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

func mSort(store []int64, left int, right int) int {

	if left == right {
		return 0
	} 
	
	var mid int = (right - left) / 2
	leftPt := left
	rightPt := mid + 1

	temp1 := make([]int64, mid + 1)
	for i:=0; i < len(temp1); i++ {
		temp1[i] = store[i]
	}
	mSort(temp1, 0, len(temp1) - 1)

	temp2 := make([]int64, len(store) - rightPt)
	for i:=0; i < len(temp2); i++ {
		temp2[i] = store[mid + 1 +i]
	}
	
	mSort(temp2, 0, len(temp2) - 1)

	temp := make([]int64, len(store))

	merge(temp, temp1, temp2)

	for i:=0; i < len(temp); i++ {
		store[i] = temp[i]
	}

	return 0
}