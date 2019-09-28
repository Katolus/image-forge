package forge

import (
	"image/color"
	"math"
	"sync"
)

// Minimum threshold of pixel difference that categorizes image for a sequential comparision
const threshold float64 = 1500

// Minimum amount of pixel required in a row in order to progress in the sequential comparison
const minReqPixInRow int = 10

// Compare image with each image in the folder.
func compare(images []*Image) chan Result {

	var wg sync.WaitGroup
	ch := make(chan Result)

	for _, needle := range images {
		for _, haystack := range images {

			if needle.name == haystack.name {
				continue
			}
			if needle.height > haystack.height {
				continue
			}
			if needle.width > haystack.width {
				continue
			}

			wg.Add(1)

			go func(n, h *Image, ch chan Result) {
				defer wg.Done()
				comparePixels(n, h, ch)
			}(needle, haystack, ch)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// Compare pixels for sets of images
func comparePixels(n, h *Image, ch chan Result) {

	var diff float64
	var diffReverse float64
	var wg sync.WaitGroup

	for i := 0; i < h.height*h.width; i++ {

		x := i % h.width
		y := i / h.width

		if h.height-y < n.height {
			break
		}

		// find the diff between the pixels
		diff = pixelDiff(n.pixels[0], h.pixels[i])
		diffReverse = pixelDiff(n.pixels[n.width-1], h.pixels[i])

		if (diff < threshold) && (x <= h.width-n.width) {
			wg.Add(1)

			go func(n, h *Image, i int, ch chan Result) {
				defer wg.Done()
				compareSequence(n, h, i, ch)
			}(n, h, i, ch)
		}

		if (diffReverse < threshold) && (x <= h.width-n.width) {

			wg.Add(1)

			go func(n, h *Image, i int, ch chan Result) {
				defer wg.Done()
				compareSequenceReverse(n, h, i, ch)
			}(n, h, i, ch)
		}
	}

	wg.Wait()
}

// Return a sum of absolute differences for each parameter of a pixel
func pixelDiff(n, h pixel) float64 {
	var diff float64

	diff += math.Abs(float64(n.r) - float64(h.r))
	diff += math.Abs(float64(n.g) - float64(h.g))
	diff += math.Abs(float64(n.b) - float64(h.b))
	diff += math.Abs(float64(n.a) - float64(h.a))

	return diff
}

// Return a sum of absolute differences for each parameter of a color
func colorDiff(nColor, hColor color.Color) float64 {
	var diff float64

	nR, nG, nB, nA := nColor.RGBA()
	hR, hG, hB, hA := hColor.RGBA()
	diff += math.Abs(float64(nR) - float64(hR))
	diff += math.Abs(float64(nG) - float64(hG))
	diff += math.Abs(float64(nB) - float64(hB))
	diff += math.Abs(float64(nA) - float64(hA))

	return diff
}

// Compare a pixels sequentially
func compareSequence(n, h *Image, hIdx int, ch chan Result) {
	compareIml(n, h, hIdx, ch, func(x int) (int, bool) { return x, false })
}

// Compare a pixels sequentially in reverse
func compareSequenceReverse(n, h *Image, hIdx int, ch chan Result) {
	compareIml(n, h, hIdx, ch, func(x int) (int, bool) { return (n.width - 1) - (x % n.width) + (n.width * (x / n.width)), true })
}

// Compare sequentially implementation
func compareIml(n, h *Image, hIdx int, ch chan Result, indexCalc func(int) (int, bool)) {
	hStartPix := hIdx
	var totalMatchedPixels int
	var counter, nIdx int
	var accumulator uint64
	var isReverse bool

	for i := 0; i < n.height*n.width; i++ {

		// So that i works for both reverse and regular iterations
		nIdx, isReverse = indexCalc(i)

		// Each row must have minim number of pixels under threshold
		// (1) did the previous row have less than 10?
		// (2) if in new row, reset counter
		// (3) if this pixel beneath threshold, increment counter

		newRow := (i%n.width == 0)
		notFirstRow := (i/n.width != 0)

		if newRow && notFirstRow {
			hIdx += (h.width - n.width)
		}

		diff := pixelDiff(n.pixels[nIdx], h.pixels[hIdx])

		if (newRow) && (notFirstRow) && counter < minReqPixInRow {
			return
		}
		if newRow {
			counter = 0
		}

		if diff < threshold {
			totalMatchedPixels++
			counter++
		}

		hIdx++
		accumulator += uint64(diff)
	}

	avgDiff := int(accumulator / uint64(n.height*n.width))
	ch <- Result{n, h, hStartPix, avgDiff, totalMatchedPixels, isReverse}
}
