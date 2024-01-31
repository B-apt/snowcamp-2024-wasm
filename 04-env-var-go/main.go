package main

import (
	"fmt"
	"os"
)

func main() {
	f := os.Getenv("EFILE")

	file, err := os.Create(f)
	if err != nil {
		fmt.Println("😡", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("👋 Hello, World! 🌍")
	if err != nil {
		fmt.Println("😡", err)
		return
	}

	data, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("😡", err)
	}

	fmt.Print(string(data))

}
