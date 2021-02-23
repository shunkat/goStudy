//sliceやmapを一つずつ反復処理するためには、rangeを使って繰り返します。１つ目の引数はindex２つ目の値はインデックスの中の要素のコピーです。他の言語で言えばfor eachのようなものですね。
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//上で定義している配列iに対して中にある要素数分だけ繰り返して、それぞれの要素をvとしてまたindexをiとして処理を繰り返している感じですね
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

