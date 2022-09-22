package basicTypes

import (
	"fmt"
	"reflect"
)

func RuneTypes() {
	r1 := "~"
	fmt.Printf("the value of r1 :%s and type: %s\n", r1, reflect.TypeOf(r1))

	s1 := "falcon"
	r2 := []rune(s1)
	fmt.Printf("the value of r2 in Unicode :%U\n", r2)

	s2 := "🐧🐧🐧"
	r3 := []rune(s2)
	fmt.Printf("the value of r3 in Unicode :%U\n", r3)

	for _, i := range r3 {
		fmt.Printf("the value of r3 :%c and type: %s\n", i, reflect.TypeOf(i))
	}

}
