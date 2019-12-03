# [完全二叉树的节点个数](https://leetcode-cn.com/problems/count-complete-tree-nodes/)

### 题目

给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

```
输入: 
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
```

### 解法

递归左右子树交换。

```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//遍历解法O(N)
func countNodesByRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodesByRecursive(root.Left) + countNodesByRecursive(root.Right) + 1
}

//算术解法
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelL, levelR := countLevel(root.Left), countLevel(root.Right)
	fmt.Println(levelL,levelR)
	if levelL == levelR {
		return countNodes(root.Right) + 1 << uint(levelL)
	} else {
		return countNodes(root.Left) + 1 << uint(levelR)
	}
}

func countLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}
```
