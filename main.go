package main

import "yfann/datastruct"

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

func main() {
	ll := datastruct.LinkedList{}
	ll.InsertAtBeginning(1)
	ll.InsertAtBeginning(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.Display()
	ll.Delete(2)
	ll.Delete(3)
	ll.Display()
}
