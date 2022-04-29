package main

import "fmt"

var (
	i2 = 2
	j2 = 5
	s2 = "riach && full"
	k2 = true
)

func main() {
	fmt.Println("variable")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(i2, j2, s2, k2)
}

func variableZeroValue() {
	var i int
	var s string
	fmt.Printf("%d %q", i, s)
	fmt.Println()
}

func variableInitValue() {
	var i, j int = 1, 2
	var s string = "pizza"
	fmt.Println(i, j, s)
}

func variableTypeDeduction() {
	var i, j, s = 1, 2, "pizza"
	fmt.Println(i, j, s)
}

func variableShorter() { // 短变量声明变量
	i, j, s := 1, 2, "pizza"
	j = 3 // 前面有冒号后面不能再有冒号
	fmt.Println(i, j, s)
}
