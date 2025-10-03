// COMMANDS FOR GENERAL PIECES OF DATA, NOT SPECIFIC TO A SINGLE TEAM YEAR OR PLAYER

package main

import (
	"fmt"
	"strconv"
	"strings"

	//"strings"

	"github.com/gocolly/colly"
)

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
	c := colly.NewCollector()
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