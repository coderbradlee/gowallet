package main

import (
// "fmt"
)

func main() {
	var choice []int
	for i := 0; i < 15; i++ {
		choice = append(choice, i)
	}
	combine(choice, 14)
	// choice := []string{"a", "b", "c", "d", "e"}
	// {
	// 	newchoice := rmstring(choice, "a")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "b")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "c")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "d")
	// 	fmt.Println(newchoice)
	// }
	// {
	// 	newchoice := rmstring(choice, "e")
	// 	fmt.Println(newchoice)
	// }
	var results []string
	permutate(choice, results)
	// fmt.Println("imtoken address:", ethaddress)
	// fmt.Printf("\n")
	// testbtcsign()
	// testfunc()
	// var results []int
	// getReward(10, results)
	// left := []int{1, 4, 8}
	// right := []int{2, 10, 20, 33, 55, 66, 77}
	// ret := mergeSort(left, right)
	// fmt.Println(ret)
	// test := []int{2, 9, 4, 10, 3, 20, 5, 33, 1, 88, 11, 25, 6}
	// ret := merges(test)
	// fmt.Println(ret)
	// to_sort := []int{3434, 3356, 67, 123, 111, 890}
	// sorted := merges(to_sort)
	// fmt.Println(sorted)
	// dv(8, results)
	// fmt.Println("===================")
	// dv(15, results)
	// dv(16, results)
	// dv(15, results)
	// test()
	// testipfs()
}
