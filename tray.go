package trayicon

/*
#cgo linux pkg-config: gtk+-3.0
#cgo linux CFLAGS: -DLINUX
#cgo windows CFLAGS: -DWIN32
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa
#include <stdlib.h>

#if defined LINUX
#include "platform/linux.c"
#elif defined WIN32
#include "platform/windows.c"
#elif defined DARWIN
#include "platform/darwin.m"
#endif

*/
import "C"
