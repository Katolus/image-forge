package forge

import (
	"sync"
	"testing"
)

func TestPixelDiffForSamePixel(t *testing.T) {

	pixel1 := pixel{165, 140, 155, 20}
	pixel2 := pixel{165, 140, 155, 20}

	if diff := pixelDiff(pixel2, pixel1); diff != 0 {
		t.Fatal("Same pixel should always return 0 diff.")
	}
}

func TestPixelDiffForDifferentPixel(t *testing.T) {
	pixel1 := pixel{165, 140, 155, 20}
	pixel2 := pixel{46092, 38922, 32637, 65535}

	if diff := pixelDiff(pixel2, pixel1); diff != 182706 {
		t.Fatal("Same pixel should always return 182706 diff.")
	}
}

func TestCompareSequenceForSamePicture(t *testing.T) {
	var wg sync.WaitGroup
	const image1Dir = "./test_images/Capture.jpg"
	const image2Dir = "./test_images/Capture.jpg"

	hIdx := 0 // Iteration index from a match between a needle and a haystack picture

	img1 := GetPixels(image1Dir) // Needle image
	img2 := GetPixels(image2Dir) // Haystack image

	ch := make(chan Result) // Channel of results

	wg.Add(1)
	go func(n, h *Image, i int, ch chan Result) {
		defer wg.Done()
		compareSequence(n, h, i, ch)
	}(img1, img2, hIdx, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	if match, ok := <-ch; !ok || match.avgDiff != 0 {
		t.Fatal("There should be a match between same pictures.")
	}
}

func TestCompareSequenceForSameReversePicture(t *testing.T) {
	var wg sync.WaitGroup
	const image1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	const image2Dir = "./test_images/revIMG_20180907_095424_Cropped.jpg"

	hIdx := 0 // Iteration index from a match between a needle and a haystack picture

	img1 := GetPixels(image1Dir) // Needle image
	img2 := GetPixels(image2Dir) // Haystack image

	ch := make(chan Result) // Channel of results

	wg.Add(1)
	go func(n, h *Image, i int, ch chan Result) {
		defer wg.Done()
		compareSequenceReverse(n, h, i, ch)
	}(img1, img2, hIdx, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	if match, ok := <-ch; !ok || match.avgDiff > int(threshold) {
		t.Logf("%v", match)
		t.Fatal("There should be a match between same pictures.")
	}
}

func TestCompareSequenceCroppedMatchedPictures(t *testing.T) {
	var wg sync.WaitGroup
	const image1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	const image2Dir = "./test_images/IMG_20180907_095424.jpg"

	hIdx := 5859148 // Iteration index from a match between a needle and a haystack picture

	img1 := GetPixels(image1Dir) // Needle image
	img2 := GetPixels(image2Dir) // Haystack image

	ch := make(chan Result) // Channel of results

	wg.Add(1)
	go func(n, h *Image, i int, ch chan Result) {
		defer wg.Done()
		compareSequence(n, h, i, ch)
	}(img1, img2, hIdx, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	if _, ok := <-ch; !ok {
		t.Fatal("CompareSequence should Result in an occupied channel.")
	}
}

func TestComparePixelsForCroppedMatchedPictures(t *testing.T) {
	var wg sync.WaitGroup

	const image1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	const image2Dir = "./test_images/IMG_20180907_095424.jpg"

	img1 := GetPixels(image1Dir) // Needle image
	img2 := GetPixels(image2Dir) // Haystack image

	ch := make(chan Result) // Channel of results

	wg.Add(1)
	go func(n, h *Image, ch chan Result) {
		defer wg.Done()
		comparePixels(n, h, ch)
	}(img1, img2, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	if _, ok := <-ch; !ok {
		t.Fatal("There should be a match between cropped pictures.")
	}
}

func TestComparePixelsForReversedPictures(t *testing.T) {
	var wg sync.WaitGroup
	const image1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	const image2Dir = "./test_images/IMG_20180907_095424.jpg"

	img1 := GetPixels(image1Dir) // Needle image
	img2 := GetPixels(image2Dir) // Haystack image

	ch := make(chan Result) // Channel of results

	wg.Add(1)
	go func(n, h *Image, ch chan Result) {
		defer wg.Done()
		comparePixels(n, h, ch)
	}(img1, img2, ch)
	// comparePixels(img1, img2, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	if _, ok := <-ch; !ok {
		t.Fatal("There should be a match between cropped pictures.")
	}
}

func TestCompareValid(t *testing.T) {

	const image1Dir = "./test_images/IMG_20180907_095424_Cropped.jpg"
	const image2Dir = "./test_images/IMG_20180907_095424.jpg"

	img1 := GetPixels(image1Dir) // Needle image
	img2 := GetPixels(image2Dir) // Haystack image

	images := []*Image{img1, img2}

	ch := compare(images)

	if _, ok := <-ch; !ok {
		t.Fatal("Compare should return a match.")
	}
}
