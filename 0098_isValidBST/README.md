# [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

### 题目

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

```
输入:
    2
   / \
  1   3
输出: true
```

示例 2:

```
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
```
解释: 输入为: [5,1,4,null,null,3,6]。  
     根节点的值为 5 ，但是其右子节点值为 4 。

### 解法

在遍历树的同时保留结点的上界与下界，在比较时不仅比较子结点的值，也要与上下界比较。

```
func isValidBST(root *TreeNode) bool {
	return helper(root, -1<<63, 1<<63-1)
}
func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	return min < node.Val && node.Val < max && helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
```
