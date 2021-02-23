vimpackage main

import "fmt"

func main() {
	var i, j int = 1, 2
	//varの代わりに:=を使うことで暗黙的に型宣言と初期化をすることができる
	//ただし関数の外では:=を使うことはできない。するとnon-declaration statement outside function bodyというエラーが出る
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

