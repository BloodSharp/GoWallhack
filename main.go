package main

/*
#cgo LDFLAGS: -lkernel32
#include "winapi_definitions.h"
*/
import "C"
import (
	"unsafe"
)

var g_hThisDll C.DWORD = 0

func init() {
	szThisDll := C.CString("GoWallhack.dll")
	g_hThisDll = C.LoadLibraryA(szThisDll)
	C.free(unsafe.Pointer(szThisDll))
	if g_hThisDll != 0 {
		HookEverything()
	}
}

func main() {

}
