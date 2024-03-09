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

// -------
// attr creation
// -------

func Test_CopyAttr(t *testing.T) {
	attr := CreateAttr()
	attrCopy := CopyAttr(attr)

	if attr == nil || attrCopy == nil {
		t.Error("error copying attr")
	}
}

// -------
// quality controls
// -------

func Test_SetMaxColors(t *testing.T) {
	attr := CreateAttr()
	expected := 2
	SetMaxColors(attr, expected)
	result := GetMaxColors(attr)
	DestroyAttr(attr)

	if result != expected {
		t.Errorf("have = %d , want %d\n", result, expected)
	}
}

func Test_SetSpeed(t *testing.T) {
	attr := CreateAttr()
	speedSet := 2
	err := SetSpeed(attr, speedSet)
	if err != 0 {
		t.Errorf("error setting speed")
	}
	speedGet := GetSpeed(attr)
	DestroyAttr(attr)

	if speedSet != speedGet {
		t.Errorf("speed: have = %d , want = %d\n", speedGet, speedSet)
	}
}

func Test_SetMinPosterization(t *testing.T) {
	attr := CreateAttr()
	minPosterizationSet := 2
	err := SetMinPosterization(attr, minPosterizationSet)
	if err != 0 {
		t.Errorf("error setting min posterization")
	}
	minPosterizationGet := GetMinPosterization(attr)
	if minPosterizationSet != minPosterizationGet {
		t.Errorf("min posterization: have = %d , want = %d\n", minPosterizationGet, minPosterizationSet)
	}
}

func Test_SetQuality(t *testing.T) {
	attr := CreateAttr()
	minSet := 2
	maxSet := 50
	err := SetQuality(attr, minSet, maxSet)
	if err != 0 {
		t.Errorf("error setting quality")
	}
	minGet := GetMinQuantity(attr)
	maxGet := GetMaxQuantity(attr)
	DestroyAttr(attr)

	if minGet != minSet {
		t.Errorf("min: have = %d , want = %d\n", minSet, minGet)
	}
	if maxGet != maxSet {
		t.Errorf("max: have = %d , want = %d\n", maxSet, maxGet)
	}
}

// -------
// image controls
// -------

func Test_AddFixedColor(t *testing.T) {
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

	var result cLiqResult
	err := QuantizeImage(attr, image, &result)
	if err != 0 {
		t.Errorf("error quantizing image")
	}

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

func Test_SimplePNGQuant(t *testing.T) {
	raw, width, height := loadRawData("./images/example.png")

	cattr := CreateAttr()
	SetQuality(cattr, 0, 25)
	SetSpeed(cattr, 1)
	cimage := CreateImageRGBA(cattr, &raw[0], width, height, 0)

	var cresult cLiqResult
	_ = QuantizeImage(cattr, cimage, &cresult)

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
