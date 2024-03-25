package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage ipGeo <ip>")
		os.Exit(-1)
	}

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	ipGeoApi, err := NewGeoApi(fmt.Sprintf("%s/apikey.txt", exPath))
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
