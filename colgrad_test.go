package forge

import (
	"image/color"
	"image/color/palette"
	"testing"
)

func TestIfPixelsAreComparable(t *testing.T) {

	const img1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	cX := 367
	cY := 353

	img := loadImage(img1Dir)
	// fmt.Println(img.At(cX, cY).RGBA())
	// fmt.Println(img.At(cX, cY+1).RGBA())
	// fmt.Println(img.At(cX+1, cY).RGBA())
	// fmt.Println(img.At(cX, cY))
	// fmt.Println(img.At(cX, cY+1))
	// fmt.Println(img.At(cX+1, cY))
	// fmt.Println(colorDiff(img.At(cX, cY), img.At(cX+1, cY)))
	// fmt.Println(colorDiff(img.At(cX, cY), img.At(cX, cY+1)))

	if colorDiff(img.At(cX, cY), img.At(cX+1, cY)) > threshold {
		t.Fatalf("Pixel not comparable under current threshold = %f", threshold)
	}
}
func TestPixelToRGBAPaletteConversion(t *testing.T) {

	const img1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	cX := 367
	cY := 353
	img := loadImage(img1Dir)
	p := palette.WebSafe
	p9 := palette.Plan9
	// fmt.Println(img.At(cX, cY))
	// fmt.Println(img.At(cX, cY).RGBA())
	// fmt.Println(color.Palette.Convert(p, img.At(cX, cY)))
	// fmt.Println(p[color.Palette.Index(p, img.At(cX, cY))])
	// fmt.Println(p[color.Palette.Index(p, img.At(cX, cY+1))])
	// fmt.Println(p9[color.Palette.Index(p9, img.At(cX, cY))])
	// fmt.Println(p9[color.Palette.Index(p9, img.At(cX, cY+1))])

	wsRGBA := color.Palette.Convert(p, img.At(cX, cY))
	p9RGBA := color.Palette.Convert(p9, img.At(cX, cY))
	if wsRGBA == p9RGBA {
		// A bit useless test
		t.Logf("Same pixel on both standards!")
	}
	// Conversion produces different pixel descriptors

}

func TestGetWebSafeRGBA(t *testing.T) {
	// Commented out - proof that on the same pixel CAN not match because of uint32 conversion
	// Altough pixel should match on WebSafe RGBA
	// p := palette.Plan9
	compareColor := pixel{41576, 7773, 10989, 65535}
	// fmt.Println(compareColor)
	expectedValues := color.RGBA{153, 51, 51, 255}
	// r, g, b, a := expectedValues.RGBA()
	// fmt.Println(expectedValues)
	// fmt.Println(pixel{r, g, b, a})
	// fmt.Println(colorDiff(compareColor, pixel{r, g, b, a}))
	// fmt.Println(p[color.Palette.Index(p, pixel{r, g, b, a})])
	r, g, b, a := getWebSafeRGBA(compareColor)
	convertedColor := color.RGBA{r, g, b, a}
	if convertedColor != expectedValues {
		t.Fatalf("I don't understand this conversion!")
	}
}

func TestGroupPixelOverColor(t *testing.T) {

	const img1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	img := GetPixels(img1Dir)

	if imgMap := img.groupPixelOverColor(); len(imgMap) == 0 {
		t.Logf("Log map -> %v", imgMap)
		t.Fatal("Map cannot be empty")
	}
}

// func TestBitConversion(t *testing.T) {
// 	const img1Dir = "./test_images/IMG_20180907_095424.jpg"
// 	img := loadImage(img1Dir)

// 	r, g, b, a := img.At(22, 22).RGBA()
// 	fmt.Println(img.At(22, 22))
// 	fmt.Printf("%b | %b | %b\n", 93, 135, 113)
// 	fmt.Printf("%T | %T | %T | %T\n", r, g, b, a)
// 	fmt.Printf("%d | %d | %d | %d\n", r, g, b, a)
// 	fmt.Printf("%b | %b | %b | %b\n", r, g, b, a)
// 	fmt.Printf("%b | %b | %b | %b\n", r<<8, g<<8, b<<8, a<<8)
// 	// r8, g8, b8, a8 := getWebSafeRGBA(img.At(22, 22))
// 	r8, g8, b8, a8 := uint8(r), uint8(g), uint8(b), uint8(a)
// 	fmt.Printf("%T | %T | %T | %T\n", r8, g8, b8, a8)
// 	fmt.Printf("%d | %d | %d | %d\n", r8, g8, b8, a8)
// 	fmt.Printf("%b | %b | %b | %b\n", r8, g8, b8, a8)
// 	fmt.Printf("%b | %b | %b | %b\n", r8<<2, g8<<2, b8<<2, a8<<2)
// 	r16, g16, b16, a16 := uint16(r), uint16(g), uint16(b), uint16(a)
// 	fmt.Printf("%T | %T | %T | %T\n", r16, g16, b16, a16)
// 	fmt.Printf("%d | %d | %d | %d\n", r16, g16, b16, a16)
// 	fmt.Printf("%b | %b | %b | %b\n", r16, g16, b16, a16)
// 	fmt.Printf("%b | %b | %b | %b\n", r16<<2, g16<<2, b16<<2, a16<<15)
// 	r, g, b, a = uint32(r8), uint32(g8), uint32(b8), uint32(a8)
// 	fmt.Printf("%b | %b | %b | %b\n", r, g, b, a)
// 	r, g, b, a = uint32(r16), uint32(g16), uint32(b16), uint32(a16)
// 	fmt.Printf("%b | %b | %b | %b\n", r, g, b, a)
// 	a = 2000
// 	fmt.Printf("%b | %b | %b\n", uint32(a), uint16(a), uint8(a))
// 	fmt.Printf("%d | %d | %d\n", uint32(a), uint16(a), uint8(a))
// 	a = 208
// 	fmt.Printf("%b | %b | %b\n", uint32(a), uint16(a), uint8(a))
// 	fmt.Printf("%d | %d | %d\n", uint32(a), uint16(a), uint8(a))
// }

func TestIsSame(t *testing.T) {

	color1 := pixel{41576, 7773, 10989, 65535}
	color2 := pixel{255, 255, 255, 65535}

	if !color1.isSame(color1) {
		t.Fatalf("Same color should return true")
	}

	if color1.isSame(color2) {
		t.Fatalf("Different color should return false")
	}
}

func TestGetColorGradient(t *testing.T) {
	img := GetPixels("./test_images/IMG_20180907_095424_Cropped.jpg")
	cGradient := img.getColorGradient()

	if cGradient.sourceImage == nil || len(cGradient.gradMatrix) == 0 {
		t.Fatal("Color gradient should create an object with source image pointer and populate the gradient matrix")
	}
}
