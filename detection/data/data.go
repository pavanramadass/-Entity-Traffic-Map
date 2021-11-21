package data

import (
	"fmt"
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

	skip := 1
	if len(filters) > 2 {
		skip = int(filters[2])
	}

	step := 1
	fmt.Print("\n", start, end, step)
	for i := 0; i < len(d.data); i = i + step {
		time := d.data[i].Timestamp
		fmt.Print("\n", time >= start, (end == -1 || time <= end))
		if time >= start && (end == -1 || time <= end) {
			filtered_data = append(filtered_data, d.data[i])
			step = skip
		}
	}

	return filtered_data
}
