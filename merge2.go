package main

import (
	"fmt"
)

func Merge(A []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q
	var L []int
	var R []int
	var k int = p
	var i int
	var j int
	for i = 0; i < n1; i++ {
		L = append(L, A[k])
		k += 1
	}
	for j = 0; j < n2; j++ {
		R = append(R, A[k])

		k += 1
	}
	L = append(L, 99999999)
	R = append(R, 99999999)
	k = p
	i = 0
	j = 0
	for k = p; k <= r; k++ {
		if L[i] >= R[j] {
			A[k] = R[j]
			j += 1

		} else {
			A[k] = L[i]
			i += 1
		}

	}
	//fmt.Println(L, R)
}
func MergeSort(A []int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q, r)

	}

}

func main() {
	a := []int{10, 5, 7, 9, 7, 5, 6, 2, 3, 1}
	MergeSort(a, 0, 9)
	fmt.Println(a)

}
