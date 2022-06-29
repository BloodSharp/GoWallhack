package main

/*
#cgo LDFLAGS: -lopengl32
#include "opengl_definitions.h"
*/
import "C"

//export HOOK_glBegin
func HOOK_glBegin(mode C.GLenum) {
	if mode == C.GL_TRIANGLE_STRIP {
		C.glDisable(C.GL_DEPTH_TEST)
	}
	C.glBegin(mode)
}
