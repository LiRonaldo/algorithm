package main

import "fmt"

/**
计数排序
第一种，就是定义一个和最大的数+1长度的数组,将元素一个个按照大小放到对应的下标中,
重复的就+1,然后依次打印出来下标（数组中的值是几，就打印几次）
*/
func main() {

	arr := []int{8, 7, 1, 99, 1, 3, 8, 2, 1, 1}
	fmt.Println(arr)
	sortarr := sort1(arr)
	print(sortarr)
}

func sort1(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max <= arr[i] {
			max = arr[i]
		}
	}
	sortArr := make([]int, max+1, max+1)
	for i := 0; i < len(arr); i++ {
		sortArr[arr[i]]++
	}
	return sortArr
}
func print(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
