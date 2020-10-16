package main

import (
	"fmt"
	"github.com/MufTho/Digital-Introduction-Golang/Basic_Package/Package_1"
	"github.com/MufTho/Digital-Introduction-Golang/Basic_Package/Package_2"
)

func main()  {
	Package_1.PrintPackage1()
	Package_1.CustomPrint("Mufid")

	fmt.Println(Package_1.Add(10, 20))
	result, flag := Package_1.AddDoubleReturn(20, 20)
	fmt.Println(fmt.Sprintf("hasil penjumlahan adalah %d, %t",result, flag))

	Package_2.ImportPrintPackage1()
}