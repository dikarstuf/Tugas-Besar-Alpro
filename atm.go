package main

import (
	"fmt"
)

func main() {
	var PIN, attempt int
	var nominal int

	fmt.Print("PIN:")
	fmt.Scan(&PIN)
	attempt = 1

	for PIN != 241231 && attempt < 3 {
		fmt.Println("PIN incorrect")
		fmt.Print("PIN:")
		fmt.Scan(&PIN)
		attempt = attempt + 1
	}

	if PIN == 241231 {
		fmt.Print("Nominal:")
		fmt.Scan(&nominal)
		fmt.Printf("%d rupiahs is withdrawn. Thank you.\n", nominal)
	} else {
		fmt.Println("Account is blocked, contact nearest branch!")
	}
}