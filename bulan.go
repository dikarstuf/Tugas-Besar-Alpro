package main

import "fmt"

func main() {
	var bulan int
	fmt.Scan(&bulan)

	switch bulan {
	case 1:
		fmt.Println("Januari")
	case 2:
		fmt.Println("Februari")
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7:
		fmt.Println("Juli")
	case 8:
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("Oktober")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("Desember")
	default:
		fmt.Println("Input is Invalid")
	}
}
