package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//下のように最初に変数的な簡単なステートメントを書くこともできます。このステートメントはifのスコープの中でだけ有効だということに注意してください。
	//Powは第一引数を第二引数で累乗した値を返すmath内の関数です。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	//ここでもしvとかを使おうと思うとスコープの外なのでエラーです。もし使いたかったらifの外に書いてください。
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
