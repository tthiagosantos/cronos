package repository

import (
	"cronos/internal/domain/user/entity"
	entity2 "cronos/internal/domain/wfm/entity"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindByEmail(username string) (*entity.User, error)
}

type TikectRepository interface {
	Save(tikects []*entity2.Ticket) error
}
