package main

import (
	"fmt"

	"github.com/dimasalakhov/go-submodules-test/a"
)

func main() {
	fmt.Println("calling " + a.Name + " from package B")
}
