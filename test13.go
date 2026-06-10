package main
import "fmt"

type arrInt [100]int

func getTens(num int) int {
	return (num / 10) % 10
}
func shouldSwap(a, b int) bool {
	tA := getTens(a)
	tB := getTens(b)
	if tA > tB {
		return true
	}
	if tA == tB && a > b {
		return true
	}
	return false
}
func insertionSort(arr *arrInt, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && shouldSwap(arr[j], key) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
func main() {
	var n int
	var data arrInt

	if fmt.Scan(&n); n <= 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Println("Data sebelum sorting:")
	for i := 0; i < n; i++ {
		fmt.Println(data[i])
	}

	fmt.Println()
	insertionSort(&data, n)
	fmt.Println("Data setelah sorting:")
	for i := 0; i < n; i++ {
		fmt.Println(data[i])
	}
}