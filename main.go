package main

import (
	"yfann/tools"
)

// func main() {
// 	var instance = patterns.GetInstance()
// 	var instance1 = patterns.GetInstance()
// 	fmt.Printf(" %T ", instance)
// 	fmt.Printf(" %T ", instance1)
// 	fmt.Println(instance1 == instance)
// }

// func main() {
// 	var nums = leetcode.TwoSum([]int{1, 2, 3, 4, 5}, 7)
// 	fmt.Println(nums)
// }

// linkedList test
// func main() {
// 	ll := datastruct.LinkedList{}
// 	ll.InsertAtBeginning(1)
// 	ll.InsertAtBeginning(2)
// 	ll.InsertAtEnd(3)
// 	ll.InsertAtEnd(4)
// 	ll.Display()
// 	ll.Delete(2)
// 	ll.Delete(3)
// 	ll.Display()
// }

// func main() {
// 	list1 := &leetcode.ListNode{Val: 1}
// 	list1.Next = &leetcode.ListNode{Val: 2}
// 	list1.Next.Next = &leetcode.ListNode{Val: 3}

// 	list2 := &leetcode.ListNode{Val: 4}
// 	list2.Next = &leetcode.ListNode{Val: 8}
// 	list2.Next.Next = &leetcode.ListNode{Val: 9}

// 	list3 := leetcode.AddTwoNumbers(list1, list2)
// 	tools.PrintListNode(list1)
// 	tools.PrintListNode(list2)
// 	tools.PrintListNode(list3)

// }

func main() {
	tools.Run()
}
