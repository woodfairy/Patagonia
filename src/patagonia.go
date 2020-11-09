package main

import (
	"Patagonia/src/core"
	"fmt"
)

func main() {
	fmt.Println("[*] (VERBOSE) Creating and binding socket")
	socketFd, err := core.CreateAndBind()
	if err != nil {
		fmt.Println("[*] (ERR)", err)
	}

	fmt.Println("[*] (VERBOSE) Listening and accepting connection")
	err = core.ListenAndAccept(socketFd)
	if err != nil {
		fmt.Println("[*] (ERR)", err)
	}
}
