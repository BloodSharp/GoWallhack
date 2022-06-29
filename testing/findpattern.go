package main_test

import "unsafe"

func DataCompare(pData uintptr, bMask []byte, szMask string) bool {
	for i := 0; i < len(szMask); i++ {
		pData++
		if szMask[i] == 'x' && ((*[1]byte)(unsafe.Pointer(pData))[:])[0] != bMask[i] { //PointerContent(pData) != bMask[i] {
			return false
		}
	}
	return true
}

func FindPattern(dwAddress uintptr, dwLen int, bMask []byte, szMask string) uintptr {
	for i := 0; i < dwLen; i++ {
		if DataCompare(uintptr(dwAddress+uintptr(i)), bMask, szMask) {
			return (dwAddress + uintptr(i) + 1) //Fix me +1
		}
	}
	return 0
}
