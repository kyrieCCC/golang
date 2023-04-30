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