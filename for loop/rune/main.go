package main

import "fmt"

func main() {

	for i := 50; i <= 140; i++ {
		//fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
	varRune := 'a'
	fmt.Println(varRune)
	fmt.Printf("%T \n", varRune)
	fmt.Printf("%T \n", rune(varRune))
}
