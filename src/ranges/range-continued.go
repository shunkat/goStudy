package main

import "fmt"

func main() {
	pow := make([]int, 10)
	//この例ではindexしか必要ない場合なので値を使わずにインデックスだけで繰り返している感じですね。配列を繰り返しの要素でしか扱わないときはこの形ですね
	//具体的にはからの配列を用意して、その中に２の要素番号乗した値を入れています。
	for nu := range pow {
		pow[nu] = 1 << uint(nu) // == 2**i
	}
	//もしインデックスが必要ない場合は_を代入することで捨てることができます。コレシないとiとして扱われてしまうということですね。
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

