package stack

import (
	"sync"
)

type Item string

type ItemStack struct {
	items []string
	lock  sync.RWMutex
}

// New 创建栈
func (s *ItemStack) New() *ItemStack {
	s.items = []string{}
	return s
}

// Push 入栈
func (s *ItemStack) Push(t string) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop 出栈
func (s *ItemStack) Pop() string {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return item
}

// Top 取栈顶
func (s *ItemStack) Top() string {
	return s.items[len(s.items)-1]
}

// IsEmpty 判空
func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}
