package main

import "fmt"
//この記法によってVertexというものを構造体として扱うということが宣言されます。
//またその中の構造も同時に書くことができます。
//この例ではint型の変数ｘとyを持つよということが宣言されました。
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

