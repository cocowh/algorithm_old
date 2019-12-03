# [在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)

### 题目

您需要在二叉树的每一行中找到最大的值。

示例：

```
输入: 
      1
     / \
    3   2
   / \   \  
  5   3   9 
输出: [1, 3, 9]
```

### 解法

深度遍历，递归，记录层次，参考[二叉树的层次遍历](https://github.com/cocowh/algorithm/tree/master/0102_levelOrder)。

也可借助栈广度遍历。

```
func helper(node *TreeNode, res []int, level int) (int, []int) {
	if node == nil {
		return level, res
	}
	if level >= len(res) { //拓深，初始化为最小值
		res = append(res, ^int(^uint(0)>>1))
	}
    //当前节点值与当前层次最大值比较
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
	if len(res) > 0 && maxLevel > len(res) { //最后多加一层
		res = res[:len(res)-1]
	}
	return res
}
```
