package main

import (
	"fmt"
	"katolus/image-forge/forge"
)

func main() {

	picturePath := "./images/wedding_photo_4.jpeg"
	img := forge.GetPixels(picturePath)

	// Creates a reversed picture.
	img.ForgeRev("./results/")
	// Example output: A reverse image under './results/revCapture.jpg' has been created.

	// Creates a gradient picture.
	img.ForgeGrad("./results/")
	// Example output: A gradient image under './results/revCapture.jpg' has been created.

	// Creates a lower quality picture.
	img.ForgeLQ("./results/", "p9")
	// Example output:

	results := forge.Analyze("./images")

	if len(results) > 0 {
		results[0].Forge("./results/")
	} else {
		fmt.Printf("Didn't find any matching images.")
	}
	// Example output: Finished analyzing ./images and found 5 images
	// Found matching pair of IMG_20180907_095424_Cropped.jpg in IMG_20180907_095424.jpg
	// A match image under ./results/Capture_IN_revCapture_DIRECTORY_844_Y_X_0_0.jpg has been created.
}
