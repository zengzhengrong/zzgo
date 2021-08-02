package zlink

import "fmt"

// LinkNode is 单向列表
type LinkNode struct {
	Value interface{}
	Next  *LinkNode
}

//Reverse is reverse LinkNode
func (head *LinkNode) Reverse() *LinkNode {
	cur := head
	var pre *LinkNode = nil
	for cur != nil {

		pre, cur, cur.Next = cur, cur.Next, pre

	}
	return pre
}

//Check is check
func Check(head *LinkNode) {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next

	}

}

/// 反转链表
func ReverseNode(head *LinkNode) *LinkNode {
	//  先声明两个变量
	//  前一个节点
	var preNode *LinkNode
	var nextNode *LinkNode
	preNode = nil
	//  后一个节点

	nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.Next
		//  将头节点指向前一个节点
		head.Next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}
