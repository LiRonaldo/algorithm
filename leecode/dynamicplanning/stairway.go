package main

import (
	"fmt"
)

/**
有一座高度是10级台阶的楼梯，从下往上走，每跨一步只能向上1级或者2级台阶。要求用程序来求出一共有多少种走法。
当差最后一步当第10级的时候，是在第8阶或者第9阶。所以f（10）=f（9）+f(8)
以此类推，f（9）=f(8)+f(7)
最终
只有一阶的时候只有一种走法 f(1)=1
两阶的时候 一种是一步一步，一种是一下两部，f(2)=2,
这样就是斐波那契数列

soul2方法是对斐波那契数列的优化
*/

func main() {
	num := soul1(10)
	fmt.Println(num)
	num1 := soul2(10)
	fmt.Println(num1)
}

/**
斐波那契
*/
func soul1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return soul1(n-1) + soul1(n-2)
}

/**
优化斐波那契
a    b    tmp
1    2    3      5
*/
func soul2(n int) int {
	a := 1
	b := 2
	tmp := 0
	for i := 2; i < n; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}
