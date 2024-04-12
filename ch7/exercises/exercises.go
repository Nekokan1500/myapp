package exercises

import "sort"

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if score1 > score2 {
		l.Wins[team1] += 1
	}
	if score2 > score1 {
		l.Wins[team2] += 1
	}
}

func (l League) Ranking() []string {
	type pair struct {
		key   string
		value int
	}
	pairs := make([]pair, 0, len(l.Wins))
	for k, v := range l.Wins {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})
	ranking := make([]string, 0, len(l.Wins))
	for _, pair := range pairs {
		ranking = append(ranking, pair.key)
	}
	return ranking
}
