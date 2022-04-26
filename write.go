package main

import (
	"fmt"
	"io"
	"os"
)

func write() string {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	var fileContent string

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fileContent = string(data[:n])
	}
	return fileContent
}
