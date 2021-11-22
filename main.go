package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

type test_struct struct {
	Request_Type string
	Start_Date string
	End_Date string
	Data_Content string
}

var (
	deviceID int
	err      error
	webcam   *gocv.VideoCapture
	stream   *mjpeg.Stream
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET": // THIS SHOULD RETURN AN OBJECT TYPE OF 'Year-Month-DD' dependent on the current schedule
		log.Println("Returning current schedule")
		w.Write([]byte(`{"Request_Type": "get_schedule", "Start_Date": "2021-January-1", "End_Date": "2021-January-1"}`))
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		if (t.Request_Type == "data_schedule") { // THIS SHOULD RETURN AN OBJECT TYPE OF 'Year-Month-DD' dependent on the current schedule
			log.Println("Data scheduling requested")
			res := []byte(`{"Request_Type":"` + t.Request_Type + `", "Start_Date": "` + t.Start_Date + `", "End_Date": "` + t.End_Date + `"}`)
			w.Write(res)
		} else if (t.Request_Type == "edit_schedule") { // THIS SHOULD RETURN AN OBJECT TYPE OF 'Year-Month-DD' dependent on the current schedule
			log.Println("Edit scheduling requested")
			res := []byte(`{"Request_Type":"` + t.Request_Type + `", "Start_Date": "` + t.Start_Date + `", "End_Date": "` + t.End_Date + `"}`)
			w.Write(res)
		} else if (t.Request_Type == "map_generation") { // Data_Content should be updated to send back the path to the img file which has the map
			log.Println("Map generation requested")
			res := []byte(`{"Request_Type":"` + t.Request_Type + `", "Data_Content": "url('res/image/filled.png')"}`)
			w.Write(res)
		} else if (t.Request_Type == "cancel_schedule") { // Nothing really needs to be returned here, just cancel the schedule
			log.Println("Cancelling schedule")
			res := []byte(`{"Request_Type:"` + t.Request_Type + `"}`)
			w.Write(res)
		}
	}
}

func mjpegCapture() {
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf.GetBytes())
		buf.Close()
	}
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/form", handleRequest)

	deviceID := 1
	host := 8080
	webcam, err = gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// create the mjpeg stream
	stream = mjpeg.NewStream()

	// start capturing
	go mjpegCapture()

	fmt.Println("Capturing. Point your browser to " + host)

	// start http server
	http.Handle("/videoStream", stream)
	log.Fatal(http.ListenAndServe(host, nil))

	log.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}