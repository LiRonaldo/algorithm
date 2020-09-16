package main

import "fmt"

/**
给定一个二进制数组, 找到含有相同数量的 0 和 1 的最长连续子数组。
*/
func main() {
	arr := []int{1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}
	//先把数组元素是0，变成-1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i] = -1
		}
	}
	var sum = 0
	var ans = 0
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == 0 {
			if i >= ans {
				ans = i + 1
			}
		}
		if _, ok := m[arr[i]]; !ok {
			m[sum] = i
			continue
		}
		temp := i - m[sum]
		if temp > ans {
			ans = temp
		}
	}
	fmt.Println(ans)

}
