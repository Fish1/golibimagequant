package golibimagequant

// #cgo LDFLAGS: -L./libs/ -limagequant_sys -lm
// #include "stdlib.h"
// #include "./libs/libimagequant.h"
import "C"
import (
	"image/color"
	_ "image/png"
	"unsafe"
)

type LiqAttr *C.struct_liq_attr
type LiqImage *C.struct_liq_image
type LiqResult *C.struct_liq_result
type LiqPalette *C.struct_liq_palette
type LiqColor *C.struct_liq_color
type LiqError int

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func NewColor(ccolor LiqColor) Color {
	return Color{
		r: uint8(ccolor.r),
		g: uint8(ccolor.g),
		b: uint8(ccolor.b),
		a: uint8(ccolor.a),
	}
}

func NewRGBA(ccolor LiqColor) color.RGBA {
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

func CreateAttr() LiqAttr {
	return C.liq_attr_create()
}

func CopyAttr(attr LiqAttr) LiqAttr {
	return C.liq_attr_copy(attr)
}

func DestroyAttr(attr LiqAttr) {
	C.liq_attr_destroy(attr)
}

// -------
// quality controls
// -------

func SetMaxColors(attr LiqAttr, colors int) LiqError {
	return LiqError(C.liq_set_max_colors(attr, C.int(colors)))
}

func GetMaxColors(attr LiqAttr) int {
	return int(C.liq_get_max_colors(attr))
}

func SetSpeed(attr LiqAttr, speed int) LiqError {
	return LiqError(C.liq_set_speed(attr, C.int(speed)))
}

func GetSpeed(attr LiqAttr) int {
	return int(C.liq_get_speed(attr))
}

func SetMinPosterization(attr LiqAttr, bits int) LiqError {
	return LiqError(C.liq_set_min_posterization(attr, C.int(bits)))
}

func GetMinPosterization(attr LiqAttr) int {
	return int(C.liq_get_min_posterization(attr))
}

func SetQuality(attr LiqAttr, min int, max int) LiqError {
	return LiqError(C.liq_set_quality(attr, C.int(min), C.int(max)))
}

func GetMinQuantity(attr LiqAttr) int {
	return int(C.liq_get_min_quality(attr))
}

func GetMaxQuantity(attr LiqAttr) int {
	return int(C.liq_get_max_quality(attr))
}

// -------
// image creation
// -------

func CreateImageRGBA(attr LiqAttr, pixels *uint8, width int, height int, gamma float64) LiqImage {
	return C.liq_image_create_rgba(attr, unsafe.Pointer(pixels), C.int(width), C.int(height), C.double(gamma))
}

func DestroyImage(image LiqImage) {
	C.liq_image_destroy(image)
}

// -------
// image controls
// -------

func AddFixedColor(image LiqImage, color Color) {
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

func QuantizeImage(attr LiqAttr, image LiqImage, result *LiqResult) LiqError {
	return LiqError(C.liq_image_quantize(image, attr, (**C.struct_liq_result)(result)))
}

func DestroyResult(result LiqResult) {
	C.liq_result_destroy(result)
}

// -------
// quantization results
// -------

func GetPalette(result LiqResult) LiqPalette {
	return C.liq_get_palette(result)
}

func WriteRemappedImage(result LiqResult, image LiqImage, buffer *uint8, bufferSize uint64) LiqError {
	return LiqError(C.liq_write_remapped_image(result, image, unsafe.Pointer(buffer), C.ulong(bufferSize)))
}
