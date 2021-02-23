//mapは配列やスライスの中でキーと値を関連付けるものです。自分の考え方では配列を辞書に変えるようなイメージだという認識ですね
package main

import "fmt"

//ここで何が行われているかをもう一度確認するとtypeメソッドは既存の型に別名をつける事のできるものなので、vertexという別名をstruct型につけている、という感じです。
//その中に、Lat.Longというfloat64型の値を入れているのがここでの処理です。
//以上復習でした。
type Vertex struct {
	Lat, Long float64
}

//本来の書き方はvar languages map[string]string = map[string]string{"go":"golang", "rb":"ruby", "js":"javascript"}のように、最初にmap[keyの型]valueの型という形で変数を宣言して、そこに同じ型の値を後ろから代入するというものです。
//ここでは前半の変数の方の宣言までは一緒ですが、その後の値の代入をgoの組み込みのmake関数を使ってやっています。
var m map[string]Vertex

func main() {
	//make関数でstring型のkeyを持ちvertex(struct)型のvalueをもつ、空の配列mを作っています。
	m = make(map[string]Vertex)
	//make関数は一つずつ値を入れていかないといけませんね、今回はbell labsというkeyとVertexの値	40.68433, -74.39967,を結びつけています。
	//mapはスライスと違って自分でラベルをつけることができるのが特徴ですね
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
