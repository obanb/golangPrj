package dto

import (
	"awesomeProject/common"
	errs "awesomeProject/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"io/ioutil"
)

type SignUpUserRequest struct {
	Name          string          `bson:"name" validate:"required,gte=2,lte=50"`
	Surname       string          `bson:"surname" validate:"required,gte=2,lte=50"`
	Nickname      string          `bson:"nickname" validate:"required,gte=4,lte=50"`
	Email         string          `bson:"email" validate:"required,email,gte=4,lte=50"`
	Position      common.Position `bson:"position" validate:"required,gte=4,lte=50"`
	Password      string          `bson:"password" json:"-" validate:"required,gte=12,lte=18"`
	PasswordAgain string          `bson:"password_again" json:"-"  validate:"required,gte=12,lte=18"`
}

func (r SignUpUserRequest) ComparePasswords() *errs.AppError {
	if r.Password != r.PasswordAgain {
		return errs.NewUnexpectedError("Passwords must be the same.")
	}
	return nil
}

func (r SignUpUserRequest) Validate() *[]string {
	validate := validator.New()

	var res []string

	err := validate.Struct(r)
	for _, e := range err.(validator.ValidationErrors) {
		res = append(res, fmt.Sprintf("Validation error, property: %s, validator: %s.", e.StructField(), e.ActualTag()))
	}
	if err != nil {
		return &res
	}
	return nil
}

func RawBodyUnmarshal(raw *io.ReadCloser) (*SignUpUserRequest, *errs.AppError) {
	// reading body stream convert it to bytes
	body, readError := ioutil.ReadAll(*raw)

	if readError != nil {
		return nil, errs.NewUnexpectedError("Reader read error: " + readError.Error())
	}
	var req SignUpUserRequest

	unmarshalErr := bson.UnmarshalExtJSON(body, true, &req)
	if unmarshalErr != nil {
		return nil, errs.NewUnexpectedError("prepareBsonQuery error: " + unmarshalErr.Error())
	}

	return &req, nil
}

//func (raw *io.ReadCloser) UnmarshalRaw(b []byte) error {
//	var j string
//	err := json.Unmarshal(b, &j)
//	if err != nil {
//		return err
//	}
//	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
//	*s = toID[j]
//	return nil
//}
