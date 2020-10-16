package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main()  {
	v := Vertex{3, 5}
	p := &v
	p.X = 10
	fmt.Println(v)

	a := &Vertex{10, 30}
	q := a
	q.X = 40
	fmt.Println(*a)

	// array
	var arrStruct []Vertex
	for i := 0; i < 5; i++{
		arrStruct = append(arrStruct, Vertex{i, i+1})
	}
	fmt.Println(arrStruct)
}
