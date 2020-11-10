package main

import (
	"Patagonia/src/core"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

func main()  {
	port := flag.Int("p", 8080, "port")
	flag.Parse()
	fmt.Println("[*] Starting Patagonia client")

	var interrupt = false
	for !interrupt {
		err := createConnectWrite(*port)
		if err != nil {
			fmt.Println("[*] (ERR)", err)
			interrupt = true
		}
	}
}

func createConnectWrite(port int) error {
	socketFd, err := core.CreateSocket()
	if err != nil {
		return err
	}

	err = syscall.Connect(socketFd, &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{127, 0, 0, 1},
	})

	fmt.Println("[*] Socket connected on fd", socketFd)

	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[*] > ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	//input := "Hello from the client"
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

	syscall.Close(socketFd)

	return nil
}
