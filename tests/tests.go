package tests

// #cgo LDFLAGS: -L../libimagequant/target/release/ -limagequant_sys -lm
// #include "../abi.c"
import "C"
import (
	"testing"
)

func Liq_version(t *testing.T) {
	expected := C.uint(40202)
	result := C.liq_version()
	if result != expected {
		t.Fatalf("version = %d, want %d", result, expected)
	}
}
