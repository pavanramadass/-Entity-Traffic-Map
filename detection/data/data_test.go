package data

import "testing"

func TestImportData(t *testing.T) {
	wanted = nil
	got = ImportData("file_test.json")

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

func TestExportData(t *testing.T) {
	wanted = nil
	got = ExportData("file_test.json")

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

// TestDatacollectionMinutes tests for collecting 1 to 2 minutes of data
func TestDataCollectionMinutes(t *testing.T) {

}

// TestDatacollectionHours tests for collecting 1 to 2 hours of data
func TestDataCollectionHours(t *testing.T) {

}

// TestDatacollectionDays tests for collecting 1 to 2 days of data
func TestDataCollectionDays(t *testing.T) {

}
