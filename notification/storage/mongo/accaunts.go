package postgres

import (
	ctx "context"
	"log"

	"github.com/dilshodforever/5-oyimtixon/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService struct {
	db *mongo.Database
}

func NewAccountService(db *mongo.Database) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateAccount(req model.Send) error {
	coll := s.db.Collection("notifications")
	id := uuid.NewString()
	_, err := coll.InsertOne(ctx.Background(), bson.M{
		"id":       id,
		"UserId":   req.Userid,
		"message":     req.Message,
	})
	if err != nil {
		log.Printf("Failed to create account: %v", err)
		return err
	}
	return nil
}
