package test

import (
	"fmt"
	"regexp"
	"xiaoma-bot/books/service"
)

/**
测试用例
@MYT 20240810
*/

// TestCal 测试方法
func TestCal() {

	pattern := regexp.MustCompile(`/0.`)
	matchString := pattern.MatchString("1/0")
	fmt.Println(matchString)
	//if matched, _ := regexp.MustCompile(`\/0`, "1/0."); matched {
	//	fmt.Println("0不能出现在/0后面")
	//}
	postfix := service.Convert("1.6/0.2")
	fmt.Println(postfix)
	calculate := service.ComputeSuffixStr(postfix)
	fmt.Println(calculate)
}
