package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello world")

	f, err := os.Open("art.ansi")
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1)
	for {
		_, err := f.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		os.Stdout.Write(buffer)
	}
}
