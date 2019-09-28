package forge

import (
	"sync"
	"testing"
)

func TestGetPathsReturnsPaths(t *testing.T) {
	paths := getPaths("./test_images")

	if len(paths) <= 0 {
		t.Fatalf("Expected no errors and paths to have items (%d).", len(paths))
	}
}

func TestGetPathsReturnsPathsFails(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expecting an error %s.", err)
		}
	}()
	getPaths("./non_existing_path")
}

func TestLoadImage(t *testing.T) {
	const imageDir = "./test_images/Capture.jpg"
	img := loadImage(imageDir)

	if img == nil {
		t.Fatalf("Unable to load image from '%s' path", imageDir)
	}
}

func TestLoadImageName(t *testing.T) {
	const imageDir = "./test_images/Capture.jpg"
	img := loadImage(imageDir)
	bounds := img.Bounds()

	if bounds.Max.X != 540 || bounds.Max.Y != 525 {
		t.Fatalf("Invalid bounds of the image under '%s'", imageDir)
	}
}

func TestGetPixels(t *testing.T) {
	const imageDir = "./test_images/Capture.jpg"
	imgP := GetPixels(imageDir)

	switch true {
	case imgP.name != "Capture.jpg", imgP.path != imageDir, len(imgP.pixels) != 283500, imgP.width != 540, imgP.height != 525:
		t.Fatalf("Failed reading pixels from %s", imageDir)
	}
}

func TestLoadImages(t *testing.T) {
	ch := loadImages("./test_images/")

	if item := <-ch; item == nil {
		t.Fatalf("Failed to read %s from the loaded channel.", item.name)
	}
}

func TestXImages(t *testing.T) {
	var wg sync.WaitGroup

	pixel := []pixel{}
	img := Image{"Name", "./test_images/IMG_20180907_095424_Cropped.jpg", pixel, 10, 10}
	ch := make(chan *Image)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- &img
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()
	images := getImages(ch)

	if len(images) != 1 {
		t.Fatal("Failed collecting images from a channel to an []*Image")
	}

}
