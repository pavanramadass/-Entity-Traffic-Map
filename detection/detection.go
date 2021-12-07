package detection

import (
	"entityDetection/detection/data"
	"entityDetection/detection/metadata"
	"fmt"
	"image"
	"time"

	"gocv.io/x/gocv"
)

const MinimumArea = 3000

// Detection captures and detects data from a video device
func Detection(kill <-chan bool, catch chan<- string, bytes chan<- []byte) {
	webcam, err := gocv.OpenVideoCapture(-1)
	if err != nil {
		fmt.Printf("Error %s opening video capture device: 0\n", err)
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
		fmt.Printf("Device closed: 0\n")
		return
	}

	buffer, err := gocv.IMEncode(gocv.PNGFileExt, img)
	buffer.GetBytes()
	if err != nil {
		fmt.Println("Error:", err)
		bytes <- buffer.GetBytes()
	}

	gocv.CvtColor(img, &img, gocv.ColorRGBToGray)

	// Initialize Data instance .
	var data data.Data

	// Create channel for passing new centroid points to data.
	// Add 60 frames buffer in case storing is slower than processing.
	points := make(chan image.Point, 60)

	// Start new go routine to consume centroid points.
	go data.StoreData(points)

	fmt.Printf("Start reading device: 0\n")

	for to_kill := false; !to_kill; {
		select {
		case to_kill := <-kill:
			if to_kill {
				fmt.Println("Kill")
			}
		default:
		}

		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: 0\n")
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
	destination := time.Now().Format(metadata.TimeLayout)
	err = data.ExportData(destination)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data capture closed")
	catch <- destination
}
