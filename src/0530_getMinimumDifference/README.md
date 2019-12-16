# [二叉搜索树的最小绝对差](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/)
### 题目

给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

示例 :

```
输入:

   1
    \
     3
    /
   2

输出:
1

解释:
最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
```

注意: 树中至少有2个节点。

### 解法

二叉搜索树，左<中<右，求差值，中序遍历。对go语法不熟悉，参考[LeetCode-in-Go](https://github.com/aQuaYi/LeetCode-in-Go/tree/master/Algorithms/0530.minimum-absolute-difference-in-bst)。

```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type res struct {
	minDiff, previous int
}

func getMinimumDifference(root *TreeNode) int {
	res := res{0, 0}
	helper(root, &res)
	return res.minDiff
}

func helper(node *TreeNode, res *res) {
	if node == nil {
		return
	}
	helper(node.Left, res)
	newDiff := diff(res.previous, node.Val)
	if res.minDiff > newDiff {
		res.minDiff = newDiff
	}
	res.previous = node.Val
	helper(node.Right, res)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
```
