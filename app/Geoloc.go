package app

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// MAP API KEY

// func GetCoords() {
// 	var lat, lng float64
// 	var AddressName string
// 	var err error
// 	file, _ := os.Create("MiniSQL.txt")
// 	data := make(map[string]map[string][]float64)
// 	for Band := 0; Band < len(Api.Artists); Band++ {
// 		for _, Address := range Api.Artists[Band].MatchingLocations {
// 			AddressName, lat, lng, err = getCoordinates(Address)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}
// 			if _, ok := data[Api.Artists[Band].Name]; !ok {
// 				data[Api.Artists[Band].Name] = make(map[string][]float64)
// 			}
// 			data[Api.Artists[Band].Name][AddressName] = []float64{lat, lng}
// 		}
// 	}
// 	file.WriteString(fmt.Sprintln(data))
// }
// func getCoordinates(address string) (string, float64, float64, error) {
// 	// url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", strings.ReplaceAll(address, " ", "+"), apiKey)
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	return "", 0, 0, err
// 	// }
// 	// defer resp.Body.Close()

// 	var result struct {
// 		Results []struct {
// 			FormattedAddress string `json:"formatted_address"`
// 			Geometry         struct {
// 				Location struct {
// 					Lat float64 `json:"lat"`
// 					Lng float64 `json:"lng"`
// 				} `json:"location"`
// 			} `json:"geometry"`
// 		} `json:"results"`
// 	}
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return "", 0, 0, err
// 	}

// 	if len(result.Results) == 0 {
// 		return "", 0, 0, fmt.Errorf("No results found")
// 	}

// 	location := result.Results[0]
// 	return location.FormattedAddress, location.Geometry.Location.Lat, location.Geometry.Location.Lng, nil
// }
// func WriteCoordsToStruct() {
// 	re := regexp.MustCompile(`(.+):\[(-?\d+\.\d+) (-?\d+\.\d+)`)
// 	Data, _ := os.ReadFile("MiniSQL.txt")
// 	DataSplit := strings.Split(string(Data), "]] ")
// 	for Band := 0; Band < len(Api.Artists); Band++ {
// 		Coord := Api.Artists[Band].Coordinate
// 		LocLngLat := strings.Split(string(DataSplit[Band]), "] ")
// 		for LLL := 0; LLL < len(LocLngLat); LLL++ {
// 			loc := re.FindStringSubmatch(LocLngLat[LLL])[1]
// 			Lat := re.FindStringSubmatch(LocLngLat[LLL])[2]
// 			Long := re.FindStringSubmatch(LocLngLat[LLL])[3]
// 			Loc := regexp.MustCompile(`.+\[`).ReplaceAllString(loc, "")
// 			Coord = append(Coord, Coordinates{
// 				Location:   Loc,
// 				Lattitude:  Lat,
// 				Longtitude: Long,
// 			})
// 		}
// 		Api.Artists[Band].Coordinate = Coord
// 	}
// }

// func LocationHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/location" {
// 		ErrorHandler(w, r, http.StatusNotFound)
// 		return
// 	}
// 	data := Api.Artists[id_int].Coordinate
// 	Data, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON data:", err)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(Data)
// }
