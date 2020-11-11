package net

import (
	"Patagonia/src/core/io"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"syscall"
	"time"
)

func CreateSocket() (fileDescriptor int, error error) {
	socketFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return socketFd, err//errors.New("socket creation failed")
	}

	return socketFd, nil
}

func bindSocket(fileDescriptor int, address syscall.Sockaddr) error {
	err := syscall.Bind(fileDescriptor, address)
	if err != nil {
		return err//errors.New("socket bind failed")
	}

	return nil
}

func listen(fileDescriptor int, backlog int) error {
	err := syscall.Listen(fileDescriptor, backlog)
	if err != nil {
		return errors.New("socket listening failed")
	}

	return nil
}

func accept(fileDescriptor int) (nfd int, sockAddr syscall.Sockaddr, err error) {
	nfd, sockAddr, err = syscall.Accept(fileDescriptor)
	if err != nil {
		return nfd, sockAddr, err//errors.New("socket accept failed")
	}

	return nfd, sockAddr, nil
}

func Connect(fileDescriptor int, port int) error {
	syscall.Connect(fileDescriptor, &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{127, 0, 0, 1},
	})

	return nil
}

func CreateAndBind(port int) (socketFd int, error error) {
	fmt.Println("[*] (VERBOSE) Creating raw socket")
	socketFd, err := CreateSocket()
	if err != nil {
		return socketFd, err
	}
	fmt.Println("[*] (VERBOSE) Socket created succesfully with fd", socketFd)

	sockAddress := &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{127, 0, 0, 1},
	}

	fmt.Println("[*] (VERBOSE) Binding socket with fd", socketFd, "on", sockAddress.Addr, ":", sockAddress.Port)
	err = bindSocket(socketFd, sockAddress)
	fmt.Println("[*] (VERBOSE) Socket bound succesfully")

	if err != nil {
		return socketFd, err
	}
	
	return socketFd, nil
}

func ListenAndAccept(socketFd int) error {
	err := listen(socketFd, 0x41)
	if err != nil {
		return err
	}

	nfd, sockAddr, err := accept(socketFd)

	if err != nil {
		return err
	}

	fmt.Println("[*] Socket with fd", socketFd, "accepting connection on fd", nfd, sockAddr)

	start := time.Now()

	input, n, err := io.Receive(nfd)
	if err != nil {
		return err
	}

	log.Printf("Received %d bytes", n)
	log.Printf("Content: %s", input)

	//data := []byte("<html><head><title>Patagonia</title></head><body><h1>Patagonia</h1><p>It's working! Hello world from the Patagonia server<p></body><html>")

	data, err := ioutil.ReadFile("index.html")

	if err != nil {
		return err
	}

	n, err = io.Send(nfd, data)
	if err != nil {
		return err
	}

	duration := time.Since(start)

	log.Printf("Wrote %d bytes", n)
	log.Printf("Content: %s\n", data)
	log.Printf("Duration: %dms\n", duration.Microseconds())

	syscall.Close(nfd)

	return nil
}

func ReceiveAndServe(port int) error {
	fmt.Println("[*] (VERBOSE) Creating and binding socket")
	socketFd, err := CreateAndBind(port)
	if err != nil {
		return err
	}

	var interrupt = false

	for !interrupt {
		fmt.Println("[*] (VERBOSE) Listening and accepting connection")
		err = ListenAndAccept(socketFd)
		if err != nil {
			interrupt = true
			return err
		}
	}

	return nil
}
