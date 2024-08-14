package service

import (
	"log"

	pb "github.com/dilshodforever/5-oyimtixon/model"
	s "github.com/dilshodforever/5-oyimtixon/storage"
)

type AccountService struct {
	stg s.InitRoot
}

func NewAccountService(stg s.InitRoot) *AccountService {
	return &AccountService{stg: stg}
}

func (s *AccountService) CreateAccount(req pb.Send) error {
	err := s.stg.Notification().CreateAccount(req)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
