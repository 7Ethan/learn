package main

import (
	"fmt"
	_ "time"
)

func main() {
	ping := make(chan bool)

	message := make(chan struct{})

	//go func(){
	//	for i := 0; i < 10;i++{
	//		<-ping
	//		fmt.Print(i)
	//		pong <- true
	//	}
	//}()
	//
	//go func(){
	//	for i := 0;i < 10;i++{
	//		ping <- true
	//		fmt.Printf("%c",'A' + i)
	//		<- pong
	//	}
	//	message <- struct{}{}
	//}()
	//
	////ping <- true
	//<- message

	go func() {
		for i := 0; i < 10; i++ {
			<-ping
			fmt.Print(i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ping <- true
			fmt.Printf("%c", 'A'+i)
		}
		message <- struct{}{}
	}()
	<-message
}

// 译者注 ：Action 结构体需要添加一个 next 字段，否则下面代码的部分操作不成立
type Action struct {
	name     string
	metadata Meta
	// Probably some other data here...
	// 这里或许还有一些其他数据
	next *Action
}

type ActionHistory struct {
	top  *Action
	size int
}

//func (history *ActionHistory) Undo() *Action {
//	topAction := history.top
//	if topAction != nil {
//		history.top = topAction.next
//	} else if topAction.next == nil {
//		history.top = nil
//	}
//	historyAction.size--
//	return topAction
//}
