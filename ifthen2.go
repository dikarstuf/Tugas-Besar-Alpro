package main
import "fmt"
func main() {
    var jumlahMahasiswa,kapasitasBus,jumlahBus,sisaMahasiswa int
	kapasitasBus = 45
	fmt.Scan(&jumlahMahasiswa)
	jumlahBus = jumlahMahasiswa / kapasitasBus
	sisaMahasiswa = jumlahMahasiswa % kapasitasBus

    if sisaMahasiswa > 0 {
        jumlahBus = jumlahBus + 1
    }

    fmt.Printf("Diperlukan %d bus untuk tamasya ke Lembang\n", jumlahBus)
}