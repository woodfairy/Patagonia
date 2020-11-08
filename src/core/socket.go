package core

import (
	"errors"
	"fmt"
	"syscall"
)

func createSocket() (fileDescriptor int, error error) {
	socketFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return 0x0, errors.New("socket creation failed")
	}

	return socketFd, nil
}

func bindSocket(fileDescriptor int, address syscall.Sockaddr) error {
	err := syscall.Bind(fileDescriptor, address)
	if err != nil {
		return errors.New("socket bind failed")
	}

	return nil
}

func CreateAndBind() error {
	fmt.Println("[*] (VERBOSE) Creating raw socket")
	socketFd, err := createSocket()
	if err != nil {
		return err
	}
	fmt.Println("[*] (VERBOSE) Socket created succesfully with fd", socketFd)

	sockAddress := &syscall.SockaddrInet4{
		Port: 8081,
		Addr: [4]byte{127, 0, 0, 1},
	}

	fmt.Println("[*] (VERBOSE) Binding socket with fd", socketFd)
	err = bindSocket(socketFd, sockAddress)
	fmt.Println("[*] (VERBOSE) Socket bound succesfully")

	if err != nil {
		return err
	}
	
	return nil
}
