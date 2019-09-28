package forge_test

import (
	"katolus/image-analysis/forge"
)

func ExampleForgeRev() {

	picturePath := "./test_images/Capture.jpg"
	img := forge.GetPixels(picturePath)

	img.ForgeRev("./results")

	// Output: A reverse image under './results/revCapture.jpg' has been created.
}

func ExampleForgeGrad() {

	picturePath := "./test_images/Capture.jpg"
	img := forge.GetPixels(picturePath)

	img.ForgeGrad("./results")

	// Output: A gradient image under './results/revCapture.jpg' has been created.
}

func ExampleAnalyze() {
	results := forge.Analyze("./test_images")

	results[0].Forge("./results")
	// Output: Finished analyzing ./test_images and found 5 images
	// Found matching pair of IMG_20180907_095424_Cropped.jpg in IMG_20180907_095424.jpg
	// A match image under ./results/Capture_IN_revCapture_DIRECTORY_844_Y_X_0_0.jpg has been created.
}
