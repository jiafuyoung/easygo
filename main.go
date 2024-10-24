package main

import (
	"fmt"

	s "github.com/jiafuyoung/easygo/stack"
)

func main() {
	stack := s.Stack{}                               // 创建一个空栈
	stack.Push(1)                                    // 入栈操作
	stack.Push(2)                                    // 入栈操作
	top, ok := stack.Peek()                          // 查看栈顶元素但不移除它
	fmt.Println("Top element:", top, "OK:", ok)      // 输出: Top element: 2 OK: true
	stack.Pop()                                      // 出栈操作但不获取元素值，仅检查是否成功移除顶部元素
	fmt.Println("Stack is empty:", !stack.IsEmpty()) // 检查栈是否为空并输出结果: Stack is empty: false
}
