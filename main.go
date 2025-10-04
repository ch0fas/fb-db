package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	// See docs for specific function features and parameters
	switch args[0] {
	case "get_team_record": 
		
		gtr_commands := flag.NewFlagSet("get_team_record", flag.ExitOnError)
		gtr_team := gtr_commands.String("team", "", "Team abbreviation")
		gtr_year := gtr_commands.Int("year", 0, "League year")
		gtr_commands.Parse(os.Args[2:])

		result := get_team_record(*gtr_team, strconv.Itoa(*gtr_year))
		fmt.Printf("=== %d %s === \n", *gtr_year, result[3])
		fmt.Printf("Wins: %s \n", result[0])
		fmt.Printf("Losses: %s \n", result[1])
		fmt.Printf("Ties: %s \n", result[2])
	
	case "get_sb_winner":
		gsw_commands := flag.NewFlagSet("get_sb_winner", flag.ExitOnError)
		gsw_year := gsw_commands.Int("year", 0, "League year")
		gsw_commands.Parse(os.Args[2:])

		sb_winner := get_superbowl_winner(strconv.Itoa(*gsw_year))
		fmt.Printf("== Super Bowl %s == \n", sb_winner[1])
		fmt.Printf("Winner: %s \n", sb_winner[0])
	case "scorigami":
		s_commands := flag.NewFlagSet("scorigami", flag.ExitOnError)
		s_winning_score := s_commands.Int("ws", 0, "Winning score")
		s_losing_score := s_commands.Int("ls", 0, "Losing score")
		s_commands.Parse(os.Args[2:])

		answer := scorigami(*s_winning_score, *s_losing_score)
		fmt.Println(answer)
		

	}
}