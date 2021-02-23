//deferで実行されるのは逆順です。筒にいれてそこから取り出すイメージですね。【意味不明】
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
//実際に使うときは重めの処理をやるときに先にやらなきゃいけない処理を進めてから最後にそれをやるようなイメージです
	fmt.Println("done")
}

