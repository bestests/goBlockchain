package testlib

import (
	"fmt"
)

var pop map[string]string

func init () {
	pop = make(map[string]string)

	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallen"
	pop["John Legend"] = "All of Me"
}

func GetMusic (singer string) string {
	return pop[singer]
}

func GetKeys () {
	fmt.Println("------------")
	for key, _ := range pop {
		fmt.Println(key)
		fmt.Println("------------")
	}
}
