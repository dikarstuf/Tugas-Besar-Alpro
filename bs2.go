package main
import "fmt"
type Penduduk struct {
	Nama  string
	Tahun int
	Kota  string
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]Penduduk, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].Nama, &data[i].Tahun, &data[i].Kota)
	}

	var cariTahun int
	fmt.Scan(&cariTahun)

	foundIdx := -1
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid].Tahun == cariTahun {
			foundIdx = mid
			break
		} else if data[mid].Tahun < cariTahun {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundIdx != -1 {
		fmt.Println(data[foundIdx].Nama)
		fmt.Println(data[foundIdx].Tahun)
		fmt.Println(data[foundIdx].Kota)
		fmt.Printf("ditemukan di index ke-%d\n", foundIdx)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}