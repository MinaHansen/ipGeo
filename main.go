package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage ipGeo <ip>")
		os.Exit(-1)
	}

	ipGeoApi, err := NewGeoApi("apikey.txt")
	if err != nil {
		log.Fatal(err)
	}

	ip := os.Args[1]
	geoLocation, err := ipGeoApi.LookupIp(ip)
	if err != nil {
		log.Fatal(err)
	}

	print(geoLocation.ToString())
}
