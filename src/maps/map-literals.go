package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
//
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
//mapはstructリテラルと似ていますが、違うのはkeyが入る点です。
//あと出力したときにmapということを示すために　、mapという文字列が最初に明示されます。
