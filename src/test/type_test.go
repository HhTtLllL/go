package try_test

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type MyInt int64			//定义变量别名

func TestImplicit(t *testing.T) {

	var a int32 = 1
	var b int64

	//b = a   				//错误
	b = int64(a) 			//正确，显示类型转换

	var c MyInt
//	c = b 					//错误，也不支持别名转化

	c = MyInt(b)			//正确，显示类型转换
}

func TestPoint(t *testing.T) {

	a := 1
	aPtr := &a

	aPtr += 1 			//错误,Go中不支持指针算术运算
}

func TestCompareArray(t *testing.T) {

	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]  int{1, 2, 3, 4}

	t.Log(a == b)
	//t.Log(a == c)					//错误，长度不同的数组不能比较
	t.Log(a == d)
}


func TestSwitchMultiCase(t *testing.T) {

	for i := 0; i < 5; i ++ {

		switch i {
		case 0, 2:
			t.Log("Even ")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {

	for i := 0; i < 5; i ++ {

		switch {

		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}

func TestArrayInit(t *testing.T) {

	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {

	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i ++ {

		t.Log(arr3[i])
	}

	for _, e := range arr3 {

		t.Log(e)
	}
}

func TestInitMap(t *testing.T) {

	m1 := map[int]int {1 : 1, 2 : 4, 3 : 9}

	m2 := map[int]int{}			//初始化一个空的map
	m2[4] = 16					//插入一个元素

	m3 := make(map[int]int, 10)//用make创建一个map
}

func TestAccessNotExistingKey(t *testing.T) {


	if v, ok := m["four"]; ok {

		t.Log("four", v)
	}else {

		t.Log("Not existing")
	}

}

func TestTravelMap(t *testing.T) {

	m1 := map[int]int{1: 2, 2: 4, 3: 9}
	for k, v := range m1 {

		t.Log(k, v)
	}
}

func TestMapWithFunValue(t *testing.T) {

	m := map[int]func(op int) int {}

	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op} \
}

func TestMapForSet(t *testing.T) {

	mySet := map[int]bool {}

	//插入元素
	mySet[1] = true

	//判断元素是否存在
	if mySet[1] {

		fmt.Println("exist ")
	}else {

		fmt.Println("is not exist")
	}

	//删除元素
	delete(mySet, 1)

	//求set的长度
	length := len(mySet)
	fmt.Println("len = %d", length)
}

func TestString(t *testing.T) {

	var s string
	t.Log(s)					//初始化为默认零值 ""

	s = "hello"
	s[1] = '3' 					//string 是不可变的byte slice

	s = "\xE4\xB8\xA5"			// 可以存储任何二进制数据

	s = "中"
}

func timeSpent(inner func(op int) int) func (op int) int {

	return func(n int) int {

		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}
}

func slowFun(op int) int {

	time.Sleep(time.Second*1)

	return op
}
func returnMultiValues() (int, int ) {

	return rand.Intn(10), rand.Intn(20)
}

//函数式编程   ---　计算机程序的构造和解释
func Testfn(t *testing.T) {

	a, _ := returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(slowFun)			//timeSpent的返回值是一个新的函数
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {

	ret := 0
	for _, op := range ops {

		ret += op
	}

	return ret
}

func Clear() {

	fmt.Println("clear resources.")
}

func TestDefer(t *testing.T) {

	defer Clear()
	fmt.Println("start")
	panic("err")						//defer依然会执行
	fmt.Println("End")			//不会执行
}


