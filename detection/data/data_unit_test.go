package data

import (
	"detection/centroid"
	"image"
	"os"
	"reflect"
	"testing"
	"time"
)

const tempFile = "temp.json"
const testData = "test_data.json"
const diffTestData = "edited_test_data.json"
const start int64 = 1636476114709
const end int64 = 1636476118081
const step int64 = 3

var stored_data = []centroid.Centroid{{Timestamp: 1636476106424, X: 320, Y: 240}, {Timestamp: 1636476114667, X: 39, Y: 214}, {Timestamp: 1636476114709, X: 45, Y: 217}, {Timestamp: 1636476114757, X: 71, Y: 197}, {Timestamp: 1636476114805, X: 81, Y: 199}, {Timestamp: 1636476114852, X: 93, Y: 197}, {Timestamp: 1636476114901, X: 102, Y: 196}, {Timestamp: 1636476114948, X: 119, Y: 201}, {Timestamp: 1636476114989, X: 127, Y: 190}, {Timestamp: 1636476115030, X: 135, Y: 189}, {Timestamp: 1636476115075, X: 148, Y: 192}, {Timestamp: 1636476115118, X: 154, Y: 186}, {Timestamp: 1636476115158, X: 164, Y: 173}, {Timestamp: 1636476115199, X: 169, Y: 169}, {Timestamp: 1636476115238, X: 173, Y: 166}, {Timestamp: 1636476115279, X: 177, Y: 157}, {Timestamp: 1636476115321, X: 179, Y: 155}, {Timestamp: 1636476115362, X: 187, Y: 153}, {Timestamp: 1636476115403, X: 193, Y: 152}, {Timestamp: 1636476115443, X: 193, Y: 150}, {Timestamp: 1636476115483, X: 194, Y: 149}, {Timestamp: 1636476115524, X: 197, Y: 153}, {Timestamp: 1636476115564, X: 198, Y: 151}, {Timestamp: 1636476115604, X: 200, Y: 152}, {Timestamp: 1636476115645, X: 198, Y: 148}, {Timestamp: 1636476115686, X: 202, Y: 149}, {Timestamp: 1636476115727, X: 203, Y: 150}, {Timestamp: 1636476115767, X: 200, Y: 147}, {Timestamp: 1636476115808, X: 206, Y: 151}, {Timestamp: 1636476115849, X: 206, Y: 145}, {Timestamp: 1636476115890, X: 207, Y: 143}, {Timestamp: 1636476115930, X: 207, Y: 150}, {Timestamp: 1636476115971, X: 207, Y: 143}, {Timestamp: 1636476116011, X: 207, Y: 142}, {Timestamp: 1636476116053, X: 207, Y: 141}, {Timestamp: 1636476116093, X: 207, Y: 141}, {Timestamp: 1636476116134, X: 206, Y: 149}, {Timestamp: 1636476116181, X: 206, Y: 139}, {Timestamp: 1636476116239, X: 206, Y: 147}, {Timestamp: 1636476116285, X: 206, Y: 146}, {Timestamp: 1636476116325, X: 206, Y: 145}, {Timestamp: 1636476116367, X: 206, Y: 142}, {Timestamp: 1636476116407, X: 206, Y: 139}, {Timestamp: 1636476116448, X: 206, Y: 141}, {Timestamp: 1636476116489, X: 207, Y: 138}, {Timestamp: 1636476116530, X: 207, Y: 138}, {Timestamp: 1636476116570, X: 208, Y: 136}, {Timestamp: 1636476116611, X: 208, Y: 131}, {Timestamp: 1636476116652, X: 208, Y: 129}, {Timestamp: 1636476116693, X: 208, Y: 126}, {Timestamp: 1636476116733, X: 208, Y: 124}, {Timestamp: 1636476116774, X: 206, Y: 116}, {Timestamp: 1636476116815, X: 204, Y: 111}, {Timestamp: 1636476116856, X: 202, Y: 108}, {Timestamp: 1636476116897, X: 200, Y: 105}, {Timestamp: 1636476116937, X: 196, Y: 102}, {Timestamp: 1636476116978, X: 188, Y: 98}, {Timestamp: 1636476117020, X: 184, Y: 96}, {Timestamp: 1636476117063, X: 180, Y: 96}, {Timestamp: 1636476117104, X: 177, Y: 95}, {Timestamp: 1636476117144, X: 172, Y: 93}, {Timestamp: 1636476117186, X: 170, Y: 92}, {Timestamp: 1636476117227, X: 168, Y: 91}, {Timestamp: 1636476117268, X: 166, Y: 90}, {Timestamp: 1636476117309, X: 164, Y: 88}, {Timestamp: 1636476117351, X: 162, Y: 87}, {Timestamp: 1636476117393, X: 162, Y: 87}, {Timestamp: 1636476117433, X: 159, Y: 85}, {Timestamp: 1636476117477, X: 157, Y: 83}, {Timestamp: 1636476117519, X: 155, Y: 82}, {Timestamp: 1636476117565, X: 154, Y: 82}, {Timestamp: 1636476117621, X: 147, Y: 72}, {Timestamp: 1636476117670, X: 148, Y: 74}, {Timestamp: 1636476117720, X: 148, Y: 76}, {Timestamp: 1636476117769, X: 151, Y: 79}, {Timestamp: 1636476117811, X: 148, Y: 77}, {Timestamp: 1636476117853, X: 149, Y: 78}, {Timestamp: 1636476117896, X: 149, Y: 79}, {Timestamp: 1636476117937, X: 154, Y: 85}, {Timestamp: 1636476117979, X: 155, Y: 87}, {Timestamp: 1636476118021, X: 155, Y: 88}, {Timestamp: 1636476118081, X: 155, Y: 90}, {Timestamp: 1636476118123, X: 160, Y: 94}, {Timestamp: 1636476118166, X: 161, Y: 96}}

