package service

import (
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

func (s *Service) AddVote(v string) error {
	vote, err := models.ParseStringToVoteType(v)
	if err != nil {
		return err
	}
	return s.store.Incr(vote)
}

func (s *Service) RegisterClient(client chan string) {
	s.store.RegisterClient(client)
}

func (s *Service) UnRegisterClient(client chan string) {
	s.store.UnRegisterClient(client)
}

func (s *Service) HandleBroadcast() {
	s.store.HandleBroadcast()
}

func (s *Service) BroadcastVotes(html string) {
	s.store.BroadcastVotes(html)
}
