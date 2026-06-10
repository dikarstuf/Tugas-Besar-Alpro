package main

import "fmt"

func main() {
	var char rune
	var hasil bool
	
	fmt.Scanf("%c", &char)
	
	hurufbesar := (char >= 'A' && char <= 'Z' )
	hurufkecil := (char >= 'a' && char <= 'z' )
	
	hasil = hurufbesar || hurufkecil
	
	fmt.Println(hasil)

}    