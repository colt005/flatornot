package service

import (
	"log/slog"
	"time"

	_ "time/tzdata"

	"github.com/colt005/flatornot/constants"
	"github.com/colt005/flatornot/models"
	"github.com/colt005/flatornot/store"
)

type Service struct {
	store *store.Store
}

func New() (*Service, error) {
	s, err := store.New()
	if err != nil {
		return nil, err
	}
	return &Service{
		store: s,
	}, nil
}

func (s *Service) GetPollData() models.PollData {
	return s.store.GetPollData()
}

func (s *Service) AddVote(v string, sessionID string, metadata map[string]interface{}) error {
	vote, err := models.ParseStringToVoteType(v)
	if err != nil {
		return err
	}
	go s.AddVoteToBacklog(vote, sessionID, metadata)
	return s.store.Incr(vote)
}

func (s *Service) AddVoteToBacklog(vote models.VoteType, sessionID string, metadata map[string]interface{}) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		slog.Error("failed to load location", "err", err)
		return
	}
	now := time.Now().In(loc)

	if len(s.store.GetBacklog()) > constants.BACKLOG_THRESHOLD {
		err := s.store.SyncBacklog()
		if err != nil {
			slog.Error("Failed to sync backlog", "err", err)
		}
	}

	s.store.AddVoteToBacklog(models.VoteRow{
		ID:        nil,
		Vote:      vote,
		Metadata:  nil,
		CreatedAt: &now,
		SessionID: sessionID,
	})
}

func (s *Service) SyncBacklog() error {
	return s.store.SyncBacklog()
}

func (s *Service) RegisterClient(client chan string, sessionID string) {
	s.store.RegisterClient(client, sessionID)
}

func (s *Service) UnRegisterClient(client chan string) {
	s.store.UnRegisterClient(client)
}

func (s *Service) HandleBroadcast() {
	s.store.HandleBroadcast()
}

func (s *Service) BroadcastVotes(html string, clientID string) {
	s.store.BroadcastVotes(html, clientID)
}
