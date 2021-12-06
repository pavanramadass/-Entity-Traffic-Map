package dataCollection

import "testing"

// TestImportData tests the import data function from data.go
// If return value is nil, then the ImportData function had successfully imported the data
func TestImportData(t *testing.T) {
	wanted = nil
	got = ImportData("file_test.json")

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

// TestExportData tests the export data function from data.go
// If return value is nil, then the ExportData function had successfully exported the data
func TestExportData(t *testing.T) {
	wanted = nil
	got = ExportData("file_test.json")

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}
