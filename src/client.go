package main

import (
	"Patagonia/src/core"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

func main()  {
	fmt.Println("[*] Starting Patagonia client")

	var interrupt = false
	for !interrupt {
		err := createConnectWrite()
		if err != nil {
			fmt.Println("[*] (ERR)", err)
			interrupt = true
		}
	}
}

func createConnectWrite() error {
	socketFd, err := core.CreateSocket()
	if err != nil {
		return err
	}

	err = syscall.Connect(socketFd, &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	})

	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[*] > ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	file := os.NewFile(uintptr(socketFd), "pipe")
	_, err = file.Write([]byte(input))
	if err != nil {
		return err
	}

	log.Printf("Wrote: %s\n", input)

	data := make([]byte, 4096)
	n, err := syscall.Read(socketFd, data)
	if err != nil {
		return err
	}

	log.Printf("Received %d bytes", n)
	log.Printf("Content: %s", data)


	return nil
}
