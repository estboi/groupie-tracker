package app

import (
	"html/template"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	MemberCheck    bool
	CreationCheck  bool
	AlbumCheck     bool
	LocationCheck  bool
	FilteredButton string
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/filter.html"))
	if r.URL.Path != "/filter" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		if Api2.Artists == nil && (FilteredButton == "" || SearchButton == "") {
			Api2.NoMatches = true
			err := template.Execute(w, Api2)
			if err != nil {
				ErrorHandler(w, r, http.StatusNotFound)
			}
		} else {
			err := template.Execute(w, Api2)
			if err != nil {
				ErrorHandler(w, r, http.StatusNotFound)
			}
		}
	case "POST":
		Id_str := r.FormValue("MoreButton")
		id_int, _ = strconv.Atoi(Id_str)
		http.Redirect(w, r, "/artist", http.StatusSeeOther)
	default:
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
	}
}
func Sort(w http.ResponseWriter, r *http.Request) bool {
	if r.FormValue("Name_Reverse") == "ZA" {
		reversed := sort.Reverse(Sorting(Api.Artists))
		sort.Sort(reversed)
		for i := 0; i < len(Api.Artists)-1; i++ {
			Api.Artists[i].Id = i
		}
		return true
	} else if r.FormValue("Name_Normal") == "AZ" {
		sort.Sort(Sorting(Api.Artists))
		for i := 0; i < len(Api.Artists)-1; i++ {
			Api.Artists[i].Id = i
		}
		return true
	} else if r.FormValue("Date_Reverse") == "NewOld" {
		reversed := sort.Reverse(Sorting_Dates(Api.Artists))
		sort.Sort(reversed)
		for i := 0; i < len(Api.Artists)-1; i++ {
			Api.Artists[i].Id = i
		}
		return true
	} else if r.FormValue("Date_Normal") == "OldNew" {
		sort.Sort(Sorting_Dates(Api.Artists))
		for i := 0; i < len(Api.Artists)-1; i++ {
			Api.Artists[i].Id = i
		}
		return true
	}
	return false
}
func PressedReverse(w http.ResponseWriter, r *http.Request) bool {
	if r.FormValue("Name_Reverse") == "ZA" {
		return true
	}
	if r.FormValue("Date_Reverse") == "NewOld" {
		return true
	}
	return false
}
func PressedNormal(w http.ResponseWriter, r *http.Request) bool {
	if r.FormValue("Name_Normal") == "AZ" {
		return true
	}
	if r.FormValue("Date_Normal") == "OldNew" {
		return true
	}
	return false
}
func FilterLocations(Slice []string) {
	var STR []string
	var Check []string
	var CountryCities []string
	for i := 0; i < len(Slice); i++ {
		str := Slice[i]
		str = regexp.MustCompile(`[:][^a-zA-Z]+`).ReplaceAllString(str, ".")
		STR = strings.Split(str, ".")
		for j := range STR {
			Check = append(Check, STR[j])
		}
	}
	seen := make(map[string]bool)
	for _, element := range Check {
		if element != "" {
			if !seen[element] {
				seen[element] = true
				CountryCities = append(CountryCities, element)
			}
		}
	}
	sort.Strings(CountryCities)
	Api.FilterLocations = CountryCities
}
func FilterAlbumDates() {
	for Band := 0; Band < len(Api.Artists); Band++ {
		str := Api.Artists[Band].FirstAlbum
		str = regexp.MustCompile(`\d{2}-\d{2}-(\d{4})`).FindStringSubmatch(str)[1]
		Date, _ := strconv.Atoi(str)
		Api.Artists[Band].FilterAlbumDate = Date
	}
}

// Filter User input
func Filter(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	// Get the values of the checked checkboxes
	MembersNum := r.Form["checkbox"]
	FromCreationDate := r.FormValue("FromYearCreation")
	ToCreationDate := r.FormValue("ToYearCreation")
	FromAlbumDate := r.FormValue("FromAlbumDate")
	ToAlbumDate := r.FormValue("ToAlbumdate")
	Locations := r.FormValue("Locations")
	FilteredButton = r.FormValue("Submit")
	if FilteredButton == "" {
		return
	}
	// for loop Bands
	for Band := 0; Band < len(Api.Artists); Band++ {
		CheckedMembersNum(MembersNum, Band)
		ChoosedCreationDate(FromCreationDate, ToCreationDate, Band)
		ChoosedAlbumDate(FromAlbumDate, ToAlbumDate, Band)
		ChoosedLocations(Locations, Band)
		if MemberCheck && CreationCheck && AlbumCheck && LocationCheck {
			Api2.Artists = append(Api2.Artists, Api.Artists[Band])
			MemberCheck, CreationCheck, AlbumCheck, LocationCheck = false, false, false, false
		} else {
			MemberCheck, CreationCheck, AlbumCheck, LocationCheck = false, false, false, false
		}
	}
}
func CheckedMembersNum(Num []string, Band int) {
	var ValidNum []int
	if Num == nil {
		MemberCheck = true
		return
	}
	for _, i := range Num {
		num, _ := strconv.Atoi(i)
		ValidNum = append(ValidNum, num)
	}
	for _, i := range ValidNum {
		if Api.Artists[Band].MembersNumber == i {
			MemberCheck = true
			break
		}
	}
}
func ChoosedCreationDate(FromStr string, ToStr string, Band int) {
	if FromStr == "1940" && ToStr == "2023" {
		CreationCheck = true
		return
	}
	From, To := FromToIntReverse(FromStr, ToStr)
	for year := From; year <= To; year++ {
		if Api.Artists[Band].CreationDate == year {
			CreationCheck = true
			break
		}
	}
}
func ChoosedAlbumDate(FromStr string, ToStr string, Band int) {
	if FromStr == "1940" && ToStr == "2023" {
		AlbumCheck = true
		return
	}
	From, To := FromToIntReverse(FromStr, ToStr)
	for year := From; year <= To; year++ {
		if Api.Artists[Band].FilterAlbumDate == year {
			AlbumCheck = true
			break
		}
	}
}
func ChoosedLocations(Location string, Band int) {
	if Location == "" {
		LocationCheck = true
		return
	}
	for i := 0; i < len(Api.Artists[Band].Locations); i++ {
		if strings.Contains(Api.Artists[Band].Locations[i], Location) {
			LocationCheck = true
		}
	}
}
func FromToIntReverse(From string, To string) (int, int) {
	var FromInt, ToInt int
	Temp, _ := strconv.Atoi(From)
	ToInt, _ = strconv.Atoi(To)
	if Temp > ToInt {
		FromInt = ToInt
		ToInt = Temp
	} else {
		FromInt = Temp
	}
	return FromInt, ToInt
}
