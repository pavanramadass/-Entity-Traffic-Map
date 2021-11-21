package heatmap

import (
	"encoding/csv"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

// Metadata struct stores parsed JSON metadata.
type Metadata struct {
	DataFileName string `json:"data_filename"`
	ImageWidth   int    `json:"image_width"`
	ImageHeight  int    `json:"image_height"`
}

// GenerateHeatmapFromCSVFile does ... 
func GenerateHeatmapFromCSVFile(metaDataFileName string) {
	// Open and read metadata file
	metaDataFile, err := os.Open(metaDataFileName)
	if err != nil {
		log.Fatal("Error: Failed to open metadata file "+metaDataFileName, err)
	}
	metaDataJson, err := ioutil.ReadAll(metaDataFile)
	if err != nil {
		log.Fatal("Error: Failed to read metadata file "+metaDataFileName+": ", err)
	}
	metaDataFile.Close()

	// Parse metadata
	var metadata Metadata
	err = json.Unmarshal(metaDataJson, &metadata)
	if err != nil {
		log.Fatal("Error: Failed to parse metadata JSON: ", err)
	}

	// Open data file
	csvFile, err := os.Open(metadata.DataFileName)
	if err != nil {
		log.Fatal("Error: Failed to open data file "+metaDataFileName+": ", err)
	}

	// Parse CSV data
	reader := csv.NewReader(csvFile)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error: Failed to parse CSV file "+metaDataFileName+": ", err)
	}
	csvFile.Close()

	// Get counts for each pixel and generate heatmap image
	pixelCounts, minCount, maxCount := getPixelCounts(data, metadata.ImageWidth, metadata.ImageHeight)
	heatmap := generateHeatmapImage(pixelCounts, minCount, maxCount, metadata.ImageWidth, metadata.ImageHeight)

	// Output heatmap to PNG file
	heatmapFile, _ := os.Create("heatmap.png")
	png.Encode(heatmapFile, heatmap)
	heatmapFile.Close()
}

// GetPixelCounts does ... 
func getPixelCounts(data [][]string, imageWidth int, imageHeight int) ([][]int, int, int) {
	// Create 2D array the same size as the image
	pixelCounts := make([][]int, imageWidth)
	for i := range pixelCounts {
		pixelCounts[i] = make([]int, imageHeight)
	}

	// Count the number of data points per pixel.
	// Keep track of min and max counts across all pixels.
	minCount, maxCount := 0, 0
	for _, data := range data {
		x, _ := strconv.Atoi(data[1])
		y, _ := strconv.Atoi(data[2])

		pointCount := pixelCounts[x][y] + 1

		if pointCount > maxCount {
			maxCount = pointCount
		}
		if pointCount < minCount {
			minCount = pointCount
		} else if minCount == 0 {
			minCount = pointCount
		}

		pixelCounts[x][y] = pointCount
	}
	return pixelCounts, minCount, maxCount
}

// GenerateHeatmapImage generates a png image using the image dimensions and pixel counts. 
func generateHeatmapImage(pixelCounts [][]int, minCount int, maxCount int, imageWidth int, imageHeight int) image.Image {
	heatmapImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for x, col := range pixelCounts {
		for y, pointCount := range col {
			heatmapImage.Set(x, y, pixelCountToColor(pointCount, minCount, maxCount))
		}
	}
	return heatmapImage

}

// PixelCountToColor maps the centroid count at each pixel to a corresponding color and intensity. 
func pixelCountToColor(count int, min int, max int) color.RGBA {
	r, g, b := 0, 0, 0

	if count > 0 {
		density := int(math.Round(float64(count) / float64((max-min)+1) * 511))

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