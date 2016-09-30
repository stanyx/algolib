package main

import (
	"math/rand"
)

func RandSequence(n int) ([]int64) {
	//generated n random array
	arr1 := make([]int64, n)
	arr2 := make([]int64, n)

	for i:=0; i < n; i++ {
		arr1[i] = int64(i + 1)
	}

	seed := rand.NewSource(time.Now().UnixNano())
    r := rand.New(seed)

	for i:=0; i < n; i++ {
		idx := r.Intn(len(arr1))

		arr2[i] = arr1[idx]
		arr1_l := arr1[0:idx]
		arr1_r := arr1[idx + 1: len(arr1)]

		arr1 = make([]int64, len(arr1_l) + len(arr1_r))
		c:=0
		for j := 0; j < len(arr1); j++ {

			if j < len(arr1_l) {
				arr1[j] = arr1_l[j]
			} else {
				arr1[j] = arr1_r[c]
				c += 1
			}

		}

	}

	//fmt.Printf("rand arr -> %s\r\n", arr2)
	return arr2
}

func getMedianOf3(arr []int64, left int, right int) {
	//sort by median
	if len(arr) > 3 {

		mid := (right - left) / 2

		if arr[left] > arr[mid] {
			temp := arr[left]
			arr[left] = arr[mid]
			arr[mid] = temp
		}
		if arr[mid] > arr[right] {
			temp := arr[right]
			arr[right] = arr[mid]
			arr[mid] = temp
		}
		if arr[left] > arr[right] {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
		}

	}
}

func partition(arr []int64, left int, right int, pivot int64) (int) {

	var leftPt int
	leftPt = left - 1
	rightPt := right

	for {

		for {		
			leftPt += 1
			if (arr[leftPt] >= pivot) {
				break
			}
		}

		for {
			rightPt -= 1
			if (rightPt <= 0 || arr[rightPt] <= pivot) {
				break
			}
		}

		if (leftPt >= rightPt) {			
			break
		} else {
			//swap elements
			temp_inner := arr[rightPt]
			arr[rightPt] = arr[leftPt]
			arr[leftPt] = temp_inner
		}

	}

	//swap last elements
	temp := arr[leftPt]
	arr[leftPt] = arr[right]
	arr[right] = temp
	return leftPt
}

func qsortWithInsert(arr []int64, left int, right int) (int) {
	return 0
}

func qsort(arr []int64, left int, right int) (int) {

	if right - left <= 0 {
		return 0
	} else {

		pivot := arr[right]
		left_idx := partition(arr, left, right, pivot)

		qsort(arr, left, left_idx - 1)
		qsort(arr, left_idx + 1, right)
	}

	return 0
}


func QuickSort(arr []int64, merge bool) {

	if merge {
		qsortWithInsert(arr, 0, len(arr) - 1)
	} else {
		qsort(arr, 0, len(arr) - 1)
	}

}