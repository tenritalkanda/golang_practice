package main

import "fmt"

type Siswa struct {
	Nama, Kelas string
	Umur        int
}

func main() {
	//php => $ab = new Class()
	siswa := Siswa{}
	siswa.Nama = "Tenri"
	siswa.Kelas = "3A"
	siswa.Umur = 30

	fmt.Println(siswa)

}
