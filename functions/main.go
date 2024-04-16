package functions

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// add defer function
func funcDefer() {
	defer fmt.Println("defer final finish")
	defer fmt.Println("defer semi finish")
	fmt.Println("hello world") // defer は関数の最後に実行されるため、この行の後に実行される ↑
}

// add trim extension
func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, filepath.Ext(f)))
	}
	return out
}

// add extension function
func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
}

// add file checker function
func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func multiply() func(int) int {
	return func(i int) int {
		return i * 100
	}
}

// count up function
func countUp() func(int) int {
	count := 0
	return func(i int) int {
		count += i
		return count
	}
}

func main() {
	funcDefer() // defer 関数の実行
	// add slice
	files := []string{"test1.txt", "test2.txt", "test3.txt"}
	fmt.Println(trimExtension(files...))

	// add file checker
	name, err := fileChecker("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	// Anonymous function (即時関数)
	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)

	// Closure
	i = 2
	f1 := func(i int) int {
		return i + 1
	}

	fmt.Println(f1(2))

	// Closure2
	f2 := func(file string) string {
		return file + ".csv"
	}

	// add extension
	addExt(f2, "test")

	// multiply
	f3 := multiply()
	fmt.Println(f3(10))

	// count up
	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(i)
		fmt.Printf("%v\n", v)

	}
}
