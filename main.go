package main

import (
	"fmt"
	"go_basic/calculator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// NOTE: hello world is commented out to avoid the "no output" error
	//fmt.Println("Hello, World!")
	err := godotenv.Load()
	if err != nil {
		return
	}
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))
}
