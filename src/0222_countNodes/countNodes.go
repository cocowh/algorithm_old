package _222_countNodes

import "fmt"

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
