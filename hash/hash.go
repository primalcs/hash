package hash

/*
#include "hash.h"
*/
import "C"

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func reverseHexEndianRepresentation(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-2; i < j; i, j = i+2, j-2 {
		rns[i], rns[j] = rns[j], rns[i]
		rns[i+1], rns[j+1] = rns[j+1], rns[i+1]
	}
	return string(rns)
}

func Hash(input1, input2 string) string {
	input1_dec, _ := hex.DecodeString(reverseHexEndianRepresentation(input1))
	input2_dec, _ := hex.DecodeString(reverseHexEndianRepresentation(input2))
	in1 := C.CBytes(input1_dec)
	in2 := C.CBytes(input2_dec)
	var o [1024]byte
	out := C.CBytes(o[:])

	res := C.CHash(
		(*C.char)(unsafe.Pointer(in1)),
		(*C.char)(unsafe.Pointer(in2)),
		(*C.char)(unsafe.Pointer(out)))

	if res != 0 {
		fmt.Printf("Pedersen hash encountered an error: %s\n", C.GoBytes(unsafe.Pointer(out), 1024))
		//C.free(unsafe.Pointer(in1))
		//C.free(unsafe.Pointer(in2))
		//C.free(unsafe.Pointer(out))
		return ""
	}

	hash_result := "0x" + reverseHexEndianRepresentation(
		hex.EncodeToString(C.GoBytes(unsafe.Pointer(out), 32)))

	//C.free(unsafe.Pointer(in1))
	//C.free(unsafe.Pointer(in2))
	//C.free(unsafe.Pointer(out))

	return hash_result
}
