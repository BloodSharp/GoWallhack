#ifndef GOLANG_OPENGL_HACK_H
#define GOLANG_OPENGL_HACK_H

typedef unsigned int GLenum;
typedef unsigned char GLboolean;
typedef unsigned int GLbitfield;
typedef signed char GLbyte;
typedef short GLshort;
typedef int GLint;
typedef int GLsizei;
typedef unsigned char GLubyte;
typedef unsigned short GLushort;
typedef unsigned int GLuint;
typedef float GLfloat;
typedef float GLclampf;
typedef double GLdouble;
typedef double GLclampd;
typedef void GLvoid;

#define GL_TRIANGLE_STRIP 0x0005
#define GL_DEPTH_TEST 0x0B71

#define APIENTRY __stdcall
#define WINGDIAPI __declspec(dllimport)

WINGDIAPI void APIENTRY glBegin(GLenum mode);
WINGDIAPI void APIENTRY glDisable(GLenum cap);

#endif//GOLANG_OPENGL_HACK_H