package stack

/**
定义栈结构
@MYT 20240810
*/

// SliceStack 实现栈结构
type SliceStack struct {
	dataStack []interface{}
}

// NewSliceStack 创建一个新的初始栈
func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

// Push 在栈顶放入数据
func (s *SliceStack) Push(data interface{}) {
	s.dataStack = append(s.dataStack, data)
	//append函数：在原切片的末尾添加元素
}

// Pop 栈最上面的数据出栈
func (s *SliceStack) Pop() interface{} {
	if len(s.dataStack) == 0 {
		return nil
	}
	slice := s.dataStack[len(s.dataStack)-1]
	s.dataStack = s.dataStack[:len(s.dataStack)-1]
	//把栈顶下移一位
	return slice
}

// Peek 查看栈顶第一个元素
func (s *SliceStack) Peek() interface{} {
	if len(s.dataStack) == 0 {
		return nil
	}
	return s.dataStack[len(s.dataStack)-1]
}

// IsEmpty 判断栈是否空
func (s *SliceStack) IsEmpty() bool {
	return len(s.dataStack) == 0
}
