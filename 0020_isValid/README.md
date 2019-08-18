# [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

### 题目

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

### 解法

使用map存映射关系，遇`(`、`{`、`[`入栈，遇`]`、`}`、`)`出栈查map，最后看栈是否为空。

```
func isValid(s string) bool {
	m := map[byte]byte{
		'(':')',
		'{':'}',
		'[':']',
	}
	stack := []byte{}
	for _,v := range s {
		if v == '{' || v == '(' || v == '['  {
			stack = append(stack,byte(v))
		} else if v == '}' && v == ')' || v == ']' {
			if len(stack) > 0 || m[stack[len(stack) - 1]] == byte(v) {
				stack = stack[:len(stack) - 1]
				continue
			} else {
				return false
			}
		}
	}
	return  len(stack) == 0
}
```