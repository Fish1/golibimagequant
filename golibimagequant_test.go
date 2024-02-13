package golibimagequant

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
	"testing"
)

func Test_Version(t *testing.T) {
	var expected uint = 40202
	result := Version()
	if result != expected {
		t.Errorf("result = %d, expected = %d", result, expected)
	}
}

func Test_CreateAttr(t *testing.T) {
	CreateAttr()
}

func Test_SetMaxColors(t *testing.T) {
	attr := CreateAttr()
	expected := 2
	SetMaxColors(attr, expected)
	result := GetMaxColors(attr)
	if result != expected {
		t.Errorf("have = %d , want %d\n", result, expected)
	}
}

func Test_SetQuality(t *testing.T) {
	attr := CreateAttr()
	minSet := 2
	maxSet := 50
	SetQuality(attr, minSet, maxSet)
	minGet := GetMinQuantity(attr)
	maxGet := GetMaxQuantity(attr)
	if minGet != minSet {
		t.Errorf("min: have = %d , want = %d\n", minSet, minGet)
	}
	if maxGet != maxSet {
		t.Errorf("max: have = %d , want = %d\n", maxSet, maxGet)
	}
}

func Test_CreateImageRGBA(t *testing.T) {
	attr := CreateAttr()
	data := [4]uint8{0, 0, 0, 0}
	image := CreateImageRGBA(attr, &data[0], 1, 1, 0)
	firstColor := Color{
		r: 1,
		g: 1,
		b: 1,
		a: 1,
	}
	AddFixedColor(image, firstColor)
	secondColor := Color{
		r: 3,
		g: 3,
		b: 3,
		a: 3,
	}
	AddFixedColor(image, secondColor)

	result := QuantizeImage(attr, image)
	palette := GetPalette(result)

	firstResult := NewColor(&palette.entries[0])
	secondResult := NewColor(&palette.entries[1])

	if firstColor != firstResult {
		t.Fatalf("have = %+v , want = %+v", firstResult, firstColor)
	}

	if secondColor != secondResult {
		t.Fatalf("have = %+v , want = %+v", firstResult, firstColor)
	}
}

func loadRawData(filepath string) ([]uint8, int, int) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	size := image.Bounds()
	width := size.Max.X - size.Min.X
	height := size.Max.Y - size.Min.Y
	raw := make([]uint8, width*height*4)
	index := 0
	for y := size.Min.Y; y < size.Max.Y; y += 1 {
		for x := size.Min.X; x < size.Max.X; x += 1 {
			r, g, b, a := image.At(x, y).RGBA()
			raw[index], raw[index+1], raw[index+2], raw[index+3] = uint8(r), uint8(g), uint8(b), uint8(a)
			index += 4
		}
	}
	return raw, width, height
}

func Test_SimplePNGQuant(t *testing.T) {
	raw, width, height := loadRawData("./images/example.png")

	cattr := CreateAttr()
	cimage := CreateImageRGBA(cattr, &raw[0], width, height, 0)
	cresult := QuantizeImage(cattr, cimage)

	pixels := make([]uint8, width*height)
	WriteRemappedImage(cresult, cimage, &pixels[0], uint64(len(pixels)))

	cpalette := GetPalette(cresult)

	rectangle := image.Rect(0, 0, width, height)
	palette := make(color.Palette, 0)
	for _, entry := range cpalette.entries {
		palette = append(palette, NewRGBA(&entry))
	}

	image := image.NewPaletted(rectangle, palette)
	image.Pix = pixels

	file, err := os.Create("./images/example_compressed.png")
	if err != nil {
		t.Fatalf("error creating file")
	}
	defer file.Close()

	err = png.Encode(file, image)
	if err != nil {
		t.Fatalf("error encoding png")
	}

	DestroyResult(cresult)
	DestroyImage(cimage)
	DestroyAttr(cattr)
}
