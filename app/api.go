package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"unicode"
)

var onoff int

var (
	Client *http.Client
	Api    API
	Api2   API
)

// struct of url`s
type Links struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

type API struct {
	Artists         Artists
	Locations       Locations
	Dates           Dates
	Relation        Relation
	FilterLocations []string
	NoMatches       bool
	SearchList      map[string][]string
}

type Artists []struct {
	Id                int
	Image             string
	Name              string
	Members           []string
	MembersNumber     int
	CreationDate      int
	FirstAlbum        string
	Locations         []string
	Coordinate        []Coordinates
	Dates             []string
	Relations         map[string][]string
	FilterAlbumDate   int
	MatchingLocations []string
}
type Coordinates struct {
	Location   string
	Lattitude  string
	Longtitude string
}

type Locations struct {
	Index []struct {
		Id        int
		Locations []string
	}
}

type Dates struct {
	Index []struct {
		Id    int
		Dates []string
	}
}

type Relation struct {
	Index []struct {
		Id             int
		DatesLocations map[string][]string
	}
}

type Sorting Artists

func (a Sorting) Len() int           { return len(a) }
func (a Sorting) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a Sorting) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Sorting_Dates Artists

func (a Sorting_Dates) Len() int           { return len(a) }
func (a Sorting_Dates) Less(i, j int) bool { return a[i].CreationDate < a[j].CreationDate }
func (a Sorting_Dates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ConstructApi gets and decodes JSON data from API URL-s and adds it Artists []structs
func MakeApi() {
	links := GetLinks()
	var FilterLocationsSlice []string
	GetJSON(links.Artists, &Api.Artists)
	GetJSON(links.Locations, &Api.Locations)
	GetJSON(links.Dates, &Api.Dates)
	GetJSON(links.Relation, &Api.Relation)

	// Iterate over structs to add all data to Artists []structs
	for i := 0; i < len(Api.Artists); i++ {
		Api.Artists[i].Locations = Api.Locations.Index[i].Locations
		Api.Artists[i].Dates = Api.Dates.Index[i].Dates
		Api.Artists[i].Relations = Api.Relation.Index[i].DatesLocations
		Api.Artists[i].MembersNumber = len(Api.Artists[i].Members)

		//Correcting some data to make it look better
		for j := 0; j < len(Api.Artists[i].Locations); j++ {
			Api.Artists[i].Locations[j] = strings.Replace(Api.Artists[i].Locations[j], string(Api.Artists[i].Locations[j][0]), string(unicode.ToUpper(rune(Api.Artists[i].Locations[j][0]))), 1)
			Api.Artists[i].Dates[j] = strings.ReplaceAll(Api.Artists[i].Dates[j], "*", "")
			Api.Artists[i].Dates[j] = strings.ReplaceAll(Api.Artists[i].Dates[j], "-", ".")
			Api.Artists[i].Locations[j] = strings.ReplaceAll(Api.Artists[i].Locations[j], "usa", "USA")
			Api.Artists[i].Locations[j] = strings.ReplaceAll(Api.Artists[i].Locations[j], "uk", "UK")

			for k, val := range Api.Artists[i].Locations[j] {
				if val == '_' {
					Api.Artists[i].Locations[j] = strings.ReplaceAll(Api.Artists[i].Locations[j], "_", " ")
				} else if val == '-' {
					Api.Artists[i].Locations[j] = strings.ReplaceAll(Api.Artists[i].Locations[j], string(Api.Artists[i].Locations[j][k]), " — ")
				}
			}
			Api.Artists[i].Locations[j] = strings.Title(Api.Artists[i].Locations[j])
			Api.Artists[i].MatchingLocations = append(Api.Artists[i].MatchingLocations, Api.Artists[i].Locations[j])
			Api.Artists[i].Locations[j] += ": " + Api.Artists[i].Dates[j]
			// After filtering put them in one of filter options
			FilterLocationsSlice = append(FilterLocationsSlice, Api.Artists[i].Locations[j])

		}
	}
	sort.Sort(Sorting(Api.Artists))
	for i := 0; i < len(Api.Artists)-1; i++ {
		Api.Artists[i].Id = i
	}
	// Functions that need to be executed before server
	FilterLocations(FilterLocationsSlice)
	FilterAlbumDates()
	SearchList()
	// data, _ := os.ReadFile("MiniSQL.txt")
	// if data == nil {
	// 	GetCoords()
	// }
	// WriteCoordsToStruct()
}

// GetLinks gets and decodes JSON data from API URL and adds it to links struct to be used in ConstructApi()
func GetLinks() Links {
	url := "https://groupietrackers.herokuapp.com/api"
	var links Links
	GetJSON(url, &links)
	return links
}

// GetJSON gets and decodes JSON data from target URL and adds it to target interface.
func GetJSON(URL string, target interface{}) {
	resp, err := Client.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&target)
}
