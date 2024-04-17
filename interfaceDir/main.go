package interfaceDir

import "fmt"

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

// vehicle methods
func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed

}

// bycycle methods

func (b *bycycle) speedUp() int {
	b.speed += 5 * b.humanPower
	return b.speed
}

func (b *bycycle) speedDown() int {
	b.speed -= 3 * b.humanPower
	return b.speed
}

// any type
var i1 interface{}
var i2 any

func checkType(i any) {
	switch v := i.(type) {
	case nil:
		fmt.Println("Type is nil")
	case int:
		fmt.Printf("Type is int, value is %v\n", v)
	case string:
		fmt.Printf("Type is string, value is %v\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

func speedUpAndDown(c controller) {
	fmt.Printf("Current speed: %v\n", c.speedUp())
	fmt.Printf("Current speed: %v\n", c.speedDown())
}

func main() {
	v := &vehicle{0, 2}
	b := &bycycle{0, 1}

	speedUpAndDown(v)
	speedUpAndDown(b)

	checkType(i2) //Type is nil
	i2 = 10       //i2 に 10 を代入
	checkType(i2) //Type is int, value is 10
	i2 = "hello"  //i2 に "hello" を代入
	checkType(i2) //Type is string, value is hello
}
