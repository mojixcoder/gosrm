### GOSRM
---
**GOSRM** is an OSRM client written in Go. It implements all OSRM 5.x installations.  
If you want to get the most out of this package I highly recommend to read OSRM [docs](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md).

#### Features
---
 - [x] Nearest Service
 - [x] Route Service
 - [x] Table Service
 - [x] Match Service
 - [x] Trip Service
 - [ ] Tile Service

#### Installation
---
Requires Go >= 1.18: `go get github.com/mojixcoder/gosrm`

#### How To Use
---
```go
package main

import (
    "context"
    "fmt"

    "github.com/mojixcoder/gosrm"
)

func main() {
    osrm, err := gosrm.New("http://router.project-osrm.org")
    checkErr(err)
 
    nearestRes, err := gosrm.Nearest(context.Background(), osrm, gosrm.Request{
        Profile:     gosrm.ProfileDriving,
	Coordinates: []gosrm.Coordinate{{13.388860, 52.517037}},
    }, gosrm.WithNumber(3), gosrm.WithBearings([]gosrm.Bearing{{Value: 0, Range: 20}}))
    checkErr(err)

    fmt.Println("### Nearest Response ###")
    fmt.Printf("%#v\n", nearestRes)
    fmt.Println("##########")

    // String type represents the type of geometries returned by OSRM.
    // It can be either string or gosrm.LineString based on geometries option.
    // If you don't specify any geometries the default is polyline and you can use string.
    routeRes, err := gosrm.Route[string](context.Background(), osrm, gosrm.Request{
	Profile:     gosrm.ProfileDriving,
	Coordinates: []gosrm.Coordinate{{13.388860, 52.517037}, {13.397634, 52.529407}, {13.428555, 52.523219}},
    })
    checkErr(err)

    fmt.Println("\n### Route Response ###")
    fmt.Printf("%#v\n", routeRes)
    fmt.Println("##########")

    tableRes, err := gosrm.Table(context.Background(), osrm, gosrm.Request{
	Profile:     gosrm.ProfileDriving,
	Coordinates: []gosrm.Coordinate{{13.388860, 52.517037}, {13.397634, 52.529407}, {13.428555, 52.523219}},
    }, gosrm.WithSources([]uint16{0, 1}), gosrm.WithDestinations([]uint16{2}))
    checkErr(err)

    fmt.Println("\n### Table Response ###")
    fmt.Printf("%#v\n", tableRes)
    fmt.Println("##########")

    // This time we use geojson geometries so geometry type is gosrm.LineString not string.
    matchRes, err := gosrm.Match[gosrm.LineString](context.Background(), osrm, gosrm.Request{
	Profile:     gosrm.ProfileDriving,
	Coordinates: []gosrm.Coordinate{{13.3122, 52.5322}, {13.3065, 52.5283}},
    }, gosrm.WithAnnotations(gosrm.AnnotationsTrue), gosrm.WithGeometries(gosrm.GeometryGeoJSON))
    checkErr(err)

    fmt.Println("\n### Match Response ###")
    fmt.Printf("%#v\n", matchRes)
    fmt.Println("##########")

    tripRes, err := gosrm.Trip[string](context.Background(), osrm, gosrm.Request{
	Profile:     gosrm.ProfileDriving,
	Coordinates: []gosrm.Coordinate{{13.388860, 52.517037}, {13.397634, 52.529407}, {13.428555, 52.523219}, {13.418555, 52.523215}},
    }, gosrm.WithSource(gosrm.SourceFirst), gosrm.WithDestination(gosrm.DestinationLast))
    checkErr(err)

    fmt.Println("\n### Trip Response ###")
    fmt.Printf("%#v\n", tripRes)
    fmt.Println("##########")
}

func checkErr(err error) {
    if err != nil {
	panic(err)
    }
}
```
