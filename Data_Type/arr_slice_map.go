package main

import "fmt"

func main()  {
	//array()
	//slice()
	mapExample()
}

func array()  {
	primes := [6]int{2,3,5,7,9,11}
	fmt.Println(primes)
}

func slice()  {
	names := []string{
		"nama1",
		"nama2",
		"nama3",
	}
	//fmt.Println(names)

	a := names[0:2]
	//fmt.Println(a)

	a[0] = "__"
	fmt.Println(names)
}

type Geo struct {
	Lat float64
	Long float64
}

func mapExample()  {
	mapGeo := make(map[string]interface{})

	mapGeo["Yogyakarta"] = Geo{
		Lat: 6.09123,
		Long: 102.23231,
	}

	mapGeo["Malang"] = Geo{
		Lat: 7.12309,
		Long: 109.0999,
	}

	for i,v := range mapGeo{
		fmt.Println(fmt.Sprintf("Koordinat %s : %v", i,v))
	}
}
