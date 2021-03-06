# [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)

### 题目

翻转一棵二叉树。

示例：

输入：

```bash
     4  
   /   \  
  2     7  
 / \   / \  
1   3 6   9
```
输出：

```bash
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

备注:  
这个问题是受到 Max Howell 的 原问题 启发的 ：
>谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。

### 解法

递归左右子树交换。

```
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	invertTreeByRecursive(root)
	return root
}

func invertTreeByRecursive(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
	invertTreeByRecursive(node.Left)
	invertTreeByRecursive(node.Right)
}
```
