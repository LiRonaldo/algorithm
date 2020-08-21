package main

import (
	"fmt"
)

/**
普通版  要点，选择第一个的，依次和第二个往后开始比，找到就标记下坐标。然后和第一个做比较。
两层for循环
*/
func main() {

	arr := [10]int{4, 5, 1, 6, 3, 7, 12, 43, 53, 9}
	sort(&arr)
	fmt.Println(arr)
}
func sort(arr *[10]int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[minIndex] >= arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}
