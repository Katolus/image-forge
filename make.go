package forge

import (
	"fmt"
	stdimage "image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Forge - takes a single result struct and creates a `.jpg` overlaying image with matches, at the storePath variable
func (r Result) Forge(storePath string) {

	n := r.needle
	h := r.haystack

	of, err := os.Open(h.path)
	if err != nil {
		log.Fatalln("Error opening file in mkImg", err.Error())
	}

	defer of.Close()

	fileName := []string{storePath, "/", strings.Trim(n.name, ".jpg"), "_IN_", strings.Trim(h.name, ".jpg"), "_DIRECTORY_", strconv.Itoa(r.avgDiff), "_Y_X_", strconv.Itoa(r.hIdx / (h.width)), "_", strconv.Itoa(r.hIdx % (h.width)), ".jpg"}
	nf, err := os.Create(strings.Join(fileName, ""))

	if err != nil {
		log.Fatalln("Error creating a new file", err.Error())
	}

	defer nf.Close()

	dof, err := jpeg.Decode(of)
	if err != nil {
		log.Fatalln("Error decoding haystack", err.Error())
	}

	m := stdimage.NewRGBA(dof.Bounds())

	hIdx := r.hIdx

	for i := 0; i < h.height*h.width; i++ {
		m.Set(i%h.width, i/h.width, h.pixels[i])
	}

	for i := 0; i < n.height*n.width; i++ {

		// Check the difference between pixels
		var nIdx int
		if r.reverse {
			nIdx = func(x int) int {
				return (n.width - 1) - (x % n.width) + (n.width * (x / n.width))
			}(i)
		} else {
			nIdx = i
		}
		diff := pixelDiff(n.pixels[nIdx], h.pixels[hIdx])

		var d uint8 // down sizing color variable from unit32
		// Setting color of the match frame
		if diff < threshold {
			d = uint8(255 * (1 - (diff / threshold))) // Shades of color
		} else {
			d = 0 // Black color
		}

		// Create needle frame
		nX := i % (n.width)
		nY := i / (n.width)

		// Const color for 5 pixels on on left and right
		if nX < 5 || ((n.width - nX) < 5) {
			d = 160
		}
		// Const color for 5 pixels on top and bottom
		if nY == 0 || (nY == (n.height - 1)) {
			d = 160
		}

		hX := hIdx % (h.width)
		hY := hIdx / (h.width)

		a := uint8(math.Floor(255 * .5))
		m.Set(hX, hY, color.RGBA{d, d, d, a})

		// Next pixel is a new row add diff between end of needle and haystack
		if ((i + 1) % n.width) == 0 {
			hIdx += (h.width - n.width)
		}

		hIdx++
	}
	jpeg.Encode(nf, m, nil)
	fmt.Printf("A match image under %s has been created.\n", strings.Join(fileName, ""))
}

// ForgeRev - Produces a reversed image in a given path. Name of the file is based on the initial img file name.
//
// i.e. - revImgDir = "./results/"
func (img Image) ForgeRev(revImgDir string) {

	of, err := os.Open(img.path)

	if err != nil {
		log.Fatalf("Error :%s", err)
	}
	defer of.Close()

	rImg, err := jpeg.Decode(of)

	if err != nil {
		log.Fatalf("Error decoding read image. Error -> %s", err)
	}
	fileName := revImgDir + "/rev" + img.name
	nf, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Error creating new file %s. Error -> %s", fileName, err)
	}

	col := rImg.Bounds().Max.X
	row := rImg.Bounds().Max.Y
	size := col * row

	m := stdimage.NewRGBA(rImg.Bounds())

	for i := 0; i < size; i++ {
		x := (i % col)
		y := (i / col)
		r, g, b, a := rImg.At(x, y).RGBA()
		m.Set(col-1-x, y, pixel{r, g, b, a})
	}
	jpeg.Encode(nf, m, nil)
	fmt.Printf("A reverse image under %s has been created.\n", fileName)
}

// ForgeLQ - Produces a low quality (8bit) image with a limit range of color, dependend of palette (p9/ws : default p9)
func (img Image) ForgeLQ(lqImgDir string, paletteType string) {
	of, err := os.Open(img.path)

	if err != nil {
		log.Fatalf("Error :%s", err)
	}
	defer of.Close()

	rImg, err := jpeg.Decode(of)

	if err != nil {
		log.Fatalf("Error decoding read image. Error -> %s", err)
	}

	var RGBA func(color.Color) (r, g, b, a uint8)
	// Default Palette
	if paletteType == "ws" {
		RGBA = getWebSafeRGBA
	} else {
		RGBA = getP9RGBA
	}

	fileName := lqImgDir + paletteType + img.name
	nf, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Error creating new file %s. Error -> %s", fileName, err)
	}

	col := img.width
	row := img.height
	size := col * row

	m := stdimage.NewRGBA(rImg.Bounds())

	for i := 0; i < size; i++ {
		x := (i % col)
		y := (i / col)
		r, g, b, a := RGBA(rImg.At(x, y))
		m.Set(x, y, color.RGBA{r, g, b, a})
	}
	jpeg.Encode(nf, m, nil)
	fmt.Printf("A low quality image under %s has been created.\n", fileName)
}

// ForgeGrad - Produces an image of the gradient of diff change between pixels
//
// i.e. - endPath = "./results/gradImage.jpg"
func (img Image) ForgeGrad(endPath string) {

	nf, err := os.Create(endPath)
	if err != nil {
		log.Fatalf("Error creating new file %s. Error -> %s", endPath, err)
	}

	defer nf.Close()

	cGradient := img.getColorGradient()
	orImg := loadImage(img.path)
	m := stdimage.NewRGBA(orImg.Bounds())
	col := img.width

	size := len(cGradient.gradMatrix)
	for i := 0; i < size; i++ {
		if cGradient.gradMatrix[i] {
			m.Set(i%col, i/col, color.RGBA{255, 255, 255, 255})
		} else {
			m.Set(i%col, i/col, color.RGBA{0, 0, 0, 0})
		}
	}
	jpeg.Encode(nf, m, nil)
	fmt.Printf("A gradient image under %s has been created.\n", endPath)
}

// Util function for retriving name of a file for a given directory
func getImgName(imagePath string) string {
	return filepath.Base(imagePath)
}
