# [路径总和](https://leetcode-cn.com/problems/path-sum/)

### 题目

给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:   
给定如下二叉树，以及目标和 sum = 22，

```
          5
         / \
        4   8
       /   / \
      11  13  4
     /  \      \
    7    2      1
```
    
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

### 解法

```
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum-root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
```