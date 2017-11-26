package main

import (
	"flag"
	"fmt"
)

var filepath = flag.String("file", "empty", "Path to quiz file")

func main() {
	flag.Parse()
	fmt.Printf("file path is %s\n", *filepath)
}
