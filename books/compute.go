package books

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"xiaoma-bot/dto"
	stack "xiaoma-bot/stack"
)

/**
计算逻辑业务
@MYT 20240810
*/

const (
	info = "<@!16976044954935828807>"
	// 默认提示
	message  = "计数请@我并输入指令(/计数),{表达式},我会自动帮你计算！"
	compute  = "/计数"
	invalid  = "无法识别指令"
	errorRes = "表达式有误"
)

// Compute 计算逻辑
func Compute(loadMsg *dto.LoadMsg) string {
	content := loadMsg.Data.Content
	// 替换指令内容
	content = strings.Replace(content, info, "", -1)
	content = strings.TrimSpace(content)
	// 默认提示
	if content == "" {
		return message
	}
	if strings.HasPrefix(content, compute) {
		re := regexp.MustCompile(`{([^}]+)}`)
		matches := re.FindAllStringSubmatch(content, -1)
		if matches != nil {
			str := matches[0][1]
			err := ValidateInput(str)
			if err != nil {
				content = errorRes
			} else {
				// 中缀转后缀
				postfix := Convert(str)
				// 计算后缀表达结果
				res := ComputeSuffixStr(postfix)
				// 去除小数点后的0
				zero := removeZero(res)
				content = fmt.Sprintf("%s = %s", str, zero)
			}
		} else {
			content = errorRes
		}
	} else {
		content = invalid
	}
	return content
}

// 定义符号优先级
var priority = map[rune]int{
	'%': 2,
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'(': 0,
}

// Convert 中缀表达式转后缀表达式
func Convert(input string) []interface{} {
	// 定义一个存放处理后的字符的切片
	var inter []interface{}
	// 创建一个栈，用于存放运算符
	suffixStr := stack.NewSliceStack()
	for i := 0; i < len(input); i++ {
		// 转换输入的字段为rune类型方便值的判断
		char := rune(input[i])
		// 字符为空格时，跳过此次循环
		if char == ' ' {
			continue
		}
		//处理正负数字
		if unicode.IsDigit(char) || char == '.' || (char == '-' && (i == 0 || (i > 0 && (rune(input[i-1]) == '(' || rune(input[i-1]) == '+' || rune(input[i-1]) == '-' || rune(input[i-1]) == '*' || rune(input[i-1]) == '/')))) {
			num := ""
			if char == '-' {
				num += string(char)
				i++
				char = rune(input[i])
			}
			for i < len(input) && (unicode.IsDigit(char) || char == '.') {
				num += string(char)
				i++
				if i < len(input) {
					char = rune(input[i])
				}
			}
			i-- //回退一步，因为外层也有一个 i++
			// 将字符转化为浮点型
			value, _ := strconv.ParseFloat(num, 64)
			// 将处理过后的数据放入inter切片中存放
			inter = append(inter, value)
			// 在符号为"("的情况下直接放入字符栈中
		} else if char == '(' {
			suffixStr.Push(char)
		} else if char == ')' {
			// 在存放运算符的栈不为空，并且栈顶字符不为"("的情况下
			for !suffixStr.IsEmpty() && suffixStr.Peek().(rune) != '(' {
				// 把运算符栈中的运算符取出放进inter中
				inter = append(inter, suffixStr.Pop())
			}
			// 取出栈中"("
			suffixStr.Pop()
		} else {
			// 处理运算符优先级
			// 在栈不为空，并且新的运算符优先级比栈顶的小的情况下
			for !suffixStr.IsEmpty() && priority[suffixStr.Peek().(rune)] >= priority[char] {
				// 把运算符放入inter中
				inter = append(inter, suffixStr.Pop())
			}
			suffixStr.Push(char)
		}
	}
	// 把剩余的操作符全部放入inter切片中
	for !suffixStr.IsEmpty() {
		inter = append(inter, suffixStr.Pop())
	}
	// 返回切片
	return inter
}

// ComputeSuffixStr 计算后缀表达式
func ComputeSuffixStr(inter []interface{}) float64 {
	// 创建一个用存放数值
	stored := stack.NewSliceStack()
	// 遍历数值
	for _, value := range inter {
		// 判断遍历数值的类型
		switch value.(type) {
		case float64:
			// 当类型为float64时，把值压入栈中
			stored.Push(value)
		case rune:
			// 当值为rune时判定为计算符，调用计算函数
			CalculatingFunction(stored, value.(rune))
		}
	}
	//返回栈低最后一个值作为结果
	return stored.Pop().(float64)
}

// CalculatingFunction 计算函数
func CalculatingFunction(stored *stack.SliceStack, operator rune) {
	//取出两个操作数
	num1 := stored.Pop().(float64)
	num2 := stored.Pop().(float64)
	var result float64
	switch operator {
	case '+':
		result = num2 + num1
	case '-':
		result = num2 - num1
	case '*':
		result = num2 * num1
	case '/':
		// 被除数为零时输出错误和终止程序
		if num1 == 0 {
			fmt.Println("错误：计算过程中出现被除数等于零的情况，请检查算式合法性")
			os.Exit(1)
		} // 除法注意先后顺续不要错
		result = num2 / num1
	case '%':
		result = math.Mod(num2, num1)
	}
	// 将结果压回栈
	stored.Push(result)
}

// 动态去除小数点后的零
func removeZero(f float64) string {
	s := fmt.Sprintf("%f", f)
	s = strings.TrimRight(s, "0")
	// 检查移除多余零最后一位是否为‘.’结尾
	if strings.HasSuffix(s, ".") {
		s = strings.TrimRight(s, ".")
	}
	return s
}

// ValidateInput 校验输入的合法性
func ValidateInput(expr string) error {
	// 移除空格
	expr = strings.ReplaceAll(expr, " ", "")
	// 检查是否包含非法字符
	if matched, _ := regexp.MatchString(`[^0-9+\-*/().% ]`, expr); matched {
		return fmt.Errorf("输入字符不合法")
	}
	// 检查括号是否成对出现
	if strings.Count(expr, "(") != strings.Count(expr, ")") {
		return fmt.Errorf("输入括号不匹配")
	}
	// 检查开头和结尾是否为非法字符
	if matched, _ := regexp.MatchString(`^[*/%]|\s$`, expr); matched {
		return fmt.Errorf("开头或者结尾字符不合法")
	}
	// 检查连续的操作符
	if matched, _ := regexp.MatchString(`\+\+|--|\*\*|%%|\(\)|//`, expr); matched {
		return fmt.Errorf("出现连续的符号")
	}
	if matched, _ := regexp.MatchString(`/0`, expr); matched {
		return fmt.Errorf("0不能出现在/0后面")
	}
	return nil
}
