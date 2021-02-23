//for文の書き方は「for　初期化ステートメント	条件式	後処理ステートメントという書き方です。
//後処理のステートメントを中に書くこともできます
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
