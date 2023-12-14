package basicTypes

import (
	"bytes"
	"fmt"
	"unsafe"
)

// go doesn't have the char type as it has byte data type which holds 1 byte

func ByteTypes() {
	myByte := byte('R')
	fmt.Printf("the value of Byte var: %c and size is: %d \n", myByte, unsafe.Sizeof(myByte))

	myBytes := []byte{'V', 'a', 'n', 's', 'h', 'i', 'k', 'a'} // byte array with {}
	b1 := []byte("Vanshika")
	b2 := []byte("Resu")
	twoDimenByte := [][]byte{b1, b2}
	b3 := bytes.Join(twoDimenByte, []byte(","))
	var b4 = [][]byte{[]byte("baby"), []byte("vanshika"), []byte("Resu")}
	for i := 0; i < len(myBytes); i++ {
		fmt.Printf("the value at index :%d is :%c\n", i, myBytes[i]) //printing bytes from string

	}

	fmt.Println("The value of b1 :", string(b1))
	fmt.Println("The value of b2 :", string(b2))
	fmt.Printf("The value of twoDimenByte :%s\n", twoDimenByte)
	fmt.Printf("The value of b3 :%s\n", b3)
	fmt.Printf("The value of b4 :%s\n", bytes.Join(b4, []byte("|")))
}
