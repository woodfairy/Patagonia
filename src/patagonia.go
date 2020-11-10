package main

import (
	"Patagonia/src/core"
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("p", 8080, "port")
	flag.Parse()
	fmt.Println("[*] (VERBOSE) Creating and binding socket")
	socketFd, err := core.CreateAndBind(*port)
	if err != nil {
		fmt.Println("[*] (ERR)", err)
	}

	var interrupt = false

	for !interrupt {
		fmt.Println("[*] (VERBOSE) Listening and accepting connection")
		err = core.ListenAndAccept(socketFd)
		if err != nil {
			fmt.Println("[*] (ERR)", err)
			interrupt = true
		}
	}
}
