package main

import "fmt"

func main() {

	//第二个参数是5, 指名了该切片的长度
	s1 := make([]int, 5)

	fmt.Println("The length of s1: ", len(s1))
	fmt.Println("The capacity of s1: ", cap(s1))
	fmt.Println("The valuie of s1: ", s1)


	// 第二个参数是 s2的长度, 第三个参数是 s2的容量
	s2 := make([]int, 5, 8)

	fmt.Println("The length of s2: ", len(s2))
	fmt.Println("The cap of s2: ", cap(s2))
	fmt.Println("The value of s2", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]

	fmt.Println("The length of s4: ", len(s4))
	fmt.Println("The capacity of s4: ", cap(s4))
	fmt.Println("The value of s4: ", s4)
}

