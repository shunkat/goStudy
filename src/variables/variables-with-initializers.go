package main

import "fmt"
//変数の宣言と同時に初期化として値を導入している。
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

