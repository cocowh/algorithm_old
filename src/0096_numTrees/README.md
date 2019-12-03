# [不同的二叉搜索树](https://leetcode-cn.com/problems/unique-binary-search-trees/)

### 题目
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

>输入: 3  
输出: 5  
解释:  
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
```bash
   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

### 解法

头条面试遇到过，没解出来，陷入二叉添加子节点的思维中无法自拔。

解法参考[动态规划](https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode/)
   
```
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func numTreesByFormula(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = ret * 2 * (2*i + 1) / (i + 2)
	}
	return ret
}
```