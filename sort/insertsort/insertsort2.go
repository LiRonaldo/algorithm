package main

import "fmt"

/**
插入排序优化版，有点像选择排序。不优化的版本是从i往前找小的。找到就交换。优化之后是不交换，i往前都往后移动，然后将arr[i]放到j的位置上。
*/
func main() {
	arr := []int{70, 8, 4, 5, 8, 11, 119, 12, 15}
	sort2(arr)
	fmt.Println(arr)
}
func sort2(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		var j int
		t := arr[i]
		//此处有点类似选择排序，选择最小的，只不过这里要后移
		//j为啥等于i不能是i-1 ,因为j如果=i-1的话，arr[j-1]的话，会越界。
		for j = i; j > 0 && arr[j-1] > t; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}
