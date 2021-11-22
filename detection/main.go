package main

import (
	"detection/data"
	"detection/metadata"
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

const MinimumArea = 3000
const fileName = "data_capture.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmotion-detect [camera ID]")
		return
	}

	// Initialize Data instance and import past data.
	var data data.Data
	data.Import(fileName)

	// Create channel for passing new centroid points to data.
	// Add 60 frames buffer in case storing is slower than processing.
	points := make(chan image.Point, 60)

	// // Start new go routine to consume centroid points.
	go data.StoreData(points)

	// parse args
	deviceID := os.Args[1]

	webcam, err := gocv.OpenVideoCapture("demo.avi")
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Motion Window")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	imgDelta := gocv.NewMat()
	defer imgDelta.Close()

	imgThresh := gocv.NewMat()
	defer imgThresh.Close()

	mog2 := gocv.NewBackgroundSubtractorMOG2()
	defer mog2.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("Device closed: %v\n", deviceID)
		return
	}

	buffer, _ := gocv.IMEncode(gocv.PNGFileExt, img)

	gocv.CvtColor(img, &img, gocv.ColorRGBToGray)

	// // Initialize Data instance and import past data.
	meta := metadata.NewMetadata(buffer.GetBytes(), "2021-11-21_23:59:59")
	meta.AssociatedData = &data

	fmt.Printf("Start reading device: %v\n", deviceID)

	for f := 0; f < 4500; f++ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			break
		}

		if img.Empty() {
			fmt.Println("Empty Frame")
			continue
		}

		gocv.CvtColor(img, &img, gocv.ColorRGBToGray)
		mog2.Apply(img, &imgDelta)

		gocv.Threshold(imgDelta, &imgThresh, 25, 255, gocv.ThresholdBinary)

		contours := gocv.FindContours(imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)

		gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

		for i := 0; i < contours.Size(); i++ {
			area := gocv.ContourArea(contours.At(i))
			if area < MinimumArea {
				continue
			}
			fmt.Println("Generating a point.")
			point := gocv.MinAreaRect(contours.At(i)).Center
			points <- point

			gocv.Circle(&img, point, 5, color.RGBA{255, 0, 0, 0}, -1)

			statusColor := color.RGBA{255, 0, 0, 0}
			gocv.DrawContours(&img, contours, i, statusColor, 2)
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}

	close(points)
	source, _ := meta.Export()
	fmt.Println("Metadata exported to:", source)
}
