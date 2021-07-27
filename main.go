package main

import (
	"fmt"
	"github.com/brady-wang/go-tools/hashx"
)

func main() {
	var str string
	str = "hello"
	str = hashx.Sha256(str)
	fmt.Println(str)
}
