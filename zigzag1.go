package main

import "fmt"

func printZigZag(a [][]int, n int) {
	var i, j, r, c int
	var count, f int = 0, 0
	for i = 0; i < n; i++ {
		if f == 0 {
			r = i
			c = 0
			for r >= 0 && c <= i {
				fmt.Print(a[r][c], " ")
				r -= 1
				c += 1
				count += 1
				if count == n {
					fmt.Println()
					count = 0
				}
			}
			f = 1

		} else {
			r = 0
			c = i
			for r <= i && c >= 0 {
				fmt.Print(a[r][c], " ")
				r += 1
				c -= 1
				count += 1
				if count == n {
					fmt.Println()
					count = 0
				}
			}
			f = 0
		}
	}
	for j = 1; j < n; j++ {
		if f == 0 {
			r = n - 1
			c = j
			for r >= j && c <= n-1 {
				fmt.Print(a[r][c], " ")
				r -= 1
				c += 1
				count += 1
				if count == n {
					fmt.Println()
					count = 0
				}
				f = 1
			}

		} else {
			r = j
			c = n - 1
			for r <= n-1 && c >= j {
				fmt.Print(a[r][c], " ")
				r += 1
				c -= 1
				count += 1
				if count == n {
					fmt.Println()
					count = 0
				}
			}
			f = 0

		}
	}
}

func main() {
	// a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	printZigZag(a, 3)
}
