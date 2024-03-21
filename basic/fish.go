package basic

import "fmt"

type Fish struct{}

func (f Fish) Sound() {
	fmt.Println(".........")
}
func (f Fish) NumOfLegs() {
	fmt.Println(0)
}
