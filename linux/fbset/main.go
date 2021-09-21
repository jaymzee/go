package main

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "fb.h"
import "C"
import (
	"fmt"
)

func main() {
	var fbinfo C.struct_fb_var_screeninfo
	dev := C.CString("/dev/fb0")
	C.query_framebuffer(dev, &fbinfo)

	fmt.Printf("%d %d %d %d %d\n", fbinfo.xres, fbinfo.yres,
		fbinfo.xres_virtual, fbinfo.yres_virtual,
		fbinfo.bits_per_pixel)
}
