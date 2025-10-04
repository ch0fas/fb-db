// COMMANDS FOR GENERAL PIECES OF DATA, NOT SPECIFIC TO A SINGLE TEAM YEAR OR PLAYER

package main

import (
	"fmt"
	"strconv"
	"strings"

	//"strings"

	"github.com/gocolly/colly"
)

var c = colly.NewCollector()

func toRoman(num int) string { // To create the roman numerals
	var result string
	values := []struct {
		val int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, i := range values {
		if num == 0 {
			break
		}

		count := num / i.val
		result += strings.Repeat(i.symbol, count)
		num -= count * i.val
	}

	return result
}

func get_superbowl_winner(year string) []string { // To define the winner and the SB number
	var result []string
	
	URL := "https://www.pro-football-reference.com/years/" + year + "/"
	numeric_year, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println("This is not a valid year. Try again please")
		fmt.Println(numeric_year)
	}

	c.OnHTML("div#meta p:contains('Super Bowl Champion:')", func(h *colly.HTMLElement) {
		sb_winner := h.ChildText("a")
		sb_number := toRoman(numeric_year - 1965)

		result = append(result, sb_winner)
		result = append(result, sb_number)

	})

	switch  {
	case numeric_year > 2024, numeric_year < 1971:
		result[1] = "N/A"
	} 

	c.Visit(URL)
	return result
}

func scorigami(winner int, loser int) string {
	var result string

	URL := "https://www.pro-football-reference.com/boxscores/game-scores.htm"
	if winner < loser { // The loser can't have more points than the winner, evidently 
		result := "Error: The losing team can't have more points than the winning team :("
		return  result
	}

	score_str := strconv.Itoa(winner) + "-" + strconv.Itoa(loser)

	c.OnHTML("table#games tr", func(h *colly.HTMLElement) {
		score_found := false
		var amount string

		h.ForEach("td", func(_ int, hl *colly.HTMLElement) {
			if strings.TrimSpace(hl.Text) == score_str {
				score_found = true
			}

			if score_found {
				h.ForEach("td", func(_ int, h *colly.HTMLElement) {
					if hl.Attr("data-stat") == "counter" {
						amount = hl.Text
						if amount != "0" || amount != "" {
							result = fmt.Sprintf("No Scorigami. The score of %s has happened %s times in NFL History", score_str, amount)
						} 
					}
				})
			}
		})

	})

	c.Visit(URL)
	c.Wait()

	if result == "" {
		result = fmt.Sprintf("=== SCORIGAMI ALERT === \nAs of today, %s has never happpened before in NFL history", score_str)
	}
	return result
} 