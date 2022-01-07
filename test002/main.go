package main

import (
	"fmt"
	"test003/testlib"
)

func main () {
	fmt.Println("test!!!")
	song := testlib.GetMusic("Alicia Keys")
	println(song)
	testlib.GetKeys()

	fmt.Println("TEST!!!!")
}
