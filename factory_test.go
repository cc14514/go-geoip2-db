package geoip2db

import (
	"net"
	"testing"
)

func TestNewGeoipDbByStatik(t *testing.T) {
	db, _ := NewGeoipDbByStatik()
	defer db.Close()
	record, _ := db.City(net.ParseIP("47.104.232.24"))
	//record, _ := db.City(net.ParseIP("47.100.215.147"))
	t.Logf("Portuguese (BR) city name: %v\n", record.City.Names)
	t.Logf("English subdivision name: %v\n", record.Subdivisions)
	t.Logf("Russian country name: %v\n", record.Country)
	t.Logf("ISO country code: %v\n", record.Country.IsoCode)
	t.Logf("Time zone: %v\n", record.Location.TimeZone)
	t.Logf("Location : %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
