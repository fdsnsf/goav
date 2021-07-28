package main

import (
	"log"

	"github.com/fdsnsf/goav/avcodec"
	"github.com/fdsnsf/goav/avdevice"
	"github.com/fdsnsf/goav/avfilter"
	"github.com/fdsnsf/goav/avformat"
	"github.com/fdsnsf/goav/avutil"
	"github.com/fdsnsf/goav/swresample"
	"github.com/fdsnsf/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
