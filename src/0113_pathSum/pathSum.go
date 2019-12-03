package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	right := &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   -2,
		Left:  nil,
		Right: right,
	}
	fmt.Println(pathSum(root, -5))
}

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
