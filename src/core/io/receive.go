package io

import (
	"bufio"
	"os"
)

func Receive(fileDescriptor int) (input []byte, n int, err error) {
	file := os.NewFile(uintptr(fileDescriptor), "pipe")
	reader := bufio.NewReader(file)
	data := make([]byte, 4096)
	n, err = reader.Read(data)

	if err != nil {
		return input, len(input), err
	}

	return data, n, nil
}
