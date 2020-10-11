package main

import "fmt"

/**
快速排序 一个指针，假设a[0]最小，往后找，找到一个比a[0]小的a[j]，将a[j]和第i个(i从1开始自增，)交换,继续查找，在找到就和第2个数交换。循环完之后，就讲a[0]和a[i]交换
*/
func main() {
	arr := []int{6, 5, 4, 1}
	sort(arr, 0, len(arr)-1) //注意传值是传len-1，因为下边判断用的是<=
	fmt.Println(arr)
}
func sort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := findMid(arr, l, r)
	sort(arr, l, p-1)
	sort(arr, p+1, r)
}
func findMid(arr []int, l, r int) int {
	i := l
	v := arr[l]
	for j := l + 1; j <= r; j++ {
		if arr[j] < v {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[i] = arr[i], arr[l]
	return i
}
