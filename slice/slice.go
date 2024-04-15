package slice

import "fmt"

var a1 [3]int               // array [0, 0, 0]
var a2 = [3]int{10, 20, 30} // array [10, 20, 30]

func a3() {
	a4 := [...]int{10, 20, 30}
	fmt.Printf("%v %v %v\n", a1, a2, a4) // array [10, 20, 30] 3
}

func s1() {
	var s2 []int
	s3 := []int{}
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %[1]T %[1]v %v %v\n", s3, len(s3), cap(s3))
	fmt.Println(s2 == nil) // var で宣言されたスライスは nil で初期化される
	fmt.Println(s3 == nil)
	s2 = append(s2, 1, 2, 3) // スライスに要素を追加
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	s4 := []int{1, 2, 3}
	s3 = append(s3, s4...) // スライスにスライスを追加
	fmt.Printf("s3: %[1]T %[1]v %v %v\n", s3, len(s3), cap(s3))

	s5 := make([]int, 3, 5) // スライスの初期化
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	s5 = append(s5, 1, 2, 3)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	s6 := make([]int, 4, 6) // スライスの初期化
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	s7 := s6[1:3]                                               // 別のスライスに参照を代入
	s7[1] = 100                                                 // s6[2] にも影響
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6)) // s6 と s7 は同じ配列を参照しているため、s7 の変更が s6 にも影響
	fmt.Printf("s7: %[1]T %[1]v %v %v\n", s7, len(s7), cap(s7))

	s7Copy := make([]int, len(s7[1:3])) // スライスのコピー
	copy(s7Copy, s7[1:3])
	s7Copy[1] = 200
	fmt.Printf("s7Copy: %[1]T %[1]v %v %v\n", s7Copy, len(s7Copy), cap(s7Copy))

	fs6 := s6[1:3:3] // スライスの初期化 s6 の 1 番目から 3 番目までのスライスを作成 s6 の 1~2 番目までのスライスを作成
	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))

	fs6[0] = 100
	fs6[1] = 200
	fs6 = append(fs6, 300)
	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6)) // [100, 200, 300]
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))     // [0, 100, 200, 0]
	s6[3] = 400
	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6)) // [100, 200, 300]
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))     // [0, 100, 200, 400]
}
