package data

import (
	"image"
	"os"
	"time"

	"encoding/json"
	"io/ioutil"
)

// Centroid struct holds location data:
// Timestamp (millisecond precision), and X, Y pixel coordinates
type Centroid struct {
	Timestamp int64
	X, Y      int
}

// Data struct stores Centroids in a list.
type Data struct {
	data []Centroid
}

// ImportData imports data from a file into data.
// If the file doesn't exist, create it.
// data is cleared before import occurs.
func (d *Data) ImportData(source string) error {
	file, err := os.OpenFile(source, os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	imported_data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var temp_data []Centroid
	json.Unmarshal(imported_data, &temp_data)
	d.data = append(d.data, temp_data...)

	return nil
}

// ExportData exports data into the specified .json
// @requires destination be in the form "[filename].json"
func (d *Data) ExportData(destination string) error {
	var out_data []byte
	out_data, _ = json.Marshal(d.data)

	err := ioutil.WriteFile(destination, out_data, 0600)
	if err != nil {
		return err
	}

	return nil
}

// StoreData stores a single frame's worth of Centroids datapoint.
func (d *Data) StoreData(pts []image.Point) {
	sec := time.Now().UnixMilli()
	for _, pt := range pts {
		d.data = append(d.data, Centroid{Timestamp: sec, X: pt.X, Y: pt.Y})
	}
}

// GetData fetches data according to a set of passed in filters.
// Expressed as a slice of ints, the parameters are as follows:
// index=0 start_time of data to return in Unix milliseconds
// index=1 end_time of data to return in Unix milliseconds
// index=2 step of data to return in integer value
// -1 values mean don't apply filter.
func (d *Data) GetData(filters []int64) []Centroid {
	filtered_data := make([]Centroid, 0, len(d.data))

	start := int64(-1)
	if len(filters) > 0 {
		start = filters[0]
	}

	end := int64(-1)
	if len(filters) > 1 {
		end = filters[1]
	}

	step := int64(1)
	if len(filters) > 2 {
		step = filters[2]
	}
	for i := int64(1); i < int64(len(filters)); i = i + step {
		if d.data[i].Timestamp > start && (end == -1 || d.data[i].Timestamp < end) {
			filtered_data = append(filtered_data, d.data[i])
		}
	}

	return filtered_data
}
