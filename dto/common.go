package dto

import errs "awesomeProject/errors"

type MayFailWithIdResponse struct {
	Success bool
	Error   *errs.AppError
	Id      *string
}
