package main

import(
  "encoding/json"
  "flag"
  "fmt"
  "log"
  "math/rand"
  "net/http"

  "github.com/gorilla/mux"
)

const(
  US_EAST_LNG = 67.24250
  US_WEST_LNG = 124.28639
  US_NORTH_LAT = 48.94528
  US_SOUTH_LAT = 24.52083
)

func main() {
  flag.Parse()
  fmt.Println("Starting backend...")

  router := mux.NewRouter()
  router.HandleFunc("/coords", GetGeoCoordinates).Methods("GET")
  router.HandleFunc("/loc", GetLocationIDFromCoordinates).Methods("GET")

  // Serve on port 9090
  log.Fatal(http.ListenAndServe(":9000", router))
}

type LatLng struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"lng"`
}

// GET random geo coordinates
func GetGeoCoordinates(w http.ResponseWriter, req *http.Request) {
  lng := rand.Float64() * US_WEST_LNG
  for {
    fmt.Println("a")
    if lng < US_EAST_LNG {
      lng = rand.Float64() * US_WEST_LNG
    } else {
      break
    }
  }

  lat := rand.Float64() * US_NORTH_LAT
  for {
    fmt.Println("b")
    if lat < US_SOUTH_LAT {
      lat = rand.Float64() * US_NORTH_LAT
    } else {
      break
    }
  }

  latlng := LatLng{
    Lat: lat,
    Lng: lng*(-1),
  }

  json.NewEncoder(w).Encode(latlng)
}

type Locations struct {
  Locations []*Location `json:"data"`
}

type Location struct {
  Id int `json:"id"`
}

// GET location-id from coordinates
func GetLocationIDFromCoordinates(w http.ResponseWriter, req *http.Request) {

}

// GET recent photos from location-id
func GetRecentPhotosFromLocationID(w http.ResponseWriter, req *http.Request) {

}

// POST coordinate guess, return score
