package main

import (
	"fmt"
	"net/http" //HTTPプロトコルを利用してくれるパッケージ
)

//httpパッケージを利用してhttp.ResponseWriterという方のwriterを作ってそこにhello worldという文字列を入れています。
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!, %s", request.URL.Path[1:])
	//このように自分の好きな文字をURLを利用して表示させることができるのです。
}
func main() {
	http.HandleFunc("/", handler)
	//サーバーの中に誰かからのアクセスが来たときのハンドリングを指定しています。具体的にはパスを指定することでやっていあmす。。
	///が入力されたときにはhandlerを動かすよという指定です
	http.ListenAndServe(":8080", nil)
	//PCは様々なポートを持っているのでその中で自分自身を示すウェルノウンポート8080で自サーバーを立ち上げる指定をしています
}

//そもそもlocalhostとは自分のPC上にリクエストを送ることができます。つまりlocalhost8080は自分のPCの８０８０番ポートにアクセスしているという意味になります。
