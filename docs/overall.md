# Overall Commands

# get_sb_winner
The Super Bowl winner for a specific season. <br>
**Flags**

- `year`: The season year (e.g, 2024)

**Example Usage** <br>
```
fbdb get_sb_winner -year=2024
```
*Example Output*
```
== Super Bowl LIX ==
Winner: Philadelphia Eagles
```

# Scorigami
Lets you see if a score would be a scorigami and, if not, how many times it has happened <br>
**Flags**

- `ws`: The winning team's score (e.g, 20)
- `ls`: The losing team's score (e.g, 17)

Do note that, naturally, the winning team must always have the same or more points than the losing team.

**Example usage** <br>
```
fbdb scorigami -ws=20 -ls=17
```
*Example Output*
```
No Scorigami. The score of 20-17 has happened 298 times in NFL History
```
Alternatively:
```
fbdb scorigami -ws=70 -ls=70
```
```
=== SCORIGAMI ALERT ===
As of today, 70-70 has never happpened before in NFL history
```

# season_overview
Gives you a season's SB Winner, MVP, OROY, DROY, OPOY, DPOY winners, and Passing, Receiving and Rushing leaders <br>
**Flags**

- `year`: The season year (e.g, 2024)

**Example Usage** <br>
```
fbdb season_overview -year=2024
```
*Example Output*
```
=== 2024 SEASON ===
SB LIX Champions: Philadelphia Eagles
AP MVP: Josh Allen
AP OROY: Jayden Daniels
AP DROY: Jared Verse
AP OPOY: Saquon Barkley
AP DPOY: Patrick Surtain
Passing Leader: Joe Burrow, 4918 Yds
Rushing Leader: Saquon Barkley, 2005 Yds
Receiving Leader: Ja'Marr Chase, 1708 Yds
```