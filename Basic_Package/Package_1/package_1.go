package Package_1

import "fmt"

func PrintPackage1()  {
	fmt.Println("Ini Print dari package1")
}

func CustomPrint(str string)  {
	fmt.Println(fmt.Sprintf("Halo nama saya %s", str))
}

func Add(x, y int) int {
	return x + y
}

func AddDoubleReturn(x, y int)(int, bool)  {
	return x +y, true
}

