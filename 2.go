package main
import "fmt"
func main() {
	var s, pu, su, sq, jj, i  int
	var kaloriPu, kaloriSu, kaloriSq, kaloriJj, totalkalori, kalori float64
	
	fmt.Scan(&s)
	fmt.Scan(&pu, &su, &sq, &jj)

	for i = 1; i <= pu; i++ {
		kaloriPu += float64(i) * 0.5
	}
  
	for i = 1; i <= su; i++ {
		kaloriSu += float64(i) * 0.3
	}

	for i = 1; i <= sq; i++ {
		kaloriSq += float64(i) * 0.2
	}

	for i = 1; i <= jj; i++ {
		kaloriJj += float64(i) * 0.6
	}

	kalori = kaloriPu + kaloriSu + kaloriSq + kaloriJj
	totalkalori = float64(s) * kalori

	fmt.Printf("Total kalori terbakar hari ini sebanyak %g kalori\n", totalkalori)
}