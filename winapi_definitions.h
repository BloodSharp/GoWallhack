#ifndef GOLANG_WINAPI_HACK_H
#define GOLANG_WINAPI_HACK_H

#include <stdlib.h>

typedef void* LPVOID;
typedef void* HANDLE;
typedef unsigned long DWORD;
typedef const char* LPCSTR;
typedef int BOOL;
typedef short SHORT;

#define VK_DELETE 0x2E
#define PAGE_EXECUTE_READWRITE 0x40
#define MEM_COMMIT 0x00001000
#define MEM_RESERVE 0x00002000
#define MEM_DECOMMIT 0x00004000
#define MEM_RELEASE 0x00008000

#define WINAPI __stdcall
#define DLLIMPORT  __declspec(dllimport)

DLLIMPORT HANDLE WINAPI CreateThread(DWORD lpThreadAttributes, DWORD dwStackSize, DWORD lpStartAddress, DWORD lpParameter, DWORD dwCreationFlags, DWORD lpThreadId);
DLLIMPORT BOOL WINAPI DisableThreadLibraryCalls(DWORD hLibModule);
DLLIMPORT SHORT WINAPI GetAsyncKeyState(DWORD vKey);
DLLIMPORT DWORD WINAPI GetModuleHandleA(LPCSTR lpModuleName);
DLLIMPORT DWORD WINAPI GetProcAddress(DWORD hModule, LPCSTR szProcName);
DLLIMPORT void WINAPI FreeLibraryAndExitThread(DWORD hLibModule, DWORD dwExitCode);
DLLIMPORT DWORD WINAPI LoadLibraryA(LPCSTR szModuleName);
DLLIMPORT void WINAPI Sleep(DWORD dwMilliseconds);
DLLIMPORT LPVOID WINAPI VirtualAlloc(DWORD lpAddress, DWORD dwSize, DWORD flAllocationType, DWORD flProtect);
DLLIMPORT BOOL WINAPI VirtualFree(LPVOID lpAddress, DWORD dwSize, DWORD dwFreeType);

#endif//GOLANG_WINAPI_HACK_H