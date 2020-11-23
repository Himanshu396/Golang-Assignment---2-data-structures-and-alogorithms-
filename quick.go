package main

import (
	"fmt"
)

func Partition(A []int, p int, r int) int {
	x := A[r]
	i := p - 1
	var temp int
	for j := p; j <= r-1; j++ {
		if A[j] <= x {
			i += 1
			temp = A[i]
			A[i] = A[j]
			A[j] = temp
		}
		
	}
	temp = A[i+1]
	A[i+1] = A[r]
	A[r] = temp
	//fmt.Println(A, i+1)
	return i + 1
}

//go run QuickSort.go
func QuickSort(A []int, p int, r int) {
	if p < r {
		q := Partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)

	}

}

func main() {
	a := []int{10, 5, 7, 9, 7, 5, 6, 2, 3, 1}
	QuickSort(a, 0, 9)
	fmt.Println(a)

}