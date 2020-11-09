package main

import (
	"Patagonia/src/core"
	"fmt"
	"log"
	"os"
	"syscall"
)

func main()  {
	socketFd, err := core.CreateSocket()
	if err != nil {
		fmt.Println(err)
	}

	err = syscall.Connect(socketFd, &syscall.SockaddrInet4{
		Port: 8081,
		Addr: [4]byte{127, 0, 0, 1},
	})

	if err != nil {
		fmt.Println(err)
	}

	file := os.NewFile(uintptr(socketFd), "pipe")
	_, err = file.Write([]byte(`Hello World from the client`))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("wrote")

	data := make([]byte, 27)
	n, err := syscall.Read(socketFd, data)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Read %d bytes", n)
	log.Printf("Content: %s", data)

	_, _ = fmt.Scanln()


}
