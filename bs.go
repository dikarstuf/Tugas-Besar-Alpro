package main
import "fmt"
const NMAX = 1001

type tabMurid [NMAX]string

func main() {
	var n int
	var daftar tabMurid
	var target string

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&daftar[i])
	}

	fmt.Scan(&target)

	idx := binarySearch(daftar, n, target)

	if idx != -1 {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", idx+1)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}
}

func binarySearch(arr tabMurid, n int, x string) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}