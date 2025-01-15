package main

import (
	"log"

	aravis "github.com/sg3des/go-aravis"
)

func main() {
	aravis.UpdateDeviceList()

	n, err := aravis.GetNumDevices()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Devices:", n)

	for i := uint(0); i < n; i++ {
		log.Println(aravis.GetDeviceId(i))
	}
}
