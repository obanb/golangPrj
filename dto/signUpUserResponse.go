package dto

import (
	"awesomeProject/common"
)

type SignUpUserResponse struct {
	Name          string               `bson:"customer_id"`
	Surname       string               `bson:"opening_date"`
	Nickname      string               `bson:"account_type"`
	Position      common.Position      `bson:"position"`
	AccountStatus common.AccountStatus `bson:"account_status"`
}
