package main

import "fmt"

/**
堆排序
大根堆 就是叶子节点都比子叶子节点大，这样的话，第0个元素和最后一个元素比较，最终顺序是升序
小根堆就是叶子节点都比子叶子节点小。第0个元素和最后一个元素交换，最终是降序。
第一个非叶子节点是len/2-1  一个叶子节点的左右子叶子节点是 2*i+1,2*i+2
*/
func main() {
	arr := []int{7, 8, 1, 3, 7, 11}
	sort(arr)
	fmt.Println(arr)
}
func sort(arr []int) {
	n := len(arr)
	//先创建大根堆,让他复合大根堆的特性
	makeMaxHeap(arr, n)
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		n--
		//因为已经是大根堆了，满足大根堆的性质了,所以就不需要调用makeMaxHeap方法了。所以直接调用方法去调整大根堆。
		handlHeap(arr, n, 0)
	}
}
func makeMaxHeap(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		handlHeap(arr, n, i)
	}
}
func handlHeap(arr []int, n, i int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i
	//假设arr【i】最大，不需要循环遍历。
	//左子树最大
	if left < n && arr[left] > arr[max] {
		max = left
	}
	//右子树最大
	if right < n && arr[right] > arr[max] {
		max = right
	}
	//如果顶节点不是最大，那么交换，因为已经是大根堆了，所以只需要将左右子节点中最大的和顶节点交换
	//这样顶节点最小被左右最大的交换了，然后从max（就是左右子节点中最大的）处再开始调整，变成大根堆
	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		handlHeap(arr, n, max)
	}
}
