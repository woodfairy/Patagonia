package core

import (
	"errors"
	"fmt"
	"syscall"
)

func createSocket() (fileDescriptor int, error error) {
	socketFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return socketFd, errors.New("socket creation failed")
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
		return nfd, sockAddr, errors.New("socket accept failed")
	}

	return nfd, sockAddr, nil
}

func CreateAndBind() (socketFd int, error error) {
	fmt.Println("[*] (VERBOSE) Creating raw socket")
	socketFd, err := createSocket()
	if err != nil {
		return socketFd, err
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

	_, _ = fmt.Scanln()

	return nil
}
