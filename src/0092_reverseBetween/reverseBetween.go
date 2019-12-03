package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	list := &ListNode{}
	node := list
	for i := 0; i < 5; i++ {
		temp := &ListNode{
			Val:  i + 1,
			Next: nil,
		}
		node.Next = temp
		node = node.Next
	}
	//temp := &ListNode{
	//	Val:  3,
	//	Next: nil,
	//}
	//node.Next = temp
	//node = node.Next
	//temp = &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	//node.Next = temp
	//node = node.Next
	newList := reverseBetween(list.Next, 1,4)
	fmt.Println("------")
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
	fmt.Println("------")
	for newList != nil {
		fmt.Println(newList.Val)
		newList = newList.Next
	}
	fmt.Println("------")
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	headNode,tailNode,count,stack := head,head,0,[]*ListNode{}
	for count < n {
		count++
		if count >= m && count <= n {
			stack = append(stack,tailNode)
		}
		if count < m - 1 {
			headNode = headNode.Next
		}
		if count < n + 1 {
			 tailNode = tailNode.Next
		}
	}
	nodeNum := len(stack)
	for i := nodeNum - 1; i >= 0; i-- {
		headNode.Next = stack[i]
		if m == 1 && i == nodeNum - 1 {  //特殊点从开头开始反转，链表头改变
			head = headNode.Next
		}
		headNode = headNode.Next
	}
	headNode.Next = tailNode
	return head
}