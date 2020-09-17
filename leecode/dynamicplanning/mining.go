package main

import (
	"fmt"
	"math"
)

/**
有一个国家发现了5座金矿，每座金矿的黄金储量不同，需要参与挖掘的工人数也不同。
参与挖矿工人的总数是10人。每座金矿要么全挖，要么不挖，不能派出一半人挖取一半金矿。
要求用程序求解出，要想得到尽可能多的黄金，应该选择挖取哪几座金矿？
最优子结构： f(10,5)=max(f(10,4),f(10-worker[4],4)+money[4])
边界值： 当只有一座金矿，p<worker[0],0
         当只有一座金矿时，p>worker[0]，money【0】
状态转移方程：f(w,n)=max(f(w,n-1),f(w-woker[n-1],n-1)+money[n-1])

*/

var money []int = []int{400, 500, 200, 300, 350}
var worker []int = []int{5, 5, 3, 4, 3}

func main() {
	fmt.Println(caucal(10, 5))

}
func caucal(p, n int) int {
	if n == 1 && p < worker[0] {
		return 0
	}
	if n == 1 && p >= worker[0] {
		return money[0]
	}
	//人数不够挖第5座
	if n > 1 && p < worker[n-1] {
		caucal(p, n-1)
	}
	if n > 1 && p >= worker[n-1] {
		return int(math.Max(float64(caucal(p, n-1)), float64(caucal(p-worker[n-1], n-1)+money[n-1])))
	}
	return 0
}
