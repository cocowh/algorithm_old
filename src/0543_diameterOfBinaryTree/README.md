# [二叉树的直径](https://leetcode-cn.com/problems/diameter-of-binary-tree/)

### 题目

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

```
      1
     / \
    2   3
   / \     
  4   5  
```
  
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。

### 解法

```
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(node *TreeNode) (length, res int) {
	if node == nil {
		return
	}

	leftLen, leftDia := helper(node.Left)
	rightLen, rightDia := helper(node.Right)

	length = max(leftLen, rightLen) + 1
	res = max(leftLen+rightLen, max(leftDia, rightDia))
	return
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}
```