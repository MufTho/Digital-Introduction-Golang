package main

import (
	"fmt"
	"runtime"
)

func main()  {
	//arr := [] string{"satu", "dua", "tiga", "empat"}
	//perulangan(arr)

	//fmt.Println(fmt.Sprintf("Flag : %t", ifstatement(10, 15, 20)))
	switchstatement()
}

func perulangan(arr []string)  {
	//	For Basic
	for i := 0; i< len(arr); i++ {
		fmt.Println(arr[i])
	}

	//	For dengan prinsip (foreach)
	for i, v := range arr{
		fmt.Println(fmt.Sprintf("Index ke %d : %s", i,v))
	}

	//	For dengan prinsip (foreach) menghilankan index nya, mengambil valuenya
	for _, v := range arr{
		fmt.Println(fmt.Sprintf("valuenya adalah  %s", v))
	}

	// Simplified for (while)
	i := 0
	for {
		if i == len(arr){
			break
		}
		fmt.Println(fmt.Sprintf("valuenya adalah %s", arr[i]))
		i++
	}
}

func ifstatement(x, y, max int) bool {
	// type 1
	//if x < y{
	//	return true
	//}else {
	//	return false
	//}

	// Type 2
	if result := x * y; result < max {
		return false
	}else {return true}
}

func switchstatement()  {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s .\n", os)
		
	}
}

