package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := readFile("test.txt")

	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}

	fmt.Println("Content:")
	fmt.Println(data)
}

// Le type 'error' sert a retourner des erreurs en utilisant les retours multiples
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		//return "", errors.New("empty content")
		return "", fmt.Errorf("empty content (file=%v)", filename)
	}

	return string(data), nil
}
