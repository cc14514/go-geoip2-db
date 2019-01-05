// Package geoip2 provides an easy-to-use API for the MaxMind GeoIP2 and
// GeoLite2 databases; this package does not support GeoIP Legacy databases.
//
// The structs provided by this package match the internal structure of
// the data in the MaxMind databases.
//
// See github.com/cc14514/go-geoip2/driver for more advanced used cases.
package geoip2db

import (
	"fmt"
	"github.com/cc14514/go-geoip2"
	_ "github.com/cc14514/go-geoip2-db/static/statik"
	"github.com/rakyll/statik/fs"
	"os"
)

func NewGeoipDbByStatik() (*geoip2.DBReader, error) {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	file, err := statikFS.Open("/geoip/1.mmdb")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	return geoip2.OpenByFile(file)
}
