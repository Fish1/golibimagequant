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

type LiqAttr C.struct_liq_attr
type LiqImage C.struct_liq_image
type LiqResult C.struct_liq_result

type LiqPalette struct {
	Count   uint32
	Entries [256]LiqColor
}

type LiqColor struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type LiqError int

func NewRGBA(ccolor LiqColor) color.RGBA {
	return color.RGBA{
		R: uint8(ccolor.R),
		G: uint8(ccolor.G),
		B: uint8(ccolor.B),
		A: uint8(ccolor.A),
	}
}

func Version() uint {
	return uint(C.liq_version())
}

// -------
// attr creation
// -------

func CreateAttr() *LiqAttr {
	return (*LiqAttr)(C.liq_attr_create())
}

func CopyAttr(attr *LiqAttr) *LiqAttr {
	return (*LiqAttr)(C.liq_attr_copy((*C.struct_liq_attr)(attr)))
}

func DestroyAttr(attr *LiqAttr) {
	C.liq_attr_destroy((*C.struct_liq_attr)(attr))
}

// -------
// quality controls
// -------

func SetMaxColors(attr *LiqAttr, colors int) LiqError {
	return LiqError(C.liq_set_max_colors((*C.struct_liq_attr)(attr), C.int(colors)))
}

func GetMaxColors(attr *LiqAttr) int {
	return int(C.liq_get_max_colors((*C.struct_liq_attr)(attr)))
}

func SetSpeed(attr *LiqAttr, speed int) LiqError {
	return LiqError(C.liq_set_speed((*C.struct_liq_attr)(attr), C.int(speed)))
}

func GetSpeed(attr *LiqAttr) int {
	return int(C.liq_get_speed((*C.struct_liq_attr)(attr)))
}

func SetMinPosterization(attr *LiqAttr, bits int) LiqError {
	return LiqError(C.liq_set_min_posterization((*C.struct_liq_attr)(attr), C.int(bits)))
}

func GetMinPosterization(attr *LiqAttr) int {
	return int(C.liq_get_min_posterization((*C.struct_liq_attr)(attr)))
}

func SetQuality(attr *LiqAttr, min int, max int) LiqError {
	return LiqError(C.liq_set_quality((*C.struct_liq_attr)(attr), C.int(min), C.int(max)))
}

func GetMinQuantity(attr *LiqAttr) int {
	return int(C.liq_get_min_quality((*C.struct_liq_attr)(attr)))
}

func GetMaxQuantity(attr *LiqAttr) int {
	return int(C.liq_get_max_quality((*C.struct_liq_attr)(attr)))
}

// -------
// image creation
// -------

func CreateImageRGBA(attr *LiqAttr, pixels *uint8, width int, height int, gamma float64) *LiqImage {
	return (*LiqImage)(C.liq_image_create_rgba((*C.struct_liq_attr)(attr), unsafe.Pointer(pixels), C.int(width), C.int(height), C.double(gamma)))
}

func DestroyImage(image *LiqImage) {
	C.liq_image_destroy((*C.struct_liq_image)(image))
}

// -------
// image controls
// -------

func AddFixedColor(image *LiqImage, color LiqColor) {
	C.liq_image_add_fixed_color((*C.struct_liq_image)(image), C.struct_liq_color{
		r: C.uchar(color.R),
		g: C.uchar(color.G),
		b: C.uchar(color.B),
		a: C.uchar(color.A),
	})
}

// -------
// quantization
// -------

func QuantizeImage(attr *LiqAttr, image *LiqImage, result **LiqResult) LiqError {
	return LiqError(C.liq_image_quantize((*C.struct_liq_image)(image), (*C.struct_liq_attr)(attr), (**C.struct_liq_result)(unsafe.Pointer(result))))
}

func DestroyResult(result *LiqResult) {
	C.liq_result_destroy((*C.struct_liq_result)(result))
}

// -------
// quantization results
// -------

func GetPalette(result *LiqResult) *LiqPalette {
	return (*LiqPalette)(unsafe.Pointer(C.liq_get_palette((*C.struct_liq_result)(result))))
}

func WriteRemappedImage(result *LiqResult, image *LiqImage, buffer *uint8, bufferSize uint64) LiqError {
	return LiqError(C.liq_write_remapped_image((*C.struct_liq_result)(result), (*C.struct_liq_image)(image), unsafe.Pointer(buffer), C.ulong(bufferSize)))
}
