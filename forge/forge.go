// Copyright 2018 Katolus. All rights reserved.

// Package forge implements a simple image analysis package. It allows to
// do basic image comparison, flipped , lowQuality and gradient copies.
package forge

import (
	"fmt"
)

type pixel struct {
	r, g, b, a uint32
}

// Result - the result instance of images comparison
type Result struct {
	needle, haystack          *Image
	hIdx, avgDiff, totalMatch int
	reverse                   bool
}

// Image - Main object gathering basic information about the image.
type Image struct {
	name   string
	path   string
	pixels []pixel
	width  int
	height int
}

type gImage struct {
	sourceImage *Image
	gradMatrix  []bool
}

func (p pixel) RGBA() (r, g, b, a uint32) {
	return p.r, p.g, p.b, p.a
}

// Analyze - Scans a directory and returns a slice of results of matching images
func Analyze(path string) []Result {

	// Load images into a channel
	ch := loadImages(path)

	// Maps images from a channel into slice of image pointers
	images := getImages(ch)
	fmt.Printf("Finished analyzing %s and found %d images\n", path, len(images))

	// Compare images and return channel of Results
	chResults := compare(images)

	var results []Result

	for r := range chResults {
		fmt.Printf("Found matching pair of %s in %s\n", r.needle.name, r.haystack.name)
		results = append(results, r)
	}

	return results
}
