package tools

import (
	"fmt"
	"yfann/leetcode"
)

func PrintListNode(listNode *leetcode.ListNode) {
	for listNode != nil {
		fmt.Printf("%d ->", listNode.Val)
		listNode = listNode.Next
	}
	fmt.Println("")
}
