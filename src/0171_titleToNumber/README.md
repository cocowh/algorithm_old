# [Excel表列序号](https://leetcode-cn.com/problems/excel-sheet-column-number/)

### 题目

给定一个Excel表格中的列名称，返回其相应的列序号。


例如，

>  A -> 1  
   B -> 2    
   C -> 3   
   ...  
   Z -> 26  
   AA -> 27  
   AB -> 28   
   ...    

示例 1:

>输入: "A"  
输出: 1

示例 2:  

>输入: "AB"  
输出: 28

示例 3:  

>输入: "ZY"  
输出: 701

### 解法

从前往后比从后往前占用内存多0.1m，执行用时一会儿0ms一会儿4ms，判体系统用例不稳定。

```
func titleToNumber(s string) int {
	len,ret := len(s),0
	for i := len - 1; i >=0; i-- {
		ret += int(s[i] - 'A' + 1) * int(math.Pow(26 , float64(len - 1 - i)))
	}
	return ret
}

func titleToNumberFromHead(s string) int {
	step,ret := 0,0
	for i := len(s) - 1; i >=0; i-- {
		ret += int(s[i] - 'A' + 1) * int(math.Pow(26 , float64(step)))
		step++
	}
	return ret
}
```
