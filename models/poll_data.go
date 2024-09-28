package models

type PollData struct {
	FlatVotes   int
	RoundVotes  int
	TotalVotes  int
	RecentPolls []string
	LatestPun   string
}
