package main

import "fmt"

/**
希尔排序 ,就是分组的插入排序。只要掌握插入排序。然后 i++的时候要i+分组的
*/
func main() {
	arr := []int{7, 3, 4, 8, 18, 1}
	sort(arr)
	fmt.Println(arr)
}
func sort(arr []int) {
	n := len(arr)
	for n > 1 {
		n = n / 2
		for i := 0; i < n; i++ {
			for j := n; j < len(arr); j = j + n {
				var temp = arr[j]
				var k int
				for k = j; k > 0 && arr[k-n] > temp; k = k - n {
					arr[k] = arr[k-n]
				}
				arr[k] = temp
			}
		}
	}
}
