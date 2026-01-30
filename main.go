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
) *C.char {
	return C.CString(utils.ResizeHandler(C.GoString(path), int(width), int(height), C.GoString(output)))
}

//export GetSize
func GetSize(path *C.char) *C.char {
	return C.CString(utils.GetSizeHandler(C.GoString(path)))
}

func main() {
	fmt.Println(utils.GetSizeHandler("/Users/zhoucheng/Downloads/照片/DSC_2636.jpg"))
}
