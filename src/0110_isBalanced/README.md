# [平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/)
### 题目

给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

```
    3
   / \
  9  20
    /  \
   15   7
```
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

```
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```
返回 false 。

### 解法

```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	flag, _ := helper(root)
	return flag
}

func helper(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	flagL, levelL := helper(node.Left)
	flagR, levelR := helper(node.Right)
	if flagL && flagR && levelL-levelR >= -1 && levelL-levelR <= 1 {
		return true, max(levelR, levelL) + 1
	}
	return false, 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
