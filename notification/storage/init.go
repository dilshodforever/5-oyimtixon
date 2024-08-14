package storage

import "github.com/dilshodforever/5-oyimtixon/model"

type InitRoot interface {
	Notification() NotificationService
}

type NotificationService interface {
	CreateAccount(req model.Send) error
}
