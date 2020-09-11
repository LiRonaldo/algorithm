package main

import (
	"algorithm/tree/stack"
	"fmt"
	"strings"
)

func main() {
	b := validate()
	fmt.Println(b)
}

func validate() bool {
	var s = "{}[]"
	stack := &stack.Stack{}
	for i := 0; i < len(s); i++ {
		if strings.Compare(string(s[i]), "{") == 0 {
			stack.Push(s[i])
		} else if (strings.Compare(string(s[i]), "[")) == 0 {
			stack.Push(s[i])
		} else if (strings.Compare(string(s[i]), "(")) == 0 {
			stack.Push(s[i])
		} else if (strings.Compare(string(s[i]), "}")) == 0 {
			v := stack.Pop()
			switch t := v.(type) {
			case uint8:
				if strings.Compare(string(t), "{") == 0 {
					return true
				} else {
					return false
				}
			}
		} else if (strings.Compare(string(s[i]), "]")) == 0 {
			v := stack.Pop()
			switch t := v.(type) {
			case uint8:
				if strings.Compare(string(t), "[") == 0 {
					return true
				} else {
					return false
				}
			}
		} else if (strings.Compare(string(s[i]), ")")) == 0 {
			v := stack.Pop()
			switch t := v.(type) {
			case uint8:
				if strings.Compare(string(t), "(") == 0 {
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}
