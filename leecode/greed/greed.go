package main

import (
	"fmt"
	"sort"
)

/**
贪心
0-1背包问题 物品不可分割，首先按照性价比排序，在负重不超过的情况下，一直装，如果性价比相同，那么装轻的，如果性价比相同，重量相同，装价值高的。
一般背包问题 就是物品可以分割，在负重不超过的情况下，一直装，如果性价比相同，那么装轻的，如果性价比相同，重量相同，装价值高的，如果最后装不下一件完整的物品
就剩下的负重*性价比。
代码就不写了。比较简单
复杂的是排序
*/
type Person struct {
	weight float64
	value  float64
	xjb    float64
}

//定义一个变量，为了给person数组排序
//要重写len swap less 方法
type PersonScile []Person

func (p PersonScile) Len() int {
	return len(p)
}
func (p PersonScile) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PersonScile) Less(i, j int) bool {
	return p[j].xjb < p[i].xjb
}
func main() {
	people := []Person{
		{weight: 17, value: 50},
		{weight: 30, value: 60},
		{weight: 25, value: 70},
	}
	for i := 0; i < len(people); i++ {
		people[i].xjb = people[i].value / people[i].weight
	}
	fmt.Println(people)
	sort.Sort(PersonScile(people))
	fmt.Println(people)
}
