package main

import (
	"Patagonia/src/core/net"
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("p", 8080, "port")
	flag.Parse()
	err := net.ReceiveAndServe(*port)
	if err != nil {
		fmt.Println("[*] (ERR)", err)
	}
}
