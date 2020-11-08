package main

import (
	"Patagonia/src/core"
	"fmt"
)

func main() {
	fmt.Println("[*] (VERBOSE) Creating and binding socket")
	err := core.CreateAndBind()
	if err != nil {
		fmt.Println(err)
	}
}
