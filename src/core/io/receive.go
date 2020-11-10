package io

import (
	"bufio"
	"os"
	"strings"
)

func Receive(fileDescriptor int) (input string, n int, err error) {
	file := os.NewFile(uintptr(fileDescriptor), "pipe")
	reader := bufio.NewReader(file)
	input, err = reader.ReadString('\n')

	if err != nil {
		return input, len(input), err
	}

	input = strings.Replace(input, "\n", "", -1)

	return input, len(input), nil
}
