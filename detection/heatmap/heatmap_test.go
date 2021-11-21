package heatMap

import "testing"

func TestGenerateHeatmapFromCSVFile(t *testing.T) {
	file = "metadataTest.csv"
	wanted = true
	got = GenerateHeatmapFromCSVFile(file)

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted) 
	}
}

func TestGenerateHeatmapImage(t *testing.T) {
	file = "metadataTest.csv"
	wanted = true
	GenerateHeatmapFromCSVFile(file)
	pixelCounts, minCount, maxCount = getPixelCounts()
	heatmap, got = generateHeatmapImage(pixelCounts, minCount, maxCount, imageWidth int, imageHeight int)

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted) 
	}
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