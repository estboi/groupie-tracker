# GROUPIE-TRACKER_GEOLOCALIZATION

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.  
The site has more appealing and userfriendly UI.  
Now it has filter option. The options are: 
- Members Number
- Creation Date
- First Album Date
- Concerts Location

Now it has search option. The options are: 
- Band/Artist name
- Members 
- Creation Date
- First Album Date
- Concerts Location

And the last that we have added is Geolocolization! 
The artist/band concerts locations displayed on the MAP!

## Usage

For Audit questions see: "https://github.com/01-edu/public/blob/master/subjects/groupie-tracker/search-bar/audit.md"  
### Console functions:

* to run a server  
`go run . `
* to end the server  
`CTRL + C`

## Implementation details  

First, we read the given Api and rewrite values in structures. 

Then with `http.HandleFunc("/", HomeHandler)` by GET method we show the page with list of all groups. By clicking on button `more`, the User will see detailed information about group.

After all we simply run the server by ` http.ListenAndServe` function on certain localhost

`MiniSQL.TXT` was created as solution of many responses to Google Geocode api, which encodes name of location to its coordinates.  

#### Authors

Jefim Klimenkov - jklimenk  
Aleksandr Pavlov - Aleksander  
Elina Otchenko - eotchenk