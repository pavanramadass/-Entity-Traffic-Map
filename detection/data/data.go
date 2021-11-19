package data

import (
	"os"

	"encoding/json"
	"io/ioutil"
)

// centroid struct holds location data:
// Timestamp (millisecond precision), and X, Y pixel coordinates
type centroid struct {
	Timestamp int64
	X, Y      int
}

// Data struct stores centroids in a list.
type Data struct {
	stored_data []centroid
}

// ImportData imports data from a file into stored_data.
// If the file doesn't exist, create it.
// Stored_data is cleared before import occurs.
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

	var temp_data []centroid
	json.Unmarshal(imported_data, &temp_data)
	d.stored_data = append(d.stored_data, temp_data...)

	return nil
}

// ExportData exports stored_data into the specified .json
// @requires destination be in the form "[filename].json"
func (d *Data) ExportData(destination string) error {
	var out_data []byte
	out_data, _ = json.Marshal(d.stored_data)

	err := ioutil.WriteFile(destination, out_data, 0600)
	if err != nil {
		return err
	}

	return nil
}

// StoreData stores a new centroid datapoint.
func (d *Data) StoreData(sec int64, x, y int) {
	d.stored_data = append(d.stored_data, centroid{Timestamp: sec, X: x, Y: y})
}