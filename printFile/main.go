package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// b, _ := os.ReadFile(os.Args[1])
	// fmt.Println(string(b))

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
