package main

import "fmt"

func main() {
	var var1 string = "Variável 1"
	fmt.Println(var1)

	var2 := "Variável 2"
	fmt.Println(var2)

	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	fmt.Println(var3)
	fmt.Println(var4)

	var5, var6 := "Var5", "var6"
	fmt.Println(var5)
	fmt.Println(var6)
}
