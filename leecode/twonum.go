package main

import "fmt"

/**
求两个数的和
1.种是暴力
2.是利用map做处理。
*/

func main() {
	m := make(map[int]int)
	s := []int{1, 3, 4, 7, 7, 6}
	target := 11
	for i := 0; i < len(s); i++ {
		n := target - s[i]
		if value, ok := m[n]; ok {
			fmt.Println(value, i)
			break
		} else {
			m[s[i]] = i
		}
	}
}
