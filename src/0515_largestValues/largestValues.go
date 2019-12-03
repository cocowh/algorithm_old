package _515_largestValues

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, res []int, level int) (int, []int) {
	if node == nil {
		return level, res
	}
	if level >= len(res) {
		res = append(res, ^int(^uint(0)>>1))
	}
	if node.Val > res[level] {
		res[level] = node.Val
	}
	levelL, res := helper(node.Left, res, level+1)
	levelR, res := helper(node.Right, res, level+1)
	if levelL > levelR {
		return levelL, res
	}
	return levelR, res
}
func largestValues(root *TreeNode) []int {
	maxLevel, res := helper(root, []int{}, 0)
	if len(res) > 0 && maxLevel > len(res) {
		res = res[:len(res)-1]
	}
	return res
}
