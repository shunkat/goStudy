//スライスは値の撮る範囲を指定しなくてはいけませんが既定値もあるので省略するとその既定値で勝手にスライスを作ってくれます。

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)
//ここでは１番目から二番目までのスライスが作成されます
	s = s[:2]
	fmt.Println(s)
//コレだと2つ目から配列の最後の値までが入ったスライスが作成されます。
	s = s[1:]
	fmt.Println(s)
}

