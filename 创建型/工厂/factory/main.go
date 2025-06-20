package main

import (
	"factory/aa"
	"factory/ff"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	a := aa.Abs{Info: "aaa"}
	f := ff.Fac{Af: a}
	fmt.Println(f.Af.Info)
}
