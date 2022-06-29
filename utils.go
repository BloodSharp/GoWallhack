package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"unsafe"
)

func memcpy(dest, src []byte, iSize int) {
	for iIndex := 0; iIndex < iSize; iIndex++ {
		dest[iIndex] = src[iIndex]
	}
}

func DataCompare(pData uintptr, bMask []byte, szMask string) bool {
	for i := 0; i < len(szMask); i++ {
		pData++
		if szMask[i] == 'x' && ((*[1]byte)(unsafe.Pointer(pData))[:])[0] != bMask[i] {
			return false
		}
	}
	return true
}

func FindPattern(dwAddress uintptr, dwLen int, bMask []byte, szMask string) uintptr {
	for i := 0; i < dwLen; i++ {
		if DataCompare(uintptr(dwAddress+uintptr(i)), bMask, szMask) {
			return (dwAddress + uintptr(i) + 1) // Fix me ==> + 1
		}
	}
	return 0
}

const (
	kEngineBuildUndefined = 0
	kEngineBuildv3266     = 1
	kEngineBuildv4554     = 2
	kEngineBuildvSteam    = 3
)

func GetEngineVersion() int {
	_, err := os.ReadFile(".\\SDL2.dll")
	if err == nil {
		return kEngineBuildvSteam
	}
	fHwDll, err := os.ReadFile(".\\hw.dll")
	if err != nil {
		fmt.Println("[B#] Couldn't read hw.dll file...")
		return kEngineBuildUndefined
	}
	byHwDllHash := md5.Sum(fHwDll)
	if bytes.Compare(byHwDllHash[:], []byte{155, 213, 11, 252, 98, 200, 130, 73, 18, 147, 130, 33, 50, 219, 71, 129}) == 0 {
		return kEngineBuildv3266
	}
	if bytes.Compare(byHwDllHash[:], []byte{154, 56, 153, 223, 101, 49, 91, 73, 75, 86, 46, 68, 162, 54, 88, 212}) == 0 {
		return kEngineBuildv4554
	}
	fmt.Printf("md5.Sum(fHwDll): %v\n", byHwDllHash)
	return kEngineBuildUndefined
}
