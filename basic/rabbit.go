package basic

import "fmt"

type Rabbit struct{}

func (r Rabbit) Sound() {
	fmt.Println("hehe hehehe")
}
func (r Rabbit) NumOfLegs() {
	fmt.Println(4)
}
