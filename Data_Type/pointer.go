package main

import "fmt"

func main()  {
	i := 10
	p := &i
	fmt.Println(*p)

	*p = 20
	fmt.Println(i)
}