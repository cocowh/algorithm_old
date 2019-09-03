# [报数](https://leetcode-cn.com/problems/count-and-say/)

### 题目
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

```bash
1.     1
2.     11
3.     21
4.     1211
5.     111221
```
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

示例 1:

```bash
输入: 1
输出: "1"
```
示例 2:

```bash
输入: 4
输出: "1211"
```
### 解法

一开始没理解题意，看了解析。就是把上一个数字的读法作为下一个数字的表示

一开始使用字符串拼接解答，运时和内存感人，使用字节数组转换为，效果惊人。学到了。
    
```
func countAndSay(n int) string {
	res := "1"
	for i := 1;i < n; i++ {
		resTemp := ""
		for j:= 0;j < len(res); {
			count := 1
			for k := j + 1;k < len(res) && res[k] == res[j];k++{
				count++
			}
			resTemp = resTemp + strconv.Itoa(count) + string(res[j])
			j += count
		}
		res = resTemp
	}
	return res
}

func countAndSayBetter(n int) string {
	res := []byte{'1'}
	for i := 1;i < n; i++ {
		resTemp := []byte{}
		for j:= 0;j < len(res); {
			count := 1
			for k := j + 1;k < len(res) && res[k] == res[j];k++{
				count++
			}
			resTemp = append(resTemp, byte(count + 48), res[j])
			j += count
		}
		res = resTemp
	}
	return string(res)
}
```