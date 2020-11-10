package main

import (
	"Patagonia/src/core/io"
	"Patagonia/src/core/net"
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
	socketFd, err := net.CreateSocket()
	if err != nil {
		return err
	}

	err = net.Connect(socketFd, port)

	fmt.Println("[*] Socket connected on fd", socketFd)

	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[*] > ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	//input := "Hello from the client"
	n, err := io.Send(socketFd, []byte(input))

	log.Printf("Wrote: %d bytes", n)
	log.Printf("Wrote: %s\n", input)


	data, n, err := io.Receive(socketFd)
	if err != nil {
		return err
	}

	log.Printf("Received %d bytes", n)
	log.Printf("Content: %s", data)

	syscall.Close(socketFd)

	return nil
}
