# [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

### 题目
给定一个二叉树，返回它的中序 遍历。

示例:
```bash
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
```
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

### 解法
   
```
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func inorderTraversalByStack(root *TreeNode) []int {
	res,stack,point := []int{},[]*TreeNode{},root
	for point != nil || len(stack) > 0{
		for point != nil {
			stack = append(stack,point)
			point = point.Left
		}
		point = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res,point.Val)
		point = point.Right
	}
	return res
}
```