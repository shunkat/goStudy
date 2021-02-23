package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
//コレは先程の例より簡略化できるよということを示す例です。
//vertexという宣言をしている下のところでは再度宣言をしないでも自動で推定してくれます。
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

