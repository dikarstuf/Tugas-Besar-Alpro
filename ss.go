package main
import "fmt"
const NMAX = 999

type Rapot struct {
	matkul string
	nilai  int
}

type TabRapot [NMAX]Rapot

func main() {
	var T TabRapot
	var N, X int

	fmt.Scan(&N)

	if N > NMAX {
		N = NMAX
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&T[i].matkul, &T[i].nilai)
	}

	fmt.Scan(&X)

	var ketemu int = -1
	var k int = 0
	var foundAny bool = false

	for ketemu == -1 && k < N {
		if T[k].nilai == X {
			
			for j := k; j < N; j++ {
				if T[j].nilai == X {
					fmt.Printf("%s %d\n", T[j].matkul, T[j].nilai)
					foundAny = true
				}
			}
			ketemu = k
		}
		k = k + 1
	}

	if foundAny {
		fmt.Println("Data ditemukan!")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}