package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/gocolly/colly"
)

// Hash Map of teams and team abbreviations, will probably move this somewhere else later
var TEAMS_MAP = map[string]string{
	"buf": "Buffalo Bills",
	"nwe": "New England Patriots",
	"mia": "Miami Dolphins",
	"nyj": "New York Jets",
	"pit": "Pittsburgh Steelers",
	"cin": "Cincinnati Bengals",
	"bal": "Baltimore Ravens",
	"cle": "Cleveland Browns",
	"clt": "Indianapolis Colts",
	"jax": "Jacksonville Jaguars",
	"htx": "Houston Texans",
	"oti": "Tennessee Titans", //ADDED NAME CHANGE IF OILERS
	"sdg": "Los Angeles Chargers", //ADD NAME CHANGE IF SAN DIEGO
	"den": "Denver Broncos",
	"kan": "Kansas City Chiefs",
	"rai": "Las Vegas Raiders",
	"phi": "Philadelphia Eagles",
	"was": "Washington Commanders",
	"dal": "Dallas Cowboys",
	"nyg": "New York Giants",
	"det": "Detroit Lions",
	"gnb": "Green Bay Packers",
	"min": "Minnesota Vikings",
	"chi": "Chicago Bears",
	"tam": "Tampa Bay Buccaneers",
	"atl": "Atlanta Falcons",
	"car": "Carolina Panthers",
	"nor": "New Orleans Saints",
	"ram": "Los Angeles Rams", // ADD NAME CHANGE IF CLEVELAND OR ST. LOUIS RAMS
	"sea": "Seattle Seahawks",
	"sfo": "San Francisco 49ers",
	"crd": "Arizona Cardinals", //ADD NAME CHANGE IF PHOENIX CARDINALS, ST. LOUIS CARDINALS, CHICAGO CARDINALS, etc

}

func get_team_record(team string, year string) []string {
	var result []string // Slice which will return the values needed to print the output in main

	URL := "https://www.pro-football-reference.com/teams/" + team + "/" + year + ".htm"
	c := colly.NewCollector()
	numeric_year, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("This is not a valid year. Try again please")
	}


	
	c.OnHTML("div#meta p:contains('Record:')", func(e *colly.HTMLElement) {
    	text := e.DOM.Clone().Children().Remove().End().Text()
    	record := strings.TrimSpace(strings.TrimPrefix(text, "Record:"))
    	
		parts := strings.Split(record, ",")
		final_parts := strings.Split(parts[0],"-")
		wins := final_parts[0]
		losses := final_parts[1]
		ties := final_parts[2]
    	
		result = append(result, wins)
		result = append(result, losses)
		result = append(result, ties) // "W-L-T"

		// Handles the changing of the names of teams over time, like the Oilers becoming the Titans
		switch team {
		case "oti":
			if (numeric_year < 1999) && (numeric_year >= 1997) {
				result = append(result, "Tennessee Oilers")
			} else if numeric_year < 1997 {
				result = append(result, "Houston Oilers")
			} else {
				result = append(result, "Tennessee Titans")
			}
		case "lac":
			if (numeric_year == 1960) || (numeric_year >= 2017) {
				result = append(result, "Los Angeles Chargers")
			} else {
				result = append(result, "San Diego Chargers")
			}
		case "rai":
			if numeric_year > 2019 {
				result = append(result, "Las Vegas Raiders")
			} else if (numeric_year < 1995) && (numeric_year > 1981) {
				result = append(result, "Los Angeles Raiders")
			} else {
				result = append(result, "Oakland Raiders")
			}
		case "was":
			if numeric_year > 2021 {
				result = append(result, "Washington Commanders")
			} else if (numeric_year == 2021) || (numeric_year == 2020) {
				result = append(result, "Washington Football Team")
			} else if (numeric_year < 2020) && (numeric_year > 1936) {
				result = append(result, "Washington Redskins")
			} else if (numeric_year < 1937) && (numeric_year > 1932) {
				result = append(result, "Boston Redskins")
			} else {
				result = append(result, "Boston Braves")
			}
		case "ram":
			if (numeric_year > 2016) || (numeric_year > 1945 && numeric_year < 1995) {
				result = append(result, "Los Angeles Rams")
			} else if (numeric_year > 1994) && (numeric_year < 2016) {
				result = append(result, "St. Louis Rams")
			} else {
				result = append(result, "Cleveland Rams")
			}
		case "crd":
			if (numeric_year > 1993) {
				result = append(result, "Arizona Cardinals")
			} else if (numeric_year < 1994) && (numeric_year > 1987) {
				result = append(result, "Phoenix Cardinals")
			} else if (numeric_year < 1988) && (numeric_year > 1959) {
				result = append(result, "St. Louis Cardinals")
			} else if (numeric_year == 1944) {
				result = append(result, "Chi/Pit Cards/Steelers")
			} else {
				result = append(result, "Chicago Cardinals")
			}
		default:
			result = append(result, TEAMS_MAP[team])

		}
})
	c.Visit(URL)
	return result
}