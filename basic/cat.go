package basic

import "fmt"

type Cat struct{}

func (c Cat) Sound() {
	fmt.Println("meong meong")
}
func (c Cat) NumOfLegs() {
	fmt.Println(4)
}