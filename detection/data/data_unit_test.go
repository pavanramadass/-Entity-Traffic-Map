package data

import (
	"reflect"
	"testing"
)

var stored_data = []Centroid{Centroid{Timestamp: 1636476106424, X: 320, Y: 240}, Centroid{Timestamp: 1636476114667, X: 39, Y: 214}, Centroid{Timestamp: 1636476114709, X: 45, Y: 217}, Centroid{Timestamp: 1636476114757, X: 71, Y: 197}, Centroid{Timestamp: 1636476114805, X: 81, Y: 199}, Centroid{Timestamp: 1636476114852, X: 93, Y: 197}, Centroid{Timestamp: 1636476114901, X: 102, Y: 196}, Centroid{Timestamp: 1636476114948, X: 119, Y: 201}, Centroid{Timestamp: 1636476114989, X: 127, Y: 190}, Centroid{Timestamp: 1636476115030, X: 135, Y: 189}, Centroid{Timestamp: 1636476115075, X: 148, Y: 192}, Centroid{Timestamp: 1636476115118, X: 154, Y: 186}, Centroid{Timestamp: 1636476115158, X: 164, Y: 173}, Centroid{Timestamp: 1636476115199, X: 169, Y: 169}, Centroid{Timestamp: 1636476115238, X: 173, Y: 166}, Centroid{Timestamp: 1636476115279, X: 177, Y: 157}, Centroid{Timestamp: 1636476115321, X: 179, Y: 155}, Centroid{Timestamp: 1636476115362, X: 187, Y: 153}, Centroid{Timestamp: 1636476115403, X: 193, Y: 152}, Centroid{Timestamp: 1636476115443, X: 193, Y: 150}, Centroid{Timestamp: 1636476115483, X: 194, Y: 149}, Centroid{Timestamp: 1636476115524, X: 197, Y: 153}, Centroid{Timestamp: 1636476115564, X: 198, Y: 151}, Centroid{Timestamp: 1636476115604, X: 200, Y: 152}, Centroid{Timestamp: 1636476115645, X: 198, Y: 148}, Centroid{Timestamp: 1636476115686, X: 202, Y: 149}, Centroid{Timestamp: 1636476115727, X: 203, Y: 150}, Centroid{Timestamp: 1636476115767, X: 200, Y: 147}, Centroid{Timestamp: 1636476115808, X: 206, Y: 151}, Centroid{Timestamp: 1636476115849, X: 206, Y: 145}, Centroid{Timestamp: 1636476115890, X: 207, Y: 143}, Centroid{Timestamp: 1636476115930, X: 207, Y: 150}, Centroid{Timestamp: 1636476115971, X: 207, Y: 143}, Centroid{Timestamp: 1636476116011, X: 207, Y: 142}, Centroid{Timestamp: 1636476116053, X: 207, Y: 141}, Centroid{Timestamp: 1636476116093, X: 207, Y: 141}, Centroid{Timestamp: 1636476116134, X: 206, Y: 149}, Centroid{Timestamp: 1636476116181, X: 206, Y: 139}, Centroid{Timestamp: 1636476116239, X: 206, Y: 147}, Centroid{Timestamp: 1636476116285, X: 206, Y: 146}, Centroid{Timestamp: 1636476116325, X: 206, Y: 145}, Centroid{Timestamp: 1636476116367, X: 206, Y: 142}, Centroid{Timestamp: 1636476116407, X: 206, Y: 139}, Centroid{Timestamp: 1636476116448, X: 206, Y: 141}, Centroid{Timestamp: 1636476116489, X: 207, Y: 138}, Centroid{Timestamp: 1636476116530, X: 207, Y: 138}, Centroid{Timestamp: 1636476116570, X: 208, Y: 136}, Centroid{Timestamp: 1636476116611, X: 208, Y: 131}, Centroid{Timestamp: 1636476116652, X: 208, Y: 129}, Centroid{Timestamp: 1636476116693, X: 208, Y: 126}, Centroid{Timestamp: 1636476116733, X: 208, Y: 124}, Centroid{Timestamp: 1636476116774, X: 206, Y: 116}, Centroid{Timestamp: 1636476116815, X: 204, Y: 111}, Centroid{Timestamp: 1636476116856, X: 202, Y: 108}, Centroid{Timestamp: 1636476116897, X: 200, Y: 105}, Centroid{Timestamp: 1636476116937, X: 196, Y: 102}, Centroid{Timestamp: 1636476116978, X: 188, Y: 98}, Centroid{Timestamp: 1636476117020, X: 184, Y: 96}, Centroid{Timestamp: 1636476117063, X: 180, Y: 96}, Centroid{Timestamp: 1636476117104, X: 177, Y: 95}, Centroid{Timestamp: 1636476117144, X: 172, Y: 93}, Centroid{Timestamp: 1636476117186, X: 170, Y: 92}, Centroid{Timestamp: 1636476117227, X: 168, Y: 91}, Centroid{Timestamp: 1636476117268, X: 166, Y: 90}, Centroid{Timestamp: 1636476117309, X: 164, Y: 88}, Centroid{Timestamp: 1636476117351, X: 162, Y: 87}, Centroid{Timestamp: 1636476117393, X: 162, Y: 87}, Centroid{Timestamp: 1636476117433, X: 159, Y: 85}, Centroid{Timestamp: 1636476117477, X: 157, Y: 83}, Centroid{Timestamp: 1636476117519, X: 155, Y: 82}, Centroid{Timestamp: 1636476117565, X: 154, Y: 82}, Centroid{Timestamp: 1636476117621, X: 147, Y: 72}, Centroid{Timestamp: 1636476117670, X: 148, Y: 74}, Centroid{Timestamp: 1636476117720, X: 148, Y: 76}, Centroid{Timestamp: 1636476117769, X: 151, Y: 79}, Centroid{Timestamp: 1636476117811, X: 148, Y: 77}, Centroid{Timestamp: 1636476117853, X: 149, Y: 78}, Centroid{Timestamp: 1636476117896, X: 149, Y: 79}, Centroid{Timestamp: 1636476117937, X: 154, Y: 85}, Centroid{Timestamp: 1636476117979, X: 155, Y: 87}, Centroid{Timestamp: 1636476118021, X: 155, Y: 88}, Centroid{Timestamp: 1636476118081, X: 155, Y: 90}, Centroid{Timestamp: 1636476118123, X: 160, Y: 94}, Centroid{Timestamp: 1636476118166, X: 161, Y: 96}}

func TestImportData(t *testing.T) {
	var data Data
	data.ImportData("test.json")

	if reflect.DeepEqual(data.data, stored_data) {
		t.Errorf("got %q, wanted %q", data.data, stored_data)
	}
}

func TestExportData(t *testing.T) {
	got = ExportData("file_test.json")

	if got != wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}
}

func TestStoreData(t *testing.T) {

}

func TestGetData(t *testing.T) {

}
