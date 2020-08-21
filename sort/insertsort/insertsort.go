package main

import "fmt"

/**
插入排序  将待排序的数组分成2分,一份有序，一份无序,每次都从无序的里边拿一个与有序的比较,小的话就交换，最开始认为有序的里边只有第一个元素
*/
func main() {
	arr := [10]int{3, 1, 9, 7, 3, 55, 8, 1, 6, 4}
	sort(&arr)
	fmt.Println(arr)
}
func sort(arr *[10]int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
}
