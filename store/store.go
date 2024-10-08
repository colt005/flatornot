package store

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/colt005/flatornot/constants"
	"github.com/colt005/flatornot/models"
	"github.com/colt005/flatornot/repo"
)

type Store struct {
	pollData     models.PollData
	mu           sync.Mutex
	repo         *repo.Repo
	lastSynced   time.Time
	clients      sync.Map
	broadcast    chan models.BroadCastMessage
	backlog      []models.VoteRow
	backlogMutex sync.RWMutex
}

func New() (*Store, error) {

	r, err := repo.New()
	if err != nil {
		return nil, err
	}

	s := &Store{
		pollData: models.PollData{
			FlatVotes:   0,
			RoundVotes:  0,
			TotalVotes:  0,
			RecentPolls: make([]string, 0),
		},
		backlog:   make([]models.VoteRow, 0),
		repo:      r,
		broadcast: make(chan models.BroadCastMessage),
	}

	s.BootStrap()

	return s, nil
}

func (s *Store) SyncBacklog() error {
	// get the lock
	s.backlogMutex.Lock()
	defer s.backlogMutex.Unlock()
	// get the elements
	rows := s.backlog
	//insert to DB
	err := s.repo.CreateVoteRows(rows)
	if err != nil {
		slog.Error("Failed to create rows", "error ", err)
		return err
	}
	//clear list
	s.backlog = nil
	
	return nil
}

func (s *Store) AddVoteToBacklog(v models.VoteRow) {
	s.backlogMutex.Lock()
	defer s.backlogMutex.Unlock()

	s.backlog = append(s.backlog, v)
}

func (s *Store) GetBacklog() []models.VoteRow {
	s.backlogMutex.RLock()
	defer s.backlogMutex.RUnlock()
	res := s.backlog
	return res
}

func (s *Store) GetPollData() models.PollData {
	return s.pollData
}

func (s *Store) Incr(voteType models.VoteType) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	switch voteType {
	case models.Flat:
		s.pollData.FlatVotes++
		s.pollData.RecentPolls = append([]string{constants.FLAT_POLL}, s.pollData.RecentPolls...)
	case models.Round:
		s.pollData.RoundVotes++
		s.pollData.RecentPolls = append([]string{constants.ROUND_POLL}, s.pollData.RecentPolls...)
	}
	s.pollData.TotalVotes++

	if len(s.pollData.RecentPolls) > 5 {
		s.pollData.RecentPolls = s.pollData.RecentPolls[:5]
	}

	return nil
}

func (s *Store) BootStrap() {

	var wg sync.WaitGroup

	wg.Add(2)

	var totalVotes []models.VoteTotalResult
	var recentVotes []models.VoteRow

	var firstErr, secondErr error

	go func() {
		defer wg.Done()
		totalVotes, firstErr = s.repo.GetTotalVotes()
	}()

	go func() {
		defer wg.Done()
		recentVotes, secondErr = s.repo.GetRecent5Votes()
	}()

	if firstErr != nil {
		panic(firstErr)
	}

	if secondErr != nil {
		panic(secondErr)
	}

	wg.Wait()

	for index := range totalVotes {
		if totalVotes[index].Vote == models.Flat {
			s.pollData.FlatVotes = totalVotes[index].Total
		} else if totalVotes[index].Vote == models.Round {
			s.pollData.RoundVotes = totalVotes[index].Total
		}
	}

	for index := range recentVotes {
		if recentVotes[index].Vote == models.Flat {
			s.pollData.RecentPolls = append(s.pollData.RecentPolls, constants.FLAT_POLL)
		} else if recentVotes[index].Vote == models.Round {
			s.pollData.RecentPolls = append(s.pollData.RecentPolls, constants.ROUND_POLL)
		}
	}

	s.pollData.TotalVotes = s.pollData.FlatVotes + s.pollData.RoundVotes

	s.lastSynced = time.Now()

}

func (s *Store) RegisterClient(client chan string, sessionID string) {
	s.clients.Store(client, sessionID)
}

func (s *Store) UnRegisterClient(client chan string) {
	s.clients.Delete(client)
	close(client)
}

func (s *Store) HandleBroadcast() {
	for {
		data := <-s.broadcast
		s.clients.Range(func(key, value any) bool {
			client, ok := key.(chan string)
			sessionID, _ := value.(string)
			if sessionID != data.ClientID && ok {
				select {
				case client <- data.Message:
				default:
					fmt.Println("Channel is blocked or full")
				}
			}
			return true // Continue ranging through all items
		})
	}
}

func (s *Store) BroadcastVotes(html string, clientID string) {
	s.broadcast <- models.BroadCastMessage{
		Message:  html,
		ClientID: clientID,
	}
}
