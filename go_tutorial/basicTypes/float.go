package basicTypes

import (
	"fmt"
	"strconv"
)

func FloatTypes() {
	var f float32 = 1.3
	fmt.Println("the value of f is:", f)
	s := "1.56"
	fmt.Println("the value of s:", s)

	//convert float to string
	fmt.Println("convert float to string:", strconv.FormatFloat(float64(f), 'E', -1, 32))

	//convert string to float
	f64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("convert string to float:", f64)
}
