# [左叶子之和](https://leetcode-cn.com/problems/sum-of-left-leaves/)

### 题目

计算给定二叉树的所有左叶子之和。

示例：

```
    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
```

### 解法

递归，左子树标识

```
func helper(node *TreeNode, flag bool, res int) int  {
	if node == nil {
		return res
	}
	if node.Right == nil && node.Left == nil && flag {
		return res + node.Val
	}
	return helper(node.Left, true, res) + helper(node.Right, false, res)
}
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root.Left, true, 0) + helper(root.Right, false, 0)
}
```
