# [Excel表列名称](https://leetcode-cn.com/problems/excel-sheet-column-title/)

### 题目

给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

>    1 -> A  
    2 -> B  
    3 -> C  
    ...  
    26 -> Z  
    27 -> AA  
    28 -> AB   
    ...  

示例 1:

>输入: 1  
输出: "A"  

示例 2:  

>输入: 28  
输出: "AB"  

示例 3:  

>输入: 701  
输出: "ZY"  

### 解法

类似进制转换。

三种不同的操作，后两种内存占用和运行效率一致

```

func convertToTitle(n int) string {
	m := map[int]byte{
		1:  'A',
		2:  'B',
		3:  'C',
		4:  'D',
		5:  'E',
		6:  'F',
		7:  'G',
		8:  'H',
		9:  'I',
		10: 'J',
		11: 'K',
		12: 'L',
		13: 'M',
		14: 'N',
		15: 'O',
		16: 'P',
		17: 'Q',
		18: 'R',
		19: 'S',
		20: 'T',
		21: 'U',
		22: 'V',
		23: 'W',
		24: 'X',
		25: 'Y',
		26: 'Z',
	}
	ret := []byte{}
	for n > 0 {
		if n%26 == 0 {
			ret = append([]byte{'Z'}, ret...)
			n = n/26 - 1
		} else {
			ret = append([]byte{m[n%26]}, ret...)
			n /= 26
		}
		fmt.Println(string(ret), n)
	}
	return string(ret)
}

func convertToTitleByString(n int) string {
	m := "#ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := []byte{}
	for n > 0 {
		fmt.Println(n % 26)
		if n%26 == 0 {
			ret = append([]byte{'Z'}, ret...)
			n = n/26 - 1
		} else {
			ret = append([]byte{m[n%26]}, ret...)
			n /= 26
		}
		fmt.Println(string(ret), n)
	}
	return string(ret)
}

func convertToTitleByChar(n int) string {
	ret := []byte{}
	for n > 0 {
		fmt.Println(n % 26)
		if n%26 == 0 {
			ret = append([]byte{'Z'}, ret...)
			n = n/26 - 1
		} else {
			ret = append([]byte{byte(n%26 + 'A' - 1)}, ret...)
			n /= 26
		}
		fmt.Println(string(ret), n)
	}
	return string(ret)
}
```
