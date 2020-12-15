package main

import "fmt"

//声明/定义一个接口
type Usb interface {

	Start()
	Stop()
}

type Phone struct {

}

func (p *Phone) Start() {

	fmt.Println("Phone 启动")
}

func (p *Phone) Stop () {

	fmt.Println("Phone 停止")
}

type Camera struct {

}

func (c *Camera) Start() {

	fmt.Println("Camera 启动")
}

func (c *Camera) Stop() {

	fmt.Println("Camera 停止")
}

type Computer struct {

}

func (c *Computer) Working(usb Usb) {

	usb.Start()
	usb.Stop()
}

func main() {

	//测试
	computer := Computer{}
	phone := &Phone{}
	camera := &Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}