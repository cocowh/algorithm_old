# [计数质数](https://leetcode-cn.com/problems/count-primes/)

### 题目

统计所有小于非负整数 n 的质数的数量。

示例:

>输入: 10  
输出: 4  
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

### 解法

* 确定head.val != val，遍历head至新head
* 遍历head，确定head.Next.Val != val

```
func countPrimes(n int) int {
	count, m := 0, make([]bool, n)
	for i := 2; i < n; i++ {
		if !m[i] {
			count++
			for j := i + i; j < n; j += i {
				m[j] = true
			}
		}
	}
	return count
}

func countPrimesByBitmap(n int) int {
	count, size := 0, strconv.IntSize
	m := make([]int, n/size+1)
	for i := 2; i < n; i++ {
		if m[i/size]&(1<<uint(i&(size-1))) == 0 {
			count++
			for j := i + i; j < n; j += i {
				m[j/size] |= 1 << uint(j&(size-1))
			}
		}
	}
	return count
}
```
