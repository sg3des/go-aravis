package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"time"

	aravis "github.com/sg3des/go-aravis"
)

func init() {
	log.SetFlags(log.Lshortfile)

	flag.Float64Var(&exposureTime, "e", 10000, "Exposure time (in us)")
	flag.Float64Var(&gain, "g", 0, "Gain (in dB)")
}

func main() {
	var err error
	var numDevices uint

	flag.Parse()

	// Get devices
	aravis.UpdateDeviceList()
	if numDevices, err = aravis.GetNumDevices(); err != nil {
		log.Fatal(err)
	}

	// Must find at least one device
	if numDevices == 0 {
		log.Fatal("No devices found. Exiting.")
	}

	log.Println("Devices:", numDevices)

	for i := uint(0); i < numDevices; i++ {
		name, err := aravis.GetDeviceId(i)
		if err != nil {
			log.Println(i, err)
			continue
		}

		camera, err := aravis.NewCamera(name)
		if err != nil {
			log.Println(i, err)
			continue
		}

		// list, err := camera.GetAvailablePixelFormats()
		// if err != nil {
		// 	log.Println(i, err)
		// 	continue
		// }

		pixfmt, err := camera.GetPixelFormat()
		if err != nil {
			log.Println(i, err)
			continue
		}

		defer camera.Close()

		log.Printf("Start web handler /%d.jpg for %s with pixel format: %s", i, name, pixfmt)

		http.Handle(fmt.Sprintf("/%d.jpg", i), serveJPEG(camera))
	}

	log.Println(" => Listen: 0:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

var exposureTime float64
var gain float64

func serveJPEG(camera aravis.Camera) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		camera.SetExposureTimeAuto(aravis.AUTO_CONTINUOUS)
		// camera.SetExposureTime(exposureTime)
		camera.SetGain(gain)
		camera.SetFrameRate(3.75)
		camera.SetAcquisitionMode(aravis.ACQUISITION_MODE_SINGLE_FRAME)

		size, err := camera.GetPayloadSize()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		x, y, width, height, err := camera.GetRegion()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a stream
		stream, err := camera.CreateStream()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stream.Close()

		// Add a buffer
		buffer, err := aravis.NewBuffer(size)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stream.PushBuffer(buffer)

		// Start acquisition
		camera.StartAcquisition()
		defer camera.StopAcquisition()

		buffer, err = stream.TimeoutPopBuffer(time.Second)
		if s, _ := buffer.GetStatus(); s != aravis.BUFFER_STATUS_SUCCESS {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := buffer.GetData()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Received frame size:", len(data))

		// Image is in red-green bayer format
		img := aravis.NewBayerRG(image.Rect(x, y, width, height))
		img.Pix = data

		// Write JPEG to client
		jpeg.Encode(w, img, nil)
	})
}
