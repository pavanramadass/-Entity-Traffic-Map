// What it does:
//
// This example uses the VideoCapture class to capture frames from a connected webcam,
// and displays the video in a Window class.
//
// How to run:
//
// 		go run ./cmd/capwindow/main.go
//

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"time"

	"gocv.io/x/gocv"
)

type FPS struct {
	start, frames float64
}

func (f *FPS) Fps() float64 {
	return f.frames / (float64(time.Now().Unix()) - f.start)
}

const MinimumArea = 3000

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmotion-detect [camera ID]")
		return
	}
	fps_tracker := FPS{start: float64(time.Now().Unix()), frames: 0}

	// parse args
	deviceID := os.Args[1]

	webcam, err := gocv.OpenVideoCapture(deviceID)
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

	// writer, err := gocv.VideoWriterFile("test.avi", "MJPG", 20, img.Cols(), img.Rows(), true)
	// if err != nil {
	// 	fmt.Printf("error opening video writer device: %v\n", "test.avi")
	// 	return
	// }
	// defer writer.Close()

	status := "Ready"

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}
		gocv.CvtColor(img, &img, gocv.ColorRGBToGray)

		status = "Ready"
		statusColor := color.RGBA{0, 255, 0, 0}

		// first phase of cleaning up image, obtain foreground only
		mog2.Apply(img, &imgDelta)

		// remaining cleanup of the image to use for finding contours.
		// first use threshold
		gocv.Threshold(imgDelta, &imgThresh, 25, 255, gocv.ThresholdBinary)

		// then dilate
		kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
		defer kernel.Close()
		gocv.Dilate(imgThresh, &imgThresh, kernel)

		// now find contours
		contours := gocv.FindContours(imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)

		for i := 0; i < contours.Size(); i++ {
			area := gocv.ContourArea(contours.At(i))
			if area < MinimumArea {
				continue
			}

			status = "Motion detected"
			statusColor = color.RGBA{255, 0, 0, 0}
			gocv.DrawContours(&img, contours, i, statusColor, 2)

			rect := gocv.BoundingRect(contours.At(i))
			gocv.Rectangle(&img, rect, color.RGBA{0, 0, 255, 0}, 2)
		}

		fps_tracker.frames += 1
		fps := fps_tracker.Fps()
		fps_str := fmt.Sprintf("FPS: %.2f", fps)
		gocv.PutText(&img, status, image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, statusColor, 2)
		gocv.PutText(&img, fps_str, image.Pt(10, 40), gocv.FontHersheyPlain, 1.2, color.RGBA{0, 255, 0, 0}, 2)
		// writer.Write(img)
		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
