package detection

import (
	"entityDetection/detection/data"
	"entityDetection/detection/metadata"
	"fmt"
	"image"
	"os"
	"time"

	"gocv.io/x/gocv"
)

const MinimumArea = 3000

var Meta metadata.Metadata

func Detection(kill <-chan bool, endTime string) {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmotion-detect [camera ID]")
		return
	}

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

	buffer, err := gocv.IMEncode(gocv.PNGFileExt, img)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	gocv.CvtColor(img, &img, gocv.ColorRGBToGray)

	// Initialize Data instance and import past data.
	Meta := metadata.NewMetadata(buffer.GetBytes(), endTime)

	// Initialize Data instance and import past data.
	var data data.Data
	data.Import(Meta.DataFile)

	// Create channel for passing new centroid points to data.
	// Add 60 frames buffer in case storing is slower than processing.
	points := make(chan image.Point, 60)

	// Start new go routine to consume centroid points.
	go data.StoreData(points)

	Meta.AssociatedData = &data

	fmt.Printf("Start reading device: %v\n", deviceID)

	for Meta.StartTime.Before(time.Now()) && Meta.EndTime.After(time.Now()) {
		_, ok := <-kill
		if !ok {
			fmt.Println("Kill")
			break
		}

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

		for i := 0; i < contours.Size(); i++ {
			area := gocv.ContourArea(contours.At(i))
			if area < MinimumArea {
				continue
			}

			point := gocv.MinAreaRect(contours.At(i)).Center
			points <- point

		}
	}

	close(points)
	source, err := Meta.Export()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Metadata exported to:", source)
}