package mapDir

import "fmt"

var m1 map[string]int // key が string 型で value が int 型の mapDir

func mapFn() {
	m2 := map[string]int{}
	fmt.Printf("%v %v \n", m1, m1 == nil) // var で初期化すると nil
	fmt.Printf("%v %v \n", m2, m2 == nil)
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 30
	fmt.Printf("%v %v\n", m2, len(m2), m2["A"])
	delete(m2, "A")
	fmt.Printf("%v %v\n", m2, len(m2), m2["A"]) // 存在しないキーの値を参照すると 0 が返る
	v, ok := m2["A"]                            // key("A") に対しての v == value, ok == bool が取得できる
	fmt.Printf("\"%v %v\n", v, ok)              // "A" は存在しないため、0 false となる
	v, ok = m2["C"]
	fmt.Printf("\"%v %v\n", v, ok) // "C" は存在しているため、0 true となる

	for k, v := range m2 {
		fmt.Printf("%v  %v\n", k, v) // key value を一つずつ出力する
	}
}
