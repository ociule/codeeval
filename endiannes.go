package main

import "fmt"
import "unsafe"

func isLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}

func main() {
	if isLittleEndian() {
		fmt.Println("LittleEndian")
	} else {
		fmt.Println("BigEndian")
	}
}
