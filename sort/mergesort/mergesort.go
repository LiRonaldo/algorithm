package main

import "fmt"

func main() {
	arr := []int{5, 4}
	MergeSort(arr, len(arr))
	fmt.Println(arr)
}

func MergeSort(arr []int, n int) {
	mergesort(arr, 0, n-1)
}

func mergesort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergesort(arr, l, mid)
	mergesort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}
