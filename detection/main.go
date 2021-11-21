package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"image/color"

	"gocv.io/x/gocv"

	"detection/data"
)

const MinimumArea = 3000
const fileName = "data_capture.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmotion-detect [camera ID]")
		return
	}

	var data data.Data

	// parse args
	deviceID := os.Args[1]

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	data.ImportData(fileName)

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

	gocv.CvtColor(img, &img, gocv.ColorRGBToGray)

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			break
		}

		if img.Empty() {
			continue
		}

		gocv.CvtColor(img, &img, gocv.ColorRGBToGray)
		mog2.Apply(img, &imgDelta)

		gocv.Threshold(imgDelta, &imgThresh, 25, 255, gocv.ThresholdBinary)

		contours := gocv.FindContours(imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)

		gocv.CvtColor(img, &img, gocv.ColorGrayToBGR)

		sec := time.Now().UnixMilli()
		for i := 0; i < contours.Size(); i++ {
			area := gocv.ContourArea(contours.At(i))
			if area < MinimumArea {
				continue
			}

			point := gocv.MinAreaRect(contours.At(i)).Center
			data.StoreData(sec, point.X, point.Y)

			gocv.Circle(&img, point, 5, color.RGBA{255, 0, 0, 0}, -1)

			statusColor := color.RGBA{255, 0, 0, 0}
			gocv.DrawContours(&img, contours, i, statusColor, 2)
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
	data.ExportData(fileName)
}

// WriteFile writes to a file and returns false, nil for file, and function error message if condition is true.
// Else, it returns true, file name, and function error message if condition is false.
// The only parameter is the file name.
func WriteFile(fileName) {
	file, err := ioutil.WriteFile(fileName, outData, 0600)
	if err != nil {
		return false, nil, err
	}

	return true, file, err
}

// OpenVideoCapture opens the video capture device and returns false and error message if conditional is true.
// Else, it returns true and error message if conditional is false.
// It takes only one parameter: deviceID which is the ID of the video capture device.
func OpenVideoCapture(deviceID) {
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return false, err
	}
	defer webcam.Close()
	return true, err
}
