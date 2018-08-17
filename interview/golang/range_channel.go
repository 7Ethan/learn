package main

import (
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！	问题出在这!!!!!!!!!
	// ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem //-> 因为ch是无缓冲的,只能等待chennel的使用者讲数据读出.
			//--->如果set.s里面有很多数据,那么可能是这个for 迭代一直阻塞再 ch <- elem 这个操作上
			// 其他进程无法获取 set.s的资源(一直被锁上)
			//---->并且如果 set.s 的资源在堆上分配,对这个channel的调用这没有完全将数据取出,
			// main 进程瑞出可能会造成goroutine泄露.
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2", "3", "4"},
		// s: []interface{}{"1"},
	}
	v := <-th.Iter()
	// fmt.Printf("%s%v", "ch", v)
	var _ = v
	for {

	}
}
