package main
import "fmt"
const NMAX = 100
type Peserta struct {
	id     string
	nama   string
	nilai  int
	durasi int
}
type TabPeserta [NMAX]Peserta
func insertionSort(A *TabPeserta, N int) {
	var i, j int
	var temp Peserta

	for i = 1; i < N; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 &&
			(A[j].nilai < temp.nilai ||
				(A[j].nilai == temp.nilai &&
					A[j].durasi > temp.durasi)) {

			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}
func main() {
	var A TabPeserta
	var N, i int
	var totalNilai int
	var rata float64
	var jumlah int

	fmt.Scan(&N)

	for i = 0; i < N; i++ {
		fmt.Scan(&A[i].id, &A[i].nama, &A[i].nilai, &A[i].durasi)
		totalNilai += A[i].nilai
	}

	insertionSort(&A, N)

	fmt.Println("Data setelah diurutkan:")
	for i = 0; i < N; i++ {
		fmt.Println(A[i].id, A[i].nama, A[i].nilai, A[i].durasi)
	}

	fmt.Println()
	fmt.Println("Peserta terbaik:")
	fmt.Println(A[0].id, A[0].nama, A[0].nilai, A[0].durasi)

	rata = float64(totalNilai) / float64(N)

	for i = 0; i < N; i++ {
		if float64(A[i].nilai) > rata {
			jumlah++
		}
	}

	fmt.Println()
	fmt.Println("Jumlah peserta di atas rata-rata:", jumlah)
}