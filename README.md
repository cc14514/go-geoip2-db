# GeoIP2 DB #

Based on [geoip2-golang v1.2.1](https://github.com/oschwald/geoip2-golang)

## Installation ##

```
go get https://github.com/cc14514/go-geoip2-db
```

### gx support

* 0.0.1 : QmaotkvehotV6Yvn8roarKxN96v6cnGmdKrHn3PW6BkiiV

```
gx import QmaotkvehotV6Yvn8roarKxN96v6cnGmdKrHn3PW6BkiiV
```

## Example ##

```go
package main

import (
    "fmt"
    "github.com/cc14514/go-geoip2-db"
    "log"
    "net"
)

func main() {
	db,_ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record,_ := db.City(net.ParseIP("115.35.95.90"))
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
```

## License ##

This is free software, licensed under the ISC license.
