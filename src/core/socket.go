package core

import (
	"errors"
	"fmt"
	"log"
	"syscall"
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


	data := make([]byte, 4096)
	n, err := syscall.Read(nfd, data)
	if err != nil {
		return err
	}

	log.Printf("Received %d bytes", n)
	log.Printf("Content: %s", data)


	data = []byte("Hello world from the server")
	n, err = syscall.Write(nfd, data)
	if err != nil {
		return err
	}

	log.Printf("Wrote %d bytes", n)
	log.Printf("Content: %s\n", data)

	syscall.Close(nfd)

	return nil
}
