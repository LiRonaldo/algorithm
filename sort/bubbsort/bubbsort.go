package main

import "fmt"

/**
冒泡排序 选择最大的放到最后，10个数比较9次就好了，所以第一层循环是 n-1，第二层循环是两两比较，因为比较比较过得就不需要比较了，所以是n-i-1
*/
func main() {
	a := [10]int{4, 8, 1, 7, 9, 4, 1, 4, 0, 44}
	sort(&a)
	fmt.Println(a)
}
func sort(arr *[10]int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
