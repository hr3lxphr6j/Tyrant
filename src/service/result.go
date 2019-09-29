package service

import (
	"Tyrant/src/model"
)

type ResultService interface {
	GetAllResult(limit, offset int, desc bool) ([]*model.Result, error)
	Count() (int64, error)
}
