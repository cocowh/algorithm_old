# [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)

### 题目

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

```
    3
   / \
  9  20
    /  \
   15   7
```

返回它的最小深度  2.

### 解法

[官方](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/solution/er-cha-shu-de-zui-xiao-shen-du-by-leetcode/)

终止条件：

* 叶子节点的定义是左孩子和右孩子都为 null 
* 当 root 节点左右孩子都为空时，返回 1
* 当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
* 当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值

```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDep := int(^uint(0) >> 1)
	if root.Left != nil {
		minDep = min(minDepth(root.Left), minDep)
	}
	if root.Right != nil {
		minDep = min(minDepth(root.Right), minDep)
	}
	return minDep + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```
