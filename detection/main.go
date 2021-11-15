package main

import (
	"fmt"
	"os"
	"time"

	"encoding/json"
	"image/color"
	"io/ioutil"

	"gocv.io/x/gocv"
)

const MinimumArea = 3000

type Centroid struct {
	Timestamp int64
	X, Y      int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmotion-detect [camera ID]")
		return
	}

	// parse args
	deviceID := os.Args[1]

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	file, err := os.OpenFile("test.json", os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened test.json")
	defer file.Close()

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
	writer, err := gocv.VideoWriterFile("test.avi", "MJPG", 20, img.Cols(), img.Rows(), false)

	if err != nil {
		fmt.Printf("error opening video writer device: %v\n", "test.avi")
		return
	}
	defer writer.Close()

	var data []Centroid
	var outData []byte
	oldData, _ := ioutil.ReadAll(file)
	json.Unmarshal(oldData, &data)

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
			data = append(data, Centroid{Timestamp: sec, X: point.X, Y: point.Y})
			outData, _ = json.Marshal(data)
			pt, _ := json.Marshal(Centroid{Timestamp: sec, X: point.X, Y: point.Y})
			fmt.Println(string(pt))

			gocv.Circle(&img, point, 5, color.RGBA{255, 0, 0, 0}, -1)

			statusColor := color.RGBA{255, 0, 0, 0}
			gocv.DrawContours(&img, contours, i, statusColor, 2)
		}

		//writer.Write(imgDelta)
		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
	err = ioutil.WriteFile("test.json", outData, 0600)
	if err != nil {
		fmt.Println(err)
	}
}
