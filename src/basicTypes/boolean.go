package basicTypes

import (
	"fmt"
	"unsafe"
)

func BooleanTypes() {
	var a bool
	fmt.Println("a:::", a)
	var b bool = true
	fmt.Println("b:::", b)
	b = false
	fmt.Println("b:::", b)
	c := true
	fmt.Println("c:::", c)
	fmt.Println("c:::size", unsafe.Sizeof(c))

}
