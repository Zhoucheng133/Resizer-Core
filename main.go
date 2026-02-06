package main

import "C"
import (
	"fmt"
	"resizer_core/utils"
)

//export Resize
func Resize(
	path *C.char,
	width C.int,
	height C.int,
	output *C.char,
	stretch C.int,
) *C.char {
	return C.CString(utils.ResizeHandler(C.GoString(path), int(width), int(height), C.GoString(output), stretch != 0))
}

//export GetSize
func GetSize(path *C.char) *C.char {
	return C.CString(utils.GetSizeHandler(C.GoString(path)))
}

//export ReadDir
func ReadDir(path *C.char) *C.char {
	return C.CString(utils.ReadDirHandler(C.GoString(path)))
}

func main() {
	fmt.Println(utils.ReadDirHandler("/Users/zhoucheng/Desktop/Develop/netPlayer-Mobile/ios/Runner/Assets.xcassets/AppIcon.appiconset"))
}
