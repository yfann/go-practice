package main

import (
	"fmt"
	"yfann/leetcode"
)

// func main() {
// 	var instance = patterns.GetInstance()
// 	var instance1 = patterns.GetInstance()
// 	fmt.Printf(" %T ", instance)
// 	fmt.Printf(" %T ", instance1)
// 	fmt.Println(instance1 == instance)
// }

func main() {
	var nums = leetcode.TwoSum([]int{1, 2, 3, 4, 5}, 7)
	fmt.Println(nums)
}
