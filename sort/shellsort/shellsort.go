package main

import "fmt"

/**
希尔排序 ,就是分组的插入排序。只要掌握插入排序。然后 i++的时候要i+分组的
确定一个增量一般是n=n/2
总得思路是： 假设len（arr）=6,那么n=3,
相当于普通的插入排序，但是和普通的插入排序不一样，普通的插入排序是i++,希尔排序是i+n,
也就事原数组 a[0],a[3]a[0] 他们三个元素做插入排序。
*/
func main() {
	arr := []int{7, 3, 4, 8, 18, 1}
	sort(arr)
	fmt.Println(arr)
}
func sort(arr []int) {
	n := len(arr)
	for n > 0 {
		//确定增量,增量每次都/2.
		n = n / 2
		//相当于插入排序
		for i := n; i < len(arr); i++ {
			var temp = arr[i]
			var k int
			for k = i - n; k >= 0 && arr[k] > temp; k = k - n {
				arr[k+n] = arr[k]
			}
			arr[k+n] = temp
		}
	}
}
