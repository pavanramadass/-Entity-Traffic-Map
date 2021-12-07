package heatmap

import (
	"bytes"
	"encoding/base64"
	"entityDetection/detection/data"
	"entityDetection/detection/metadata"
	"fmt"
	"image"
	"testing"
)

func TestGenerateHeatmapFrom4Min(t *testing.T) {
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04_18.json")[0]
	imageBytes, err := base64.StdEncoding.DecodeString(m.BaseImage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var d data.Data
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04_18.json")

	heatmap := NewHeatmap(img.Bounds().Max.X, img.Bounds().Max.Y)
	heatmap.GenerateHeatmap(d.GetData([]int64{}), "test.png")
}

func TestGetPixelCounts(t *testing.T) {
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04_18.json")[0]
	imageBytes, err := base64.StdEncoding.DecodeString(m.BaseImage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	heatmap := NewHeatmap(img.Bounds().Max.X, img.Bounds().Max.Y)

	var d data.Data
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04_18.json")

	pixelCountsActual := heatmap.getPixelCounts(d.GetData([]int64{}))

	if pixelCountsActual != nil {
		fmt.Println("Pixels succesfully found!")
		return
	}
}

// Mathetmatical checs to be implemented, for now non-crashing is sufficient
func TestGenerateHeatmapImage(t *testing.T) {
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04_18.json")[0]
	imageBytes, err := base64.StdEncoding.DecodeString(m.BaseImage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	heatmap := NewHeatmap(img.Bounds().Max.X, img.Bounds().Max.Y)

	var d data.Data
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04:18.json")

	pixelCounts := heatmap.getPixelCounts(d.GetData([]int64{}))

	heatmap.generateHeatmapImage(pixelCounts)
}

// Mathetmatical checs to be implemented, for now non-crashing is sufficient
func TestPixelCountToColor(t *testing.T) {
	var d data.Data
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04:18.json")

	var h Heatmap

	h.pixelCountToColor(20)
}