func TestImportEqual(t *testing.T) {
	var data Data
	data.Import(testData)

	if !reflect.DeepEqual(data.data, stored_data) {
		t.Errorf("Imported data != to %s's data", testData)
		t.Errorf("Length: ImportedData: %d || StoredData: %d", len(data.data), len(stored_data))
	}
}

func TestImportNotEqual(t *testing.T) {
	var data Data
	data.Import(diffTestData)

	if reflect.DeepEqual(data.data, stored_data) {
		t.Errorf("Imported data == to %s's different data", diffTestData)
		t.Errorf("Length: ImportedData: %d || StoredData: %d", len(data.data), len(stored_data))
	}
}

func TestExportData(t *testing.T) {
	var data Data
	data.data = stored_data

	if file, err := os.Stat(tempFile); err == nil {
		t.Errorf("%s existed before creation!", tempFile)
		t.Errorf("Size: %d", file.Size())
	}

	data.ExportData(tempFile)

	if _, err := os.Stat(tempFile); err != nil {
		t.Errorf("%s not created!", tempFile)
		return
	}

	if file, _ := os.Stat(tempFile); file.Size() != 3664 {
		t.Errorf("%s data file of incorrect size!", tempFile)
	}

	os.Remove(tempFile)
}

func TestStoreData(t *testing.T) {
	var data Data
	data.data = stored_data
	pt := image.Point{X: 161, Y: 96}

	diff := make(chan image.Point)
	go data.StoreData(diff)
	diff <- pt
	sec := time.Now().UnixMilli()

	lastPoint := data.data[len(data.data)-1]

	if lastPoint.X != 161 || lastPoint.Y != 96 || lastPoint.Timestamp != sec {
		t.Errorf("Point: %+v with Time: %d != Centroid %+v", pt, sec, lastPoint)
	}
	close(diff)
}

func TestGetDataNoFilters(t *testing.T) {
	var data Data
	data.data = stored_data

	var diffData Data
	diffData.Import(diffTestData)

	if !reflect.DeepEqual(data.GetData([]int64{}), data.data) {
		t.Errorf("GetData with 0 filters != to stored data")
		t.Errorf("Length: GetData(%d) || StoredData: %d", len(data.data), len(data.GetData([]int64{})))
	}
}

func TestGetDataStartFilter(t *testing.T) {
	var data Data
	data.data = stored_data

	gotten := data.GetData([]int64{start})
	if !reflect.DeepEqual(gotten, data.data[2:]) {
		t.Errorf("GetData with start filter = %d != to stored data", start)
		t.Errorf("Length: GetData(%d) || StoredData: %d", len(gotten), len(data.data[2:]))
	}
}

func TestGetDataEndFilter(t *testing.T) {
	var data Data
	data.data = stored_data
	gotten := data.GetData([]int64{-1, end})

	if !reflect.DeepEqual(gotten, data.data[:len(data.data)-2]) {
		t.Errorf("GetData with end filter = %d != to stored data", end)
		t.Errorf("Length: GetData(%d) || StoredData: %d", len(gotten), len(data.data[:len(data.data)-2]))
	}
}

func TestGetDataStepFilter(t *testing.T) {
	var data Data
	data.data = stored_data
	gotten := data.GetData([]int64{-1, -1, step})
	step := int(step)

	for i, v := range gotten {
		if v != data.data[i*step] {
			t.Errorf("GetData with step filter = 3 != every third stored data point")
			t.Errorf("(Index, Value): GetData(%d, %+v) || StoredData(%d, %+v)", i, v, i*step, data.data[i*step])
			break
		}
	}
}

func TestGetDataAllFilters(t *testing.T) {
	var data Data
	data.data = stored_data
	gotten := data.GetData([]int64{start, end, step})
	step := int(step)

	if len(data.data[2:len(data.data)-2])/step+1 != len(gotten) {
		t.Errorf("GetData length != supposed length")
		t.Errorf("Length: GetData(%d) || StoredData: %d", len(gotten), len(data.data[2:len(data.data)-2])/step+1)
	}

	for i, v := range gotten {
		if v != data.data[i*step+2] {
			t.Errorf("GetData with step filter = 3 != every third stored data point")
			t.Errorf("(Index, Value): GetData(%d, %+v) || StoredData(%d, %+v)", i, v, i*3+2, data.data[i*step+2])
			break
		}
	}
}
