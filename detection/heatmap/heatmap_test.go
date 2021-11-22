package heatmap

import (
	"testing"
)

const maxY, maxX int = 500, 600

func TestGenerateHeatmapFromCSVFile(t *testing.T) {
	heatmap := NewHeatmap(maxX, maxY)
	heatmap.GenerateHeatmap(StoredData, "test.png")
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

/*
// TestHeatMapMinutes tests the heatmap to graph 1 to 2 mins of data.
func TestHeatMapMinutes(t *testing.T) {

}

// TestHeatMapHours tests the heatmap to graph 1 to 2 hours of data.
func TestHeatMapHours(t *testing.T) {

}

// TestHeatMapDays tests the heatmap to graph 1 to 2 days of data.
func TestHeatMapDays(t *testing.T) {

}
*/
