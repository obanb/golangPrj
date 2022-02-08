package domain

import (
	"awesomeProject/common"
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty`
	Name          string               `bson:"name"`
	Surname       string               `bson:"surname"`
	Nickname      string               `bson:"nickname"`
	Position      common.Position      `bson:"position"`
	AccountStatus common.AccountStatus `bson:"account_status"`
	AccountType   common.AccountType   `bson:"account_type"`
	CreatedAt     string               `bson:"created_at"`
	Password      string               `bson:"password" json:"-"`
}

func (u User) ToDto() dto.SignUpUserResponse {
	return dto.SignUpUserResponse{
		Name:          u.Name,
		Surname:       u.Surname,
		Nickname:      u.Nickname,
		Position:      u.Position,
		AccountStatus: u.AccountStatus,
	}
}

type UserRepository interface {
	Save(User) (*User, *errs.AppError)
}
