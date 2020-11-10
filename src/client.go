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
	socketFd, err := core.CreateSocket()
	if err != nil {
		fmt.Println(err)
	}

	err = syscall.Connect(socketFd, &syscall.SockaddrInet4{
		Port: 8080,
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

	reader := bufio.NewReader(os.Stdin)

	var interrupt = false

	for !interrupt {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		_, err = file.Write([]byte(input))
		if err != nil {
			fmt.Println(err)
			interrupt = true
		}

		log.Printf("Write: %s", input)

		data := make([]byte, 4096)
		n, err := syscall.Read(socketFd, data)
		if err != nil {
			fmt.Println(err)
			interrupt = true
		}

		log.Printf("Read %d bytes", n)
		log.Printf("Content: %s", data)


	}


	_, _ = fmt.Scanln()

}
