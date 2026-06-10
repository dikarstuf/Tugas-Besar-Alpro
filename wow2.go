package main

import "fmt"

func main() {
	var a, b, c int
	set, total := 0, 0
	for {
		fmt.Scan(&a, &b, &c)
		set++
		jumlah := a + b + c
		total += jumlah

		if a >= 90 || b >= 90 || c >= 90 || jumlah >= 210 {
			break
		}
	}
	rata := float64(total) / float64(set*3)
	fmt.Printf("total set : %d (rata-rata : %.2f WPM)\n", set, rata)
}
