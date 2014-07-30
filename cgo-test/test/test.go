package test

// #include "test.h"
import "C"

import "fmt"

//export Println
func Println(num C.int) {
	fmt.Printf("%d\n", num)
}

func RunThings() {
	C.RunStuff()
}
