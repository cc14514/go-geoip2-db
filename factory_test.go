package geoip2db

import (
	"fmt"
	"net"
	"testing"
)

func TestNewGeoipDbByStatik(t *testing.T) {
	db, _ := NewGeoipDbByStatik()
	defer db.Close()
	record, _ := db.City(net.ParseIP("115.35.95.90"))
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
