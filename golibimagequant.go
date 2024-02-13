package golibimagequant

// #cgo LDFLAGS: -L./libimagequant/target/release/ -limagequant_sys -lm
// #include "./libimagequant/imagequant-sys/libimagequant.h"
import "C"
import (
	"image/color"
	_ "image/png"
	"unsafe"
)

type cAttr *C.struct_liq_attr
type cImage *C.struct_liq_image
type cResult *C.struct_liq_result
type cPalette *C.struct_liq_palette
type cColor *C.struct_liq_color

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func NewColor(ccolor cColor) Color {
	return Color{
		r: uint8(ccolor.r),
		g: uint8(ccolor.g),
		b: uint8(ccolor.b),
		a: uint8(ccolor.a),
	}
}

func NewRGBA(ccolor cColor) color.RGBA {
	return color.RGBA{
		R: uint8(ccolor.r),
		G: uint8(ccolor.g),
		B: uint8(ccolor.b),
		A: uint8(ccolor.a),
	}
}

func Version() uint {
	return uint(C.liq_version())
}

func CreateAttr() cAttr {
	return C.liq_attr_create()
}

func DestroyAttr(attr cAttr) {
	C.liq_attr_destroy(attr)
}

func SetMaxColors(attr cAttr, colors int) {
	C.liq_set_max_colors(attr, C.int(colors))
}

func GetMaxColors(attr cAttr) int {
	return int(C.liq_get_max_colors(attr))
}

func SetQuality(attr cAttr, min int, max int) {
	C.liq_set_quality(attr, C.int(min), C.int(max))
}

func GetMinQuantity(attr cAttr) int {
	return int(C.liq_get_min_quality(attr))
}

func GetMaxQuantity(attr cAttr) int {
	return int(C.liq_get_max_quality(attr))
}

func CreateImageRGBA(attr cAttr, pixels *uint8, width int, height int, gamma float64) cImage {
	return C.liq_image_create_rgba(attr, unsafe.Pointer(pixels), C.int(width), C.int(height), C.double(gamma))
}

func DestroyImage(image cImage) {
	C.liq_image_destroy(image)
}

func AddFixedColor(image cImage, color Color) {
	C.liq_image_add_fixed_color(image, C.struct_liq_color{
		r: C.uchar(color.r),
		g: C.uchar(color.g),
		b: C.uchar(color.b),
		a: C.uchar(color.a),
	})
}

func QuantizeImage(attr cAttr, image cImage) cResult {
	return C.liq_quantize_image(attr, image)
}

func DestroyResult(result cResult) {
	C.liq_result_destroy(result)
}

func GetPalette(result cResult) cPalette {
	return C.liq_get_palette(result)
}

func WriteRemappedImage(result cResult, image cImage, buffer *uint8, bufferSize uint64) {
	C.liq_write_remapped_image(result, image, unsafe.Pointer(buffer), C.ulong(bufferSize))
}
