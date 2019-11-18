# [路径总和 II](https://leetcode-cn.com/problems/path-sum/)

### 题目

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:  
给定如下二叉树，以及目标和 sum = 22，

```
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
```

返回:

```
[
   [5,4,11,2],
   [5,8,4,5]
]
```

### 解法

golang知识点：copy值复制，=复制指针复制。=复制的值底层改变时，影响已经复制的副本。

```
func helper(node *TreeNode, sum int, res *[][]int, stack []int) {
	if node == nil {
		return
	} else if sum-node.Val == 0 && node.Left == nil && node.Right == nil {
		temp := make([]int, len(stack)+1)
		copy(temp, append(stack, node.Val))
		*res = append(*res, temp)
	} else {
		helper(node.Left, sum-node.Val, res, append(stack, node.Val))
		helper(node.Right, sum-node.Val, res, append(stack, node.Val))
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	helper(root, sum, &res, []int{})
	return res
}
```
