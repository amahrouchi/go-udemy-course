package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//r := strings.NewReader("Hello Gophers! Readers are awesome")
	r, err := os.Open("test.txt")
	buf, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}
