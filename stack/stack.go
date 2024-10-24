package stack

type Stack struct {
	elements []interface{} // 存储元素的切片
}

// Push 向栈中添加一个元素
func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

// Pop 从栈中移除并返回顶部元素
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false // 栈为空时返回nil和false
	}
	top := s.elements[len(s.elements)-1]        // 获取顶部元素
	s.elements = s.elements[:len(s.elements)-1] // 移除顶部元素
	return top, true                            // 返回顶部元素和true表示成功移除
}

// Peek 返回顶部元素但不移除它
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false // 栈为空时返回nil和false
	}
	return s.elements[len(s.elements)-1], true // 返回顶部元素和true表示成功获取
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0 // 如果切片为空则返回true，否则返回false
}

// Size 返回栈中的元素数量
func (s *Stack) Size() int {
	return len(s.elements) // 返回切片长度，即栈中的元素数量
}
