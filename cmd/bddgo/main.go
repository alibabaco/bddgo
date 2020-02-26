package main

import (
	"fmt"
	"flag"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	fmt.Print("Hello bddgo")
}
