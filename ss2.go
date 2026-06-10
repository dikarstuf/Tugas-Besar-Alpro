package main
import "fmt"
const NMAX = 100

type student struct {
	nim, name string
	grade     float64
}

type students [NMAX]student

func main() {
	var data students
	var inputNIM string
	var inputName string
	var inputGrade float64
	var n int = 0

	for {
		fmt.Scan(&inputNIM)
		if inputNIM == "STOP" || n >= NMAX {
			break
		}
		fmt.Scan(&inputName, &inputGrade)

		var isDuplicate bool = false
		var k int = 0
		for k < n {
			if data[k].nim == inputNIM {
				isDuplicate = true
				break
			}
			k = k + 1
		}

		if !isDuplicate {
			data[n].nim = inputNIM
			data[n].name = inputName
			data[n].grade = inputGrade
			n = n + 1
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %v\n", data[i].nim, data[i].name, data[i].grade)
	}
}