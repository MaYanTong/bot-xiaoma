package test

import (
	"fmt"
	"xiaoma-bot/books"
)

/**
测试用例
@MYT 20240810
*/

// TestCal 测试方法
func TestCal() {
	postfix := books.Convert("1+1")
	fmt.Println(postfix)
	calculate := books.ComputeSuffixStr(postfix)
	fmt.Println(calculate)
}
