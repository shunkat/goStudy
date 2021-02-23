// poiterを通じてstrucrの中の値にアクセスする事もできます。

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
//ここではpにvのポインタを代入して、その下でポインタ越しに中身のxに値を表示しています。
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

