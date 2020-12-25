package main

import "fmt"

type use1 struct {

	test int
}

func (u use1)change() int {

	u.test = 1

	return u.test
}

func (u* use1)change1() int {

	u.test = 3

	return u.test
}

func main() {

	u := use1{test : 2}
	u1 := &use1{test : 2}



	fmt.Println(u.change())
	fmt.Println(u1.change1())




}