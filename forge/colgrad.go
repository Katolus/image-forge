package forge

import (
	"image/color"
	"image/color/palette"
)

// Helper function to compare pixel with each other
func (p pixel) isSame(cP pixel) bool {
	isTheSameColor := false
	fColor := getP9RGBA
	var c1 color.RGBA
	var c2 color.RGBA
	c1.R, c1.G, c1.B, c1.A = fColor(p)
	c2.R, c2.G, c2.B, c2.A = fColor(cP)

	if c1.R == c2.R && c1.G == c2.G && c1.B == c2.B && c1.A == c2.A {
		isTheSameColor = true
	}
	return isTheSameColor
}

// Return the RGBA values for the WebSafe representation
func getWebSafeRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	p := palette.WebSafe
	r, g, b, a := color.Palette.Convert(p, c).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}

// Return the RGBA values for the P9 representation
func getP9RGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	p := palette.Plan9
	r, g, b, a := color.Palette.Convert(p, c).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}

// Helper function to better understand the complexity of a image's colors
func (img Image) groupPixelOverColor() map[color.Color]int {
	col := img.height
	row := img.width

	m := make(map[color.Color]int)
	for i := 0; i < col*row; i++ {

		r, g, b, a := getP9RGBA(img.pixels[i])
		currentColor := color.RGBA{r, g, b, a}

		if _, ok := m[currentColor]; ok {
			m[currentColor]++
		} else {
			m[currentColor] = 0
		}
	}

	return m
}

// Return gImage for a given image
func (img Image) getColorGradient() gImage {
	col := img.height
	row := img.width

	var r, g, b, a uint32
	var previousColor pixel

	var gImg gImage
	gA := make([]bool, col*row)
	for i := 0; i < col*row; i++ {
		x := i % col

		r, g, b, a = img.pixels[i].RGBA()
		currentColor := pixel{r, g, b, a}

		// Set gradient to false on x = 0 as nothing to compare to
		if x == 0 {
			gA[i] = false
			previousColor = currentColor
			continue
		}

		if currentColor.isSame(previousColor) {
			gA[i] = true
		}
		previousColor = currentColor

	}
	gImg.sourceImage = &img
	gImg.gradMatrix = gA
	return gImg
}
