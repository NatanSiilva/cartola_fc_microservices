package entity

import (
	"strconv"
	"time"
)

type MatchResults struct {
	teamAScore int
	teamBScore int
}

func NewMatchResults(teamAScore, teamBScore int) *MatchResults {
	return &MatchResults{
		teamAScore: teamAScore,
		teamBScore: teamBScore,
	}
}

func (m *MatchResults) GetResult() string {
	return strconv.Itoa(m.teamAScore) + "-" + strconv.Itoa(m.teamBScore)
}

type Match struct {
	ID      string
	TeamA   *Team
	TeamB   *Team
	TeamAID string
	TeamBID string
	Date    time.Time
	Status  string
	Results MatchResults
	Action  []GameAction
}

func NewMatch(id string, teamA *Team, teamB *Team, date time.Time) *Match {
	return &Match{
		ID:      id,
		TeamA:   teamA,
		TeamB:   teamB,
		TeamAID: teamA.ID,
		TeamBID: teamB.ID,
		Date:    date,
	}
}
