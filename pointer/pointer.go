package pointer

import (
	"fmt"
	"unsafe"
)

// NOTE: uint16 2bytes
var ui1 uint16
var ui2 uint16

// NOTE: pointer
var p1 *uint16

// NOTE: double pointer
var pp1 **uint16 = &p1

func DefaultUis() {
	fmt.Printf("defaultUi1 address: %p\n", &ui1)
	fmt.Printf("defaultUi2 address: %p\n", &ui2)
}

func P1WithUi1() {
	p1 = &ui1
	// NOTE: p1 は ui1 の先頭のアドレスを指す
	fmt.Printf("p1 address: %p\n", p1)
}

func SizeOfP1() {
	// NOTE: p1 のサイズは8bytes になる
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
}

func DoublePointer() {
	// NOTE: pp1 は p1 のアドレスを指す
	fmt.Printf("pp1 address: %p\n", pp1)
}

func DereferencePointer() {
	// NOTE: p1 は ui1 の先頭のアドレスを指す
	fmt.Printf("p1 Dereference: %d\n", *pp1)
	// NOTE: ui1 の値を取得する
	fmt.Printf("pp1 Dereference: %d\n", **pp1)
}

func ChangeValueUi1(num uint16) {
	// NOTE: p1 は ui1 の先頭のアドレスを指す
	**pp1 = num
	fmt.Printf("value of ui1: %v\n", ui1)
}

// NOTE: ShadowingSample は変数のスコープを確認するためのサンプル
func ShadowingSample() {
	ok, result := true, "A"
	if ok {
		// NOTE: result を := とすると if のスコープでしか使えない
		result = "B"
		fmt.Println(result)
	} else {
		result = "C"
		fmt.Println(result)
	}
	fmt.Println(result)
}
