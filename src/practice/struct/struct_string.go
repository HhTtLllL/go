package main

import "fmt"

type anail struct{

	fly string
	round string
}

func (a anail)String() string{

	return fmt.Sprintf(a.fly)
}

/*
	在go语言中, 我们可以通过为一个类型编写名为String的方法, 来自定义该类型的字符串表示形式
	这个String不需要任何参数声明, 但需要有一个string类型的结果声明
*/
func main(){

	a := anail{fly: "brid"}

	fmt.Printf("the test %s", a)
}