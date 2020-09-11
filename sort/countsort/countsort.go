package main

import "fmt"

/**
计数排序
第一种,就是定义一个和最大的数+1长度的数组,将元素一个个按照大小放到对应的下标中,
重复的就+1,然后依次打印出来下标（数组中的值是几，就打印几次）
第二种,优化,假设待排序的数组数值都比较偏大，{101,109,108,102,110,107,103} 那么要创建一个111长度的数组，但是
比较浪费空间,优化办法是创建待排序数组的最大值-最小值+1的长度数组就可以了.
第三种，优化就是假如待排序数组有重复的，相同元素按照本来的顺序的排列
*/
func main() {

	/*arr := []int{8, 7, 1, 99, 1, 3, 8, 2, 1, 1}
	fmt.Println(arr)
	sortarr1 := sort1(arr)
	print(sortarr1)*/
	arr := []int{101, 109, 107, 103, 108, 102, 103, 110, 107, 103}
	/*fmt.Println(arr)
	sortArr := sort2(arr)
	fmt.Println(sortArr)*/
	final := sort3(arr)
	fmt.Println(final)
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
	count := make([]int, max-min+1, max-min+1)
	result := make([]int, max-min+1, max-min+1)
	for i := 0; i < len(arr); i++ {
		count[arr[i]-min]++
	}
	var k int = 0
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			result[k] = count[i] + min + i - 1
			k++
		}
	}
	return result
}

//区分相同元素的原始顺序
func sort3(arr []int) []int {
	//先确定最大最小值，相减确定要创建的数组大小
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
	//计数数组
	count := make([]int, max-min+1)
	//最终结果数组
	result := make([]int, max-min+1)
	//开始计数
	for i := 0; i < len(arr); i++ {
		count[arr[i]-min]++
	}
	for i := 1; i < len(arr); i++ {
		count[i] += count[i-1]
	}
	//开始往result数组里赛值
	//从后往前开始遍历原数组，因为重复的数据在后边，所以从后往前，
	for i := len(arr) - 1; i >= 0; i-- {
		//arr[i]元素在原始元素中是第几个，比如是第3个，落到result数组中是3-1位置(因为count中的是数量，就好比10个元素，最后一个是a[9]一样)
		result[count[arr[i]-min]-1] = arr[i]
		//因为已经赋值了一个重复元素了，所以计数数组要减去1
		//如果不想从后往前遍历arr，要从前往后的话count[arr[i]-min]++，原理自己琢磨。
		count[arr[i]-min]--
	}
	return result
}
