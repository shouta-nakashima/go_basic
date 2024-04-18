package conditionals

import (
	"fmt"
	"time"
)

type item struct {
	price float32
}

func main() {
	// if文
	a := -1
	if a == 0 {
		fmt.Println("ZERO")
	} else if a > 0 {
		fmt.Println("POSITIVE")
	} else {
		fmt.Println("NEGATIVE")
	}
	// for文
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for 文の無限ループ
	//for {
	//	fmt.Println("infinite loop")
	//	time.Sleep(1 * time.Second)
	//}

	var i int
	for {
		if i > 3 { // i が 3 より大きい場合はループを抜ける
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}

	// for と switch の組み合わせ
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			break loop // loop ラベルを指定してループを抜ける
		default:
			fmt.Printf("%v ", i)
		}
		fmt.Printf("\n")
	}
	// for range
	items := []item{
		{price: 100.},
		{price: 200.},
		{price: 300.},
	}
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("%+v\n", items) // items の中身は変わらない(for range はコピーを渡すため)
	for i := range items {
		items[i].price *= 1.1 // items の中身が変わる(items の値を変更するためには index を使う)
	}
	fmt.Printf("%+v\n", items) // [110 220 330]
}
