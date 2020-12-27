package forge

import (
	stdimage "image"
	"image/jpeg"
	"os"
	"testing"
)

// func TestMkImg(t *testing.T) {

// 	const image1Dir = "IMG_20180907_095424_Cropped.jpg"
// 	const image2Dir = "IMG_20180907_095424.jpg"

// 	img1 := GetPixels(image1Dir) // Needle image
// 	img2 := GetPixels(image2Dir) // Haystack image

// 	// ch := make(chan Result)
// 	match := Result{img1, img2, 5859148, 2241, 90372}
// 	mkImg(match)

// 	filename := "./results/test-a-nonexistent-file"
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		t.Fatalf("File does not exist")
// 	}
// 	// else {
// 	// os.Remove(filename)
// 	// }

// }

func TestMakeReverseImage(t *testing.T) {

	dir := "./results/"
	GetPixels("./test_images/IMG_20180907_095424_Cropped.jpg").ForgeRev(dir)
	if file, err := os.Stat("./results/revIMG_20180907_095424_Cropped.jpg"); os.IsNotExist(err) {
		t.Fatalf("File -> %s doesn't exist.", file.Name())
	}
}

func TestGetImgName(t *testing.T) {
	if name := getImgName("./test_images/Capture.jpg"); name != "Capture.jpg" {
		t.Fatalf("Image name is not correct. %s != Capture", name)
	}
}

func TestPixelMatchAfterTransformation(t *testing.T) {
	of, err := os.Open("./test_images/Capture.jpg")

	if err != nil {
		t.Fatalf("Error :%s", err)
	}
	defer of.Close()

	const revImgDir = "./results/"

	rImg, err := jpeg.Decode(of)
	m := stdimage.NewRGBA(rImg.Bounds())
	if err != nil {
		t.Fatalf("Error decoding read image. Error -> %s", err)
	}

	r, g, b, a := rImg.At(0, 0).RGBA()
	oPixel := pixel{r, g, b, a}
	m.Set(539, 0, pixel{r, g, b, a})
	r, g, b, a = m.At(539, 0).RGBA()
	fPixel := pixel{r, g, b, a}
	if diff := pixelDiff(oPixel, fPixel); diff >= threshold {
		t.Fatalf("Pixel difference (%f) match not within threshold.", diff)
	}

}

func TestMakeImgWithWSPixels(t *testing.T) {

	img := GetPixels("./test_images/IMG_20180907_095424_Cropped.jpg")

	dir := "./results/"
	img.ForgeLQ(dir, "ws")

	if file, err := os.Stat("./results/wsIMG_20180907_095424_Cropped.jpg"); os.IsNotExist(err) {
		t.Fatalf("File -> %s doesn't exist.", file.Name())
	}
}

func TestForgeGradientImage(t *testing.T) {

	img := GetPixels("./test_images/IMG_20180907_095424_Cropped.jpg")
	endPath := "./results/gradientCropped.jpg"
	img.ForgeGrad(endPath)

	if file, err := os.Stat(endPath); os.IsNotExist(err) {
		t.Fatalf("File -> %s doesn't exist.", file.Name())
	}
}
