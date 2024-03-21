package main

import (
	"fmt"
	"project-go/basic"
)

func main() {

	//unexported
	fmt.Println(isOdd(10))

	//exported
	fmt.Println(basic.GetFullnme("dita", "rizki"))

	deklarasiData()
	perulangan()
	arraySlice()
	tryMap()
	basic.Pointer()

	//manifest type
	//var mypost basic.Post
	// mypost.Title = "ketika cinta bertasbih"
	// mypost.Content = "yaa isinya begitulah pokoknya"
	// mypost.Author = "dita"
	// fmt.Println(mypost)

	//interface type
	// mypost := basic.Post{
	// 	Title:   "judul",
	// 	Content: "ini isinya",
	// 	Author:  "dita",
	// }
	// fmt.Printf("%s\n%s\n%s", mypost.Title, mypost.Content, mypost.Author)

	myuser := basic.NewUser("dita", "rizki", 26)
	mypost := basic.NewPost("ini judul", "ini isinya", *myuser)
	fmt.Printf("%s\n%s\n%v", mypost.Title, mypost.Content, mypost.Author)
	mypost.EditContent("ini isi baru")
	fmt.Printf("%s\n%s\n%s\n", mypost.Title, mypost.Content, mypost.Author.Getfullname())

	cat := basic.Cat{}
	rabbit := basic.Rabbit{}
	fish := basic.Fish{}
	basic.Animal(cat)
	basic.Animal(rabbit)
	basic.Animal(fish)

	fmt.Printf("\n=====================================\n")
	basic.TryDefer()
	basic.TryExit()
}

func penjumlahan(num1 int, num2 int) int {
	return num1 + num2
}

func isEven(num int) bool {
	if num == 0 {
		return false
	}
	if num%2 == 0 {
		return true
	}
	return false
}

func deklarasiData() {
	//manifest type
	//var/const namaVar tipeData
	var firstName string = "dita"
	//const pi float32 = 3.14
	firstName = "dita Rizki"

	//type interface
	age := 23
	isMarried := false
	fmt.Println(firstName, age, isMarried)
	fmt.Printf("Nama : %s\nage : %d\nisMarried: %v\n", firstName, age, isMarried)
}
func perulangan() {
	//foreach
	//for idx, value bisa kosong diganti dengan _
	drinks := []string{"soda", "jus", "coffe", "tea"} //ini slice
	fmt.Println("Tersedia minuman : ")
	for index, value := range drinks {
		fmt.Printf("%d. %s\n", index+1, value)
	}

}
func arraySlice() {
	//array
	var buah [3]string

	buah[0] = "jeruk"
	buah[1] = "apel"
	buah[2] = "mangga"

	//var kendaraan = [3]string{"mobil", "motor", "bus"}

	//slice
	//perbedaan slice dan array dari cara definisi dan cara slice
	buahslice := [3]string{"apel", "mangga", "pisang"}
	huah := buahslice
	buahslice[2] = "tomat"
	fmt.Println(buahslice, huah)

	var hasil int = penjumlahan(2, 2)
	fmt.Println(hasil, isOdd(hasil))
}

func tryMap() {
	//map
	//asosiasi key,value
	//map tanpa assign nilai
	var siswa = make(map[string]string)
	siswa["nama"] = "dita"
	siswa["umur"] = "23"
	siswa["jenis_kelamin"] = "perempuan"

	fmt.Println(siswa)

	//map dengan langsung assign nilai
	person := map[string]string{
		"nama": "dita",
		"umur": "23",
	}
	fmt.Println(person)

}
