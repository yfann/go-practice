package main

import (
	"fmt"
	"yfann/patterns"
)

func main() {
	var instance = patterns.GetInstance()
	var instance1 = patterns.GetInstance()
	fmt.Printf(" %T ", instance)
	fmt.Printf(" %T ", instance1)
	fmt.Println(instance1 == instance)
}
