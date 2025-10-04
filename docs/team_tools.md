# Team Commands

# get_team_record
This function lets you look up the Win-Loss-Tie record of a certain team for a specific season. <br>
**Flags**

- `team`: The team's abbreviation (see abbreviations [here](glossary.md))
- `year`: The season year (e.g, 2024)

**Example Usage** <br>
```
fbdb get_team_record -team=was -year=2024
```
*Example Output*
```
=== 2024 Washington Commanders ===
Wins: 12
Losses: 5
Ties: 0
```