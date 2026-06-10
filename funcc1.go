package main
import "fmt"
func isPrima(bilangan int) bool {
	if bilangan < 2 {
		return false
	}
	for i := 2; i*i <= bilangan; i++ {
		if bilangan%i == 0 {
			return false
		}
	}
	return true
}

func sumAllDigit(bilangan int) int {
	total := 0
	for bilangan > 0 {
		total += bilangan % 10
		bilangan /= 10
	}
	return total
}

func sumUntilBeforeLimit(bil1, bil2, limit int) int {
	for bil1+bil2 < limit {
		bil1 += bil2
	}
	return bil1
}

func main() {
	var n, baru, penambah int
	fmt.Scan(&n)

	if n < 100 {
		penambah = n % 10
		if penambah == 0 {
			penambah = n / 10
		}
		baru = sumUntilBeforeLimit(n, penambah, 1000)
	} else if n < 1000 {
		baru = sumUntilBeforeLimit(n, sumAllDigit(n), 10000)
	} else {
		baru = n
	}

	status := "BUKAN PRIMA"
	if isPrima(baru) {
		status = "PRIMA"
	}

	fmt.Printf("%d --> %d %s\n", n, baru, status)
}