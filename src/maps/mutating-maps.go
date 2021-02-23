package main

import "fmt"

func main() {
	m := make(map[string]int)
//mapに値の挿入、と取り出して出力
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
//mapの値の更新と出力
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
//deleteメソッドにより値の削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
//keyに対する要素が存在しているか、もし存在していたらokがそうでなければ
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

