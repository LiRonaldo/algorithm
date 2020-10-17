package main

import "fmt"

/**
优化版
*/
func main() {
	arr := [10]int{2, 1, 5, 8, 1, 3, 48, 1, 6, 10}
	sort2(&arr, 10)
	fmt.Println(arr)
}

func sort2(arr *[10]int, n int) {
	left := 0
	right := n - 1
	for left < right {
		minindex := left
		maxindex := right

		if arr[minindex] > arr[maxindex] {
			arr[minindex], arr[maxindex] = arr[maxindex], arr[minindex]
		}

		for i := left + 1; i < right; i++ {
			if arr[i] < arr[minindex] {
				minindex = i
			} else if arr[i] > arr[maxindex] {
				maxindex = i
			}
		}

		arr[minindex], arr[left] = arr[left], arr[minindex]
		arr[maxindex], arr[right] = arr[right], arr[maxindex]
		left++
		right--
	}
}
