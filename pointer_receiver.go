package main

import (
	"fmt"
)

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {
	siswa := Student{1, "Tenri", 3.15}

	fmt.Println(siswa.Name)

	graduate1(siswa)

	fmt.Println("setelah graduate 1 ", siswa.Name)

	graduate2(&siswa)

	fmt.Println("setelah graduate 2 ", siswa.Name)

	siswa.graduate3()

	fmt.Println("setelah graduate 3 ", siswa.Name)

}

func (student *Student) graduate3() {
	student.Name = student.Name + " M.Si"
}

func graduate1(student Student) {
	student.Name = student.Name + " S.Komp"
}

func graduate2(student *Student) {
	student.Name = student.Name + " S.Komp"
}
