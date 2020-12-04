package main

import (
	"errors"
	"fmt"
	"time"
)

type ReusableObj struct {
	//空类型
}

type ObjPool struct {

	bufchan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {

	objPool := ObjPool{}
	objPool.bufchan = make(chan *ReusableObj, numOfObj)

	for i := 0; i < numOfObj; i ++ {

		objPool.bufchan <- &ReusableObj{}
	}

	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {

	select {
	case ret := <- p.bufchan:
			return ret, nil
	case <-time.After(timeout):
			return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {

	select {

	case p.bufchan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func main() {

	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {


	}

	for i := 0; i < 10; i ++ {

		if v, err := pool.GetObj(time.Second * 1); err != nil {

		}else {

			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {

				fmt.Println("no obj Reseale")
			}
		}
	}
}