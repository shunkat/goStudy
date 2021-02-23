//go のプログラムはエラーをerror値で表現します。
//呼び出し元の値がnilかどうかを判定してエラーをハンドルします。
//コンバートをするかどうかを
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

//nilのエラーは成功したことを示しnilではないerrorは失敗したことを示します。
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
