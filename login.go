package main

import "fmt"

func main() {
	var x, y string
	fmt.Scan(&x, &y)

	if x == "admin" && y == "12345" {
		fmt.Println("Login Berhasil")
	} else if x == "admin" && y != "12345" {
		fmt.Println("Password Salah")
	} else if x != "admin" {
		fmt.Println("Username Tidak Ditemukan")

		if y != "12345" {
			fmt.Println("Password Salah")
		}
	}
}
