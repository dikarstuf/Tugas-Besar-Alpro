package main
import "fmt"
const NMAX = 100
type Applicant struct {
	id           string
	name         string
	testGrade    float64
	testDuration int
}
func main() {
	var applicants [NMAX]Applicant
	var count int = 0
	for count < NMAX {
		var id string
		fmt.Scan(&id)
		if id == "END" {
			break
		}
		var name string
		var testGrade float64
		var testDuration int
		fmt.Scan(&name, &testGrade, &testDuration)
		applicants[count] = Applicant{
			id:           id,
			name:         name,
			testGrade:    testGrade,
			testDuration: testDuration,
		}
		count++
	}
	fmt.Println("Data awal:")
	for i := 0; i < count; i++ {
		app := applicants[i]
		fmt.Printf("%s %s %.1f %d\n", app.id, app.name, app.testGrade, app.testDuration)
	}
	fmt.Println()
	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			tukar := applicants[j].testGrade < applicants[j+1].testGrade
			if applicants[j].testGrade == applicants[j+1].testGrade {
				tukar = applicants[j].testDuration > applicants[j+1].testDuration
			}
			if tukar {
				temp := applicants[j]
				applicants[j] = applicants[j+1]
				applicants[j+1] = temp
			}
		}
	}
	fmt.Println("Data setelah sortir:")
	for i := 0; i < count; i++ {
		app := applicants[i]
		fmt.Printf("%s %s %.1f %d\n", app.id, app.name, app.testGrade, app.testDuration)
	}
}