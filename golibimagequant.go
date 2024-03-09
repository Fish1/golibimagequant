package golibimagequant

// #cgo LDFLAGS: -L./libimagequant/target/release/ -limagequant_sys -lm
// #include "stdlib.h"
// #include "./libimagequant/imagequant-sys/libimagequant.h"
import "C"
import (
	"image/color"
	_ "image/png"
	"unsafe"
)

type cLiqAttr *C.struct_liq_attr
type cLiqImage *C.struct_liq_image
type cLiqResult *C.struct_liq_result
type cLiqPalette *C.struct_liq_palette
type cLiqColor *C.struct_liq_color
type cLiqError int

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func NewColor(ccolor cLiqColor) Color {
	return Color{
		r: uint8(ccolor.r),
		g: uint8(ccolor.g),
		b: uint8(ccolor.b),
		a: uint8(ccolor.a),
	}
}

func NewRGBA(ccolor cLiqColor) color.RGBA {
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

// -------
// attr creation
// -------

func CreateAttr() cLiqAttr {
	return C.liq_attr_create()
}

func CopyAttr(attr cLiqAttr) cLiqAttr {
	return C.liq_attr_copy(attr)
}

func DestroyAttr(attr cLiqAttr) {
	C.liq_attr_destroy(attr)
}

// -------
// quality controls
// -------

func SetMaxColors(attr cLiqAttr, colors int) cLiqError {
	return cLiqError(C.liq_set_max_colors(attr, C.int(colors)))
}

func GetMaxColors(attr cLiqAttr) int {
	return int(C.liq_get_max_colors(attr))
}

func SetSpeed(attr cLiqAttr, speed int) cLiqError {
	return cLiqError(C.liq_set_speed(attr, C.int(speed)))
}

func GetSpeed(attr cLiqAttr) int {
	return int(C.liq_get_speed(attr))
}

func SetMinPosterization(attr cLiqAttr, bits int) cLiqError {
	return cLiqError(C.liq_set_min_posterization(attr, C.int(bits)))
}

func GetMinPosterization(attr cLiqAttr) int {
	return int(C.liq_get_min_posterization(attr))
}

func SetQuality(attr cLiqAttr, min int, max int) cLiqError {
	return cLiqError(C.liq_set_quality(attr, C.int(min), C.int(max)))
}

func GetMinQuantity(attr cLiqAttr) int {
	return int(C.liq_get_min_quality(attr))
}

func GetMaxQuantity(attr cLiqAttr) int {
	return int(C.liq_get_max_quality(attr))
}

// -------
// image creation
// -------

func CreateImageRGBA(attr cLiqAttr, pixels *uint8, width int, height int, gamma float64) cLiqImage {
	return C.liq_image_create_rgba(attr, unsafe.Pointer(pixels), C.int(width), C.int(height), C.double(gamma))
}

func DestroyImage(image cLiqImage) {
	C.liq_image_destroy(image)
}

// -------
// image controls
// -------

func AddFixedColor(image cLiqImage, color Color) {
	C.liq_image_add_fixed_color(image, C.struct_liq_color{
		r: C.uchar(color.r),
		g: C.uchar(color.g),
		b: C.uchar(color.b),
		a: C.uchar(color.a),
	})
}

// -------
// quantization
// -------

func QuantizeImage(attr cLiqAttr, image cLiqImage, result *cLiqResult) cLiqError {
	return cLiqError(C.liq_image_quantize(image, attr, (**C.struct_liq_result)(result)))
}

func DestroyResult(result cLiqResult) {
	C.liq_result_destroy(result)
}

// -------
// quantization results
// -------

func GetPalette(result cLiqResult) cLiqPalette {
	return C.liq_get_palette(result)
}

func WriteRemappedImage(result cLiqResult, image cLiqImage, buffer *uint8, bufferSize uint64) cLiqError {
	return cLiqError(C.liq_write_remapped_image(result, image, unsafe.Pointer(buffer), C.ulong(bufferSize)))
}
