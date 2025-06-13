package app

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var SearchButton string

func SearchList() {
	SearchList := make(map[string][]string)
	// "artist/band", "members", "locations", "first album date", "creation date"
	var NameValue, MembersValue, LocationsValue, AlbumValue, CreationValue []string
	for Band := 0; Band < len(Api.Artists); Band++ {
		NameValue = append(NameValue, Api.Artists[Band].Name)
		MembersValue = append(MembersValue, Api.Artists[Band].Members...)
		AlbumValue = append(AlbumValue, Api.Artists[Band].FirstAlbum)
		DateToStr := strconv.Itoa(Api.Artists[Band].CreationDate)
		CreationValue = append(CreationValue, DateToStr)
	}
	LocationsValue = append(LocationsValue, Api.FilterLocations...)
	SearchList["artist/band"] = NameValue
	SearchList["members"] = MembersValue
	SearchList["locations"] = LocationsValue
	SearchList["first album date"] = AlbumValue
	SearchList["creation date"] = CreationValue
	Api.SearchList = SearchList
}
func SearchReciever(w http.ResponseWriter, r *http.Request) {
	SearchInput := strings.ToLower(r.FormValue("Search"))
	if strings.Contains(SearchInput, " : ") {
		SearchValue := regexp.MustCompile(`(.+) : ([^-]*$)`).FindStringSubmatch(SearchInput)[1]
		SearchKey := regexp.MustCompile(`(.+) : ([^-]*$)`).FindStringSubmatch(SearchInput)[2]
		for Band := 0; Band < len(Api.Artists); Band++ {
			switch SearchKey {
			case "artist/band":
				if strings.ToLower(Api.Artists[Band].Name) == SearchValue {
					Api2.Artists = append(Api2.Artists, Api.Artists[Band])
				}
			case "members":
				for i := 0; i < Api.Artists[Band].MembersNumber; i++ {
					if strings.ToLower(Api.Artists[Band].Members[i]) == SearchValue {
						Api2.Artists = append(Api2.Artists, Api.Artists[Band])
					}
				}
			case "locations":
				for i := 0; i < len(Api.Artists[Band].Locations); i++ {
					if strings.Contains(strings.ToLower(Api.Artists[Band].Locations[i]), SearchValue) {
						Api2.Artists = append(Api2.Artists, Api.Artists[Band])
					}
				}
			case "first album date":
				if Api.Artists[Band].FirstAlbum == SearchValue {
					Api2.Artists = append(Api2.Artists, Api.Artists[Band])
				}
			case "creation date":
				SearchValueInt, _ := strconv.Atoi(SearchValue)
				if Api.Artists[Band].CreationDate == SearchValueInt {
					Api2.Artists = append(Api2.Artists, Api.Artists[Band])
				}
			}
		}
	} else {
	Next:
		for Band := 0; Band < len(Api.Artists); Band++ {
			if strings.Contains(strings.ToLower(Api.Artists[Band].Name), SearchInput) {
				Api2.Artists = append(Api2.Artists, Api.Artists[Band])
				continue Next
			}
			for i := 0; i < Api.Artists[Band].MembersNumber; i++ {
				if strings.Contains(strings.ToLower(Api.Artists[Band].Members[i]), SearchInput) {
					Api2.Artists = append(Api2.Artists, Api.Artists[Band])
					continue Next
				}
			}
			for i := 0; i < len(Api.Artists[Band].Locations); i++ {
				if strings.Contains(strings.ToLower(Api.Artists[Band].Locations[i]), SearchInput) {
					Api2.Artists = append(Api2.Artists, Api.Artists[Band])
					continue Next
				}
			}
			if strings.Contains(Api.Artists[Band].FirstAlbum, SearchInput) {
				Api2.Artists = append(Api2.Artists, Api.Artists[Band])
				continue Next
			}
			SearchValueInt := strconv.Itoa(Api.Artists[Band].CreationDate)
			if strings.Contains(SearchValueInt, SearchInput) {
				Api2.Artists = append(Api2.Artists, Api.Artists[Band])
				continue Next
			}
		}
	}

}
