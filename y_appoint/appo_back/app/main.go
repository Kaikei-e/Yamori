package main

import (
	"appoint/inittialize"
	"fmt"
)

func init() {
	inittialize.Check()
}

func main() {
	fmt.Printf("hello world")
}
