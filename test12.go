package main
import "fmt"
const NMAX = 9999
type arrString [NMAX]string
func bacaData(A *arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}
func cetakData(A arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
		if i < n-1 {
			fmt.Print(", ")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func selectionSortAscend(A *arrString, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}
		temp := A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}
func selectionSortDescend(A *arrString, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if A[j] > A[idxMax] {
				idxMax = j
			}
		}
		temp := A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 || n > NMAX {
		return
	}

	var data arrString
	bacaData(&data, n)

	selectionSortAscend(&data, n)
	fmt.Println("Data setelah diurutkan secara Ascending:")
	cetakData(data, n)

	selectionSortDescend(&data, n)
	fmt.Println("Data setelah diurutkan secara Descending:")
	cetakData(data, n)
}