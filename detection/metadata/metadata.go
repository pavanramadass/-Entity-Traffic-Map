// Package metadata handles the creation and importation of metadata associated with a dataset
package metadata

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "image/png"

	"entityDetection/detection/centroid"
	"entityDetection/detection/data"
)

const TimeLayout string = "2006-01-02_15_04"

var loc, _ = time.LoadLocation("EST")

// Metadata stores parsed metadata read from JSON
type Metadata struct {
	StartTime   time.Time `json:"starttime"`
	EndTime     time.Time `json:"endtime"`
	BaseImage   string    `json:"baseimage"`
	DataFile    string    `json:"datafile"`
	HeatmapFile string    `json:"heatmapfile"`

	AssociatedData *data.Data `json:"-"`
}

// NewMetadata creates initial metadata for a dataset
func NewMetadata(base_image []byte, endTime string) *Metadata {
	m := new(Metadata)
	m.StartTime = time.Now()
	m.DataFile = m.StartTime.Format(TimeLayout) + ".json"
	m.BaseImage = base64.StdEncoding.EncodeToString(base_image)
	m.EndTime, _ = time.ParseInLocation(TimeLayout, endTime, loc)
	m.AssociatedData = nil

	return m
}

// ImportMetadata reads a JSON metadata file and parses it
func ImportMetadata(source string) []Metadata {
	var metas []Metadata

	fileInfo, err := os.Stat(source)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	if fileInfo.IsDir() {
		items, err := ioutil.ReadDir(".")
		if err != nil {
			return nil
		}
		for _, item := range items {
			metas = append(metas, ImportMetadata(item.Name())...)
		}
	} else {
		file, err := os.Open(source)
		if err != nil {
			log.Fatal("Error: Failed to open metadata file "+source, err)
			return nil
		}
		defer file.Close()

		imported_data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal("Error: Failed to read metadata file "+source, err)
			return nil
		}

		var meta Metadata
		json.Unmarshal(imported_data, &meta)
		metas = append(metas, meta)
	}
	return metas
}

// Export exports the data and metadata and returns the Datafile's file name
func (m *Metadata) Export() (string, error) {
	var out_data []byte
	out_data, _ = json.Marshal(m)

	fileName := "meta_" + m.DataFile
	err := ioutil.WriteFile(fileName, out_data, 0600)
	if err != nil {
		return "", err
	}

	err = m.AssociatedData.ExportData(m.DataFile)
	if err != nil {
		return "", err
	}

	return m.DataFile, nil
}

// GetData applies filters and returns appropriate data.
// Allows for more flexibility in data filtering than data's GetData.
func (m *Metadata) GetData(filters map[string]string) ([]centroid.Centroid, error) {
	dataFilters := []int64{-1, -1, -1}

	if start_str, ok := filters["start"]; ok {
		start, err := time.ParseInLocation(TimeLayout, start_str, loc)
		if err != nil {
			return nil, err
		}
		dataFilters[0] = int64(start.UnixMilli())
	}

	if end_str, ok := filters["end"]; ok {
		end, err := time.ParseInLocation(TimeLayout, end_str, loc)
		if err != nil {
			return nil, err
		}
		dataFilters[1] = int64(end.UnixMilli())
	}

	if step_str, ok := filters["step"]; ok {
		step, err := time.ParseInLocation(TimeLayout, step_str, loc)
		if err != nil {
			return nil, err
		}
		dataFilters[2] = int64(step.UnixMilli())
	}

	m.AssociatedData = &data.Data{}
	m.AssociatedData.Import(m.DataFile)

	return m.AssociatedData.GetData(dataFilters), nil
}
