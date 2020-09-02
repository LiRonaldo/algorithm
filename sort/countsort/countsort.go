package main

import "fmt"

/**
计数排序
第一种,就是定义一个和最大的数+1长度的数组,将元素一个个按照大小放到对应的下标中,
重复的就+1,然后依次打印出来下标（数组中的值是几，就打印几次）
第二种,优化,假设待排序的数组数值都比较偏大，{101,109,108,102,110,107,103} 那么要创建一个111长度的数组，但是
比较浪费空间,优化办法是创建待排序数组的最大值-最小值+1的长度数组就可以了.
*/
func main() {

	/*arr := []int{8, 7, 1, 99, 1, 3, 8, 2, 1, 1}
	fmt.Println(arr)
	sortarr1 := sort1(arr)
	print(sortarr1)*/
	arr := []int{101, 109, 108, 102, 110, 107, 103}
	fmt.Println(arr)
	sortArr := sort2(arr)
	fmt.Println(sortArr)
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

func sort2(arr []int) []int {
	min := arr[0]
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
		}
		if arr[i] >= max {
			max = arr[i]
		}
	}
	sortArr := make([]int, max-min+1, max-min+1)
	sortArr2 := make([]int, max-min+1, max-min+1)
	for i := 0; i < len(arr); i++ {
		sortArr[arr[i]-min]++
	}
	var k int = 0
	for i := 0; i < len(sortArr); i++ {
		for j := 0; j < sortArr[i]; j++ {
			sortArr2[k] = sortArr[i] + min + i - 1
			k++
		}
	}
	return sortArr2
}
