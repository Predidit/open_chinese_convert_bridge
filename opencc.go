package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"log"
	"unsafe"

	"github.com/longbridgeapp/opencc"
)

//export TranS2TW
func TranS2TW(_in *C.char) *C.char {
	in := C.GoString(_in)
	s2t, err := opencc.New("s2tw")
	if err != nil {
		log.Fatal(err)
	}
	out, err := s2t.Convert(in)
	if err != nil {
		log.Fatal(err)
	}
	return C.CString(out)
}

//export TranTW2S
func TranTW2S(_in *C.char) *C.char {
	in := C.GoString(_in)
	t2s, err := opencc.New("tw2s")
	if err != nil {
		log.Fatal(err)
	}
	out, err := t2s.Convert(in)
	if err != nil {
		log.Fatal(err)
	}
	return C.CString(out)
}

//export FreeString
func FreeString(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func main() {
	// Test
}
