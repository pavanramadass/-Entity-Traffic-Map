// Package heatmap takes a dataset and generates a PNG heatmap
package heatmap

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"entityDetection/detection/centroid"
)

// Heatmap stores a heatmap image with its maximum point count
type Heatmap struct {
	maxCount     int
	heatmapImage *image.RGBA
}

// NewHeatmap generates new Heatmap instance.
// Don't compute unless requested to.
func NewHeatmap(imageWidth int, imageHeight int) *Heatmap {
	heatmapImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	return &Heatmap{heatmapImage: heatmapImage}
}

// getRect finds the maximum coordinate from a list of points
func getRect(data []centroid.Centroid) (int, int) {
	bottom := 0
	left := 0
	for _, point := range data {
		if bottom < point.Y {
			bottom = point.Y
		}

		if left < point.X {
			left = point.X
		}
	}
	return left, bottom
}

// GenerateHeatmap chains several helper functions to calculate the heatmap,
// then it writes the resultant image to a PNG. If the heatmap does not have
// preset image dimensions it searches through the data for the maximum X and Y
// coordinates and uses those to determine the image dimensions.
func (h *Heatmap) GenerateHeatmap(data []centroid.Centroid, destination string) string {
	if h.heatmapImage == nil {
		x, y := getRect(data)
		h.heatmapImage = image.NewRGBA(image.Rect(0, 0, x, y))
	}

	pixelCounts := h.getPixelCounts(data)
	h.generateHeatmapImage(pixelCounts)

	// Output heatmap to PNG file
	heatmapFile, err := os.Create(destination)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	png.Encode(heatmapFile, h.heatmapImage)
	heatmapFile.Close()

	return destination
}

// GetPixelCounts generates a 2D array of the same dimensions as the image.
// It then counts the presence of centroids at each pixel in the given data.
// It keeps track of min and max counts across all pixels.
func (h *Heatmap) getPixelCounts(locations []centroid.Centroid) [][]int {
	pixelCounts := make([][]int, h.heatmapImage.Rect.Size().X)
	for i := range pixelCounts {
		pixelCounts[i] = make([]int, h.heatmapImage.Rect.Size().Y)
	}

	for _, loc := range locations {
		pixelCounts[loc.X][loc.Y] += 1
		pointCount := pixelCounts[loc.X][loc.Y]

		if pointCount > h.maxCount {
			h.maxCount = pointCount
		}
	}
	return pixelCounts
}

// generateHeatmapImage generates a PNG image using the image dimensions and pixel counts.
func (h *Heatmap) generateHeatmapImage(pixelCounts [][]int) {
	for x, col := range pixelCounts {
		for y, count := range col {
			h.heatmapImage.Set(x, y, h.pixelCountToColor(count))
		}
	}
}

// pixelCountToColor maps the centroid count at each pixel to a corresponding color and intensity.
func (h *Heatmap) pixelCountToColor(count int) color.RGBA {
	r, g, b := 0, 0, 0

	if count > 0 {
		density := int(math.Round(float64(count) / float64(h.maxCount+1) * 511))
		if density < 256 {
			r = density
			g = 255
		} else {
			r = 255
			g = 255 - (density - 256)
		}
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
