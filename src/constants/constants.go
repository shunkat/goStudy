//定数ももちろんあります
//文字、文字列、boolean、数値でのみ使えまうす。
//ただし:=を使って宣言することはできません。
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

