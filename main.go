package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_basic/calculator"
	"go_basic/pointer"
	"go_basic/variables"
	"os"
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
	fmt.Println(variables.Hello("John"))
	fmt.Println(variables.GetOsName(variables.Windows))
	pointer.DefaultUis()
	pointer.P1WithUi1()
	pointer.SizeOfP1()
	pointer.DoublePointer()
	pointer.DereferencePointer()
	pointer.ChangeValueUi1(10)
	pointer.ShadowingSample()
}
