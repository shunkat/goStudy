package main

import (
	"fmt"
	"time"
)

func oneSecondWait() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("1second passed")
}
func twoSecondWait() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2second passed")
}
func threeSecondWait() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("3second passed")
}
func main() {
	fmt.Println(time.Now())
	go threeSecondWait()
	oneSecondWait()
	twoSecondWait()
	fmt.Println(time.Now())
	//これがgoroutineの良い例なのですが、goroutineに入れた処理は他の処理が全てお話あってしまうと実行されずに終わってしまうので注意、そこで解決策としてchannelというものがあります。
}
