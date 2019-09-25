# [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

### 题目

给定一个二叉树，返回它的 前序 遍历。
 示例:

```bash
输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [1,2,3]
```
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

### 解法

前序遍历：根->左->右。
 
 * 先输出根
 * 输出根左子至左子为nil
 * 回溯到有右子的根继续输出
   
```
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	ret := append([]int{},root.Val)
	ret = append(ret,preorderTraversal(root.Left)...)
	ret = append(ret,preorderTraversal(root.Right)...)
	return  ret
}

func preorderTraversalByStack(root *TreeNode) []int {
	ret,stack,point := []int{},[]*TreeNode{},root
	for point != nil || len(stack) > 0{
		if point != nil {
			ret = append(ret,point.Val)
			stack = append(stack,point)
			point = point.Left
		} else {
			point = stack[len(stack) - 1]
			point = point.Right
			stack = stack[:len(stack) - 1]
		}
	}
	return ret
}
```