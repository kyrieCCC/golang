package main

type ListNode struct {
	Val int
    Next *ListNode
 }

func main(){
	var ln ListNode
	ln.Val = 1
	ln.Next = &ln
	removeElements(&ln, 1)
}

func removeElements(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Next: head}
	for node := newNode; node.Next != nil ;{
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return newNode.Next
}	
//循环解法
//首先遍历每一个元素得值，并找出指定的元素
//找到指定的元素后进行一个跳过操作
//即将next = next.next
//最后返回一个链表