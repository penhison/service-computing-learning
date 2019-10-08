package main

// #include "selpg.c"
import "C"

import "unsafe"

import (
	"fmt"
	"os"
)

func main() {
	argv := make([](*C.char), 0)
	for _, s := range os.Args {//
		cs := C.CString(s)
		defer C.free(unsafe.Pointer(cs))
		argv = append(argv, (*C.char)(unsafe.Pointer(cs)))
	}

	C._main((C.int)(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))
	fmt.Println("done")
}