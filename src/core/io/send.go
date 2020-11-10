package io

import (
	"os"
)

func Send(fileDescriptor int, data []byte) (n int, err error) {
	file := os.NewFile(uintptr(fileDescriptor), "conn")
	n, err = file.Write(data)
	if err != nil {
		return n, err
	}

	return n, nil
}
