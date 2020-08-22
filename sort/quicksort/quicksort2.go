package main

import "fmt"

/**
快速排序，优化两路指针。  为什么 要从基数的对面开始，自己从左边跑一下就知道了，因为最终目的是左边的数都比
基数小，右边的数都基数大，从左边开始的话，为违反。
思想相比单路的，多了一个从后往前的指针，
i从0位开始,0位也是基数，
j从末位开始，
j先寻找，找到比基数小的就a[i]=a[j]
然后从i开始往右寻找。
找到比基数大的就 a[j]=a[i]
最终i和j交汇，就将a[i]=基数，返回i
然后递归调用。
*/
func main() {
	a := []int{6, 5, 7, 1, 3, 8, 1}
	sort2(a, 0, len(a)-1)
	fmt.Println(a)
}
func sort2(arr []int, l, r int) {
	if l < r {
		p := findMid2(arr, l, r)
		sort2(arr, l, p-1)
		sort2(arr, p+1, r)
	}
}
func findMid2(arr []int, l, r int) int {
	v := arr[l]
	for l < r {
		for l < r && arr[r] >= v {
			r--
		}
		if l < r {
			arr[l] = arr[r]
		}
		for l < r && arr[l] <= v {
			l++
		}

		if l < r {
			arr[r] = arr[l]
		}
	}
	arr[l] = v
	return l
}
