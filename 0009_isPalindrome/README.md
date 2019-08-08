# [回文数](https://leetcode-cn.com/problems/palindrome-number/)

### 题目

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。


进阶:

你能不将整数转为字符串来解决这个问题吗？


### 解法

 * 边界：负数定为非回文数，个位数必为回文数
 * 回文数反转值与原值相等
 * 相等数异或值为0

```
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if 0 <= x && x < 9 {
        return true
    }
    reverse := 0
	temp := x
    for temp != 0 {
        rem := temp % 10 //求余
		temp /= 10
		reverse = reverse*10 + rem
    }
    if reverse ^ x == 0 {
        return true
    } else {
        return false
    }
}
```

