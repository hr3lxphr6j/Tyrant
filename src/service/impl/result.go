package impl

import (
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	"Tyrant/src/model"
	"Tyrant/src/service"
)

type SqliteResultService struct {
	db *xorm.Engine
}

func New(dbPath string) (service.ResultService, error) {
	_, err := os.Stat(dbPath)
	if err != nil {
		return nil, err
	}
	db, err := xorm.NewEngine("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &SqliteResultService{
		db: db,
	}, nil
}

func (s *SqliteResultService) GetAllResult(limit, offset int, desc bool) ([]*model.Result, error) {
	if s == nil || s.db == nil {
		return nil, nil
	}
	var results []*model.Result
	sess := s.db.NewSession()
	if desc {
		sess.Desc("id")
	}
	err := sess.Limit(limit, offset).Find(&results)
	return results, err
}

func (s *SqliteResultService) Count() (int64, error) {
	if s == nil || s.db == nil {
		return 0, nil
	}
	return s.db.Count(new(model.Result))
}
