package main

import "fmt"

type Biodata struct {
	Nama string
	Alamat string
	Hp int
}

type Sisi struct {
	Panjang int
	Lebar int
}

func main()  {
	s := Sisi{3,4}
	s.ScaleUp(2)
	//cara lama
	//fmt.Println(Luas(s))

	//cara punya GO
	fmt.Println(s.Luas())
	//fmt.Println(s)

	bio := Biodata{
		"Mufid",
		"Malang",
		6281334455588,
	}

	bio.rubahAlamat("Yogyakarta")
	fmt.Println(bio)
}
//cara lama
func Luas(s Sisi) int {
	return s.Panjang * s.Lebar
}
//cara baru (punya e go)
func (s Sisi) Luas()int  {
	return s.Panjang * s.Lebar
}

func (s *Sisi) ScaleUp(i int) {
	s.Lebar = s.Lebar*i
	s.Panjang = s.Panjang*i
}

func (b *Biodata) rubahAlamat (newAlamat string)  {
	b.Alamat = newAlamat
}