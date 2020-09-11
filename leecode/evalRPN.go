package main

import (
	"algorithm/tree/stack"
	"fmt"
	"strconv"
	"strings"
)

/**
go语言的interface 相当于java的object ，使用的时候 用switch 判断 v.(type) 相当于 java的 object interfaceof  v ，属于的话，在做处理，
但是也可以不适用判断，直接 s：=v.(string) 直接将v转化字符串s
*/

func main() {
	sum := eval("4,13,5,/,+")
	fmt.Println(sum)
}

func eval(str string) interface{} {
	s := strings.Split(str, ",")
	stack := &stack.Stack{}
	for i := 0; i < len(s); i++ {
		if strings.Compare(s[i], "+") == 0 {
			value := stack.Pop()
			s2 := value.(string)
			s3, _ := strconv.Atoi(s2)
			value1 := stack.Pop()
			s21 := value1.(string)
			s31, _ := strconv.Atoi(s21)
			stack.Push(strconv.Itoa(s3 + s31))

		} else if strings.Compare(s[i], "-") == 0 {
			value := stack.Pop()
			s2 := value.(string)
			s3, _ := strconv.Atoi(s2)
			value1 := stack.Pop()
			s21 := value1.(string)
			s31, _ := strconv.Atoi(s21)
			stack.Push(strconv.Itoa(s31 - s3))

		} else if strings.Compare(s[i], "*") == 0 {
			value := stack.Pop()
			s2 := value.(string)
			s3, _ := strconv.Atoi(s2)
			value1 := stack.Pop()
			s21 := value1.(string)
			s31, _ := strconv.Atoi(s21)
			stack.Push(strconv.Itoa(s3 * s31))
		} else if strings.Compare(s[i], "/") == 0 {
			value := stack.Pop()
			s2 := value.(string)
			s3, _ := strconv.Atoi(s2)
			value1 := stack.Pop()
			s21 := value1.(string)
			s31, _ := strconv.Atoi(s21)
			stack.Push(strconv.Itoa(s31 / s3))
		} else {
			stack.Push(s[i])
		}
	}
	return stack.Pop()
}
