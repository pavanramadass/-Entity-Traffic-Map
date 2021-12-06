package heatMap
package png 

import "testing"

// The TestGenerateHeatmapFromCSVFile tests whether a heatmap was generated from a csv file 
// if the return value is true, then the function has successfully generated a heatmap from a csv file
func TestGenerateHeatmapFromCSVFile(t *testing.T) {
	file = "metadataTest.csv"
	wanted = true
	got = GenerateHeatmapFromCSVFile(file)

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted) 
	}
}

// The TestGenerateHeatmapImage tests whether a heatmap image was correctly generated 
func TestGenerateHeatmapImage(t *testing.T) {
	file = "metadataTest.csv"
	heatmap1 = Decode("heatmap_test.png") 
	wanted = true
	GenerateHeatmapFromCSVFile(file)
	pixelCounts, minCount, maxCount = getPixelCounts()
	heatmap2, got = generateHeatmapImage(pixelCounts, minCount, maxCount, imageWidth int, imageHeight int)

	if heatmap1 != heatmap2 {
		t.Errorf("Heatmaps are not equal.") 
	}
}