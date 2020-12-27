package forge

import (
	stdimage "image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Walks over a directory and returns a slice of paths to images with suffix .jpg
func getPaths(directory string) []string {
	var paths []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".jpg") {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		log.Println("Error getting paths", err)
	}

	return paths
}

// Based on the path directory loads the images into a channel using goroutines
func loadImages(directory string) chan *Image {

	paths := getPaths(directory)

	var wg sync.WaitGroup
	ch := make(chan *Image)

	for _, path := range paths {
		wg.Add(1)
		go func(p string) {
			ch <- GetPixels(p)
			wg.Done()

		}(path)
	}

	// Closing a channel indicates to no more values will be added to the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// GetPixels - Given an image path loads the image into a custom image struct
func GetPixels(path string) *Image {

	// Load image
	img := loadImage(path)
	bounds := img.Bounds()
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx() // 4%2592 = 4 / 2500%2592 = 2500
		y := i / bounds.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
	}

	image := Image{
		name:   filepath.Base(path),
		path:   path,
		pixels: pixels,
		width:  bounds.Dx(),
		height: bounds.Dy(),
	}
	return &image
}

func loadImage(filePath string) stdimage.Image {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

// Translate a channel results into a slice of image references
func getImages(ch chan *Image) []*Image {
	var images []*Image

	for imgPointer := range ch {
		images = append(images, imgPointer)
	}

	return images
}
