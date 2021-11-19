package metadata

import (
	"fmt"
	"image"
	"strconv"

	_ "image/png"
)

type Metadata struct {
	Data_filename string
	base_image image.Image
	Start_time int64
	End_time int64
	image_width int
	image_height int
}

func NewMetadata(base_image image.Image, start_time int64, end_time int64) *Metadata {
	bounds := base_image.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	m := new(Metadata)
	m.Data_filename = strconv.FormatInt(start_time, 10) + ".csv"
	m.base_image = base_image
	m.Start_time = start_time
	m.End_time = end_time
	m.image_width = width
	m.image_height = height

	return m
}


