package main

// #include <stdlib.h>
// #include <android/log.h>
// void go_android_info(const char* tag, const char* str) {
//   __android_log_print(ANDROID_LOG_INFO, tag, "%s", str);
// }
import "C"
import "unsafe"

func log(str string) {
	ctag := C.CString("govrhelloworld")
	cstr := C.CString(str)
	C.go_android_info(ctag, cstr)
	C.free(unsafe.Pointer(ctag))
	C.free(unsafe.Pointer(cstr))
}
