# [二进制求和](https://leetcode-cn.com/problems/add-binary/)

### 题目
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

### 解法

简单难度。

考虑进位。直接遍历或递归。
   
```
func addBinary(a string, b string) string {
	an,bn := []byte(a),[]byte(b)
	lena,lenb := len(a),len(b)
	if lena < lenb {
		for i := lena;i < lenb; i++ {
			an = append([]byte{'0'},an...)
		}
	} else if lena > lenb {
		for i := lenb;i < lena; i++{
			bn = append([]byte{'0'},bn...)
		}
	}
	fmt.Println(string(an),string(bn))
	s,c := []byte{},false
	lena = len(an)
	for i := 0; i < lena; i++ {
		fmt.Println(i,lena - 1 - i)
		bytea,byteb,sum := an[lena - 1 - i],bn[lena - 1 - i],byte(0)
		if c {
			sum += bytea + byteb - '0' - '0' + 1
		} else {
			sum += bytea + byteb - '0' - '0'
		}
		if sum / 2 > 0 {
			c = true
		} else {
			c = false
		}

		s  = append([]byte{ sum % 2 + '0'}, s...)
	}
	if c {
		return string(append([]byte{'1'},s...))
	}
	return string(s)
	/*if len(a) < len(b) {
		temp := a
		a = b
		b = temp
	}
	lena,lenb,s,c := len(a),len(b),[]byte{},false
	return string(sumTwoBinaryString(lena,lenb,0, a, b, s, c))*/
}

func sumTwo (a,b byte, flag bool) (bool,byte) {
	sum := byte(0)
	if flag {
		sum = a - '0' + b - '0' + 1
	} else {
		sum = a - '0' + b - '0'
	}
	if sum / 2 > 0{
		return true,sum % 2 + '0'
	} else {
		return false,sum %2 + '0'
	}
}

func sumTwoBinaryString(lena, lenb, point int, a, b string, s []byte, c bool) []byte {
	fmt.Println("----")
	if point >= lena {
		if c {
			return append([]byte{'1'}, s...)
		} else {
			return s
		}
	}
	bytea,byteb := a[lena - 1 - point],byte('0')
	if point >= lenb  {
		byteb = '0'
	} else {
		byteb = b[lenb - 1 - point]
	}
	fmt.Println(string([]byte{bytea,byteb}),c)
	status,r := sumTwo(bytea,byteb,c)
	fmt.Println(status,string([]byte{r}))
	s = append([]byte{r},s...)
	c = status
	point++
	fmt.Println(point, string(s), c)
	return sumTwoBinaryString(lena,lenb,point,a,b,s,c)
}
```