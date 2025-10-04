# Overall Commands

# get_sb_winner
This function lets you see the Super Bowl winner for a specific season. <br>
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
This function lets you see if a score would be a scorigami and, if not, how many times it has happened <br>
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