package basic

import (
	"fmt"
	"os"
)

func GetFullnme(firstname string, lastname string) string {
	return fmt.Sprintf("%s %s", firstname, lastname)
}

func TryDefer() {
	defer fmt.Printf("sampai jumpa")
	fmt.Println("halo selamat datang")
}

func TryExit() {
	defer fmt.Printf("sampai jumpa")
	fmt.Println("halo selamat datang")
	os.Exit(1)
}

func Pointer() {
	var date int = 20
	var pdate *int = &date

	date = 21

	fmt.Println(date, *pdate)
	fmt.Println(&date, pdate)

}

type IAnimal interface{
	Sound()
	NumOfLegs()
}

func Animal(animal IAnimal){
	animal.Sound()
	animal.NumOfLegs()
}
