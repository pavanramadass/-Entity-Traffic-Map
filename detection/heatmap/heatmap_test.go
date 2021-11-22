package heatmap

import (
	"bytes"
	"detection/data"
	"detection/metadata"
	"encoding/base64"
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

func TestNewHeatmap(t *testing.T) {

}

func TestGenerateHeatmap(t *testing.T) {

}

func TestgetPixelCounts(t *testing.T) {

}

func TestgenerateHeatmapImage(t *testing.T) {

}

func TestpixelCountToColor(t *testing.T) {

}
