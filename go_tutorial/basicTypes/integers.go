package basicTypes

import (
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

/*
Go has fixed-size signed and unsigned integers:

int8, uint8 //The value range of int8 is – 128 ~ 127 , If you want to convert, you need to use bytevalue = 256 + int8value
byte is an alias for uint8
int16, uint16
int32, uint32 //on 32-bit processors
rune is an alias for int32 //stores Unicode characters or can store an integer value of at most 32-bits.
int64, uint64 //on 64-bit processors
*/
func IntegerTypes() {
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println("size of Int:::", unsafe.Sizeof(i1))
	fmt.Println("size of Int8:::", unsafe.Sizeof(i2))
	fmt.Println("size of Int16:::", unsafe.Sizeof(i3))
	fmt.Println("size of Int32:::", unsafe.Sizeof(i4))
	fmt.Println("size of Int64:::", unsafe.Sizeof(i5))

	fmt.Println("Convert int to string with strconv.Itoa", strconv.Itoa(i1))
	fmt.Println("Convert int to string with fmt.Sprintf", fmt.Sprintf("%d", i2))

	s := "40"
	its, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("error while converting string to int", err)
	}
	fmt.Println("Convert string to int with strconv.Atoi", its)
	var i int
	_, err = fmt.Sscanf(s, "%d", &i)
	fmt.Println("Convert string to int with fmt.Sscanf", i)
}
