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
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04:18.json")[0]
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
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04:18.json")

	fmt.Println(img.Bounds().Max.X, img.Bounds().Max.Y)
	heatmap := NewHeatmap(img.Bounds().Max.X, img.Bounds().Max.Y)
	heatmap.GenerateHeatmap(d.GetData([]int64{}), "test.png")
}

func TestGetPixelCounts(t *testing.T) {
	pixelCountsExpected = "Need To Fill In"
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04:18.json")[0]

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
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04:18.json")

	pixelCountsActual := getPixelCounts(d.GetData([]int64{}))

	if pixelCountsExpected != pixelCountsActual {
		fmt.Println("Error: pixel counts are not the same!")
		return
	}
}

func TestGenerateHeatmapImage(t *testing.T) {
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04:18.json")[0]

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
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04:18.json")

	pixelCounts := getPixelCounts(d.GetData([]int64{}))

	heatmap := generateHeatmapImage(pixelCounts)
}

func TestPixelCountToColor(t *testing.T) {
	m := metadata.ImportMetadata("/home/pi/Entity-Traffic-Map/detection/meta_2021-1122-411_04:18.json")[0]

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
	d.Import("/home/pi/Entity-Traffic-Map/detection/2021-1122-411_04:18.json")

	pixelCounts := getPixelCounts(d.GetData([]int64{}))

	color := pixelCountToColor(pixelCounts)
}
