package main

import "golang.org/x/tour/pic"

//まずはdx,dyというint型の引数をとって二次元配列を返すような関数を宣言しています。
func Pic(dx, dy int) [][]uint8 {
	//まず返すためにimageという値を作っています。　その時の行の数はdyです。
	image := make([][]uint8, dy)
	//ここではまず行ごとに中身を作っていくために縦のfor文くりかえしで処理をするよという宣言します。
	for i := 0; i < dy; i++ {
		//その中でi行目には長さdxの配列を代入するよ
		image[i] = make([]uint8, dx)
		//更にその各行のなかでj列目の中身を作る処理をしています。
		for j := 0; j < dx; j++ {
			image[i][j] = uint8((j + 22222) % (i + 12))
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
