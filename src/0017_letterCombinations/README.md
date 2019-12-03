# [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)

### 题目

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

![电话号码](https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png)

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

### 解法

```
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	ret := []string{""}
	m := map[byte][]string{
		'2':[]string{"a","b","c"},
		'3':[]string{"d","e","f"},
		'4':[]string{"g","h","i"},
		'5':[]string{"j","k","l"},
		'6':[]string{"m","n","o"},
		'7':[]string{"p","q","r","s"},
		'8':[]string{"t","u","v"},
		'9':[]string{"w","x","y","z"},
	}
	for _,v := range digits{
		copy := []string{}
		retLen := len(ret)
		for j := 0; j < retLen; j++ {
			arrLen := len(m[byte(v)])
			for i := 0; i < arrLen; i++ {
				copy = append(copy, ret[j] + m[byte(v)][i])
			}
		}
		ret = copy
	}
	return ret
}
```