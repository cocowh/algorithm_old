# [对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/)

### 题目

给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

```bash
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

```bash
    1
   / \
  2   2
   \   \
   3    3
```

说明:

 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

### 解法

递归

```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return (node1.Val == node2.Val) && isMirror(node1.Right, node2.Left) && isMirror(node1.Left, node2.Right)
}
```
