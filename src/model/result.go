package model

import (
	"encoding/json"
	"fmt"
	"time"

	"Tyrant/src/utils"
)

/*
CREATE TABLE result(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	file_name TEXT NOT NULL,
	before_size INTEGER NOT NULL,
	after_size  INTEGER NOT NULL,
	start_time_unix INTEGER NOT NULL,
	finish_time_unix INTEGER NOT NULL
);
*/
type Result struct {
	ID         int       `xorm:"id",json:"id"`
	FileName   string    `xorm:"file_name",json:"file_name"`
	BeforeSize int       `xorm:"before_size",json:"before_size"`
	AfterSize  int       `xorm:"after_size",json:"after_size"`
	StartTime  time.Time `xorm:"start_time_unix",json:"start_time"`
	FinishTime time.Time `xorm:"finish_time_unix",json:"finish_time"`
}

func (r *Result) MarshalJSON() ([]byte, error) {
	s := struct {
		ID            int    `json:"id"`
		FileName      string `json:"file_name"`
		BeforeSize    string `json:"before_size"`
		AfterSize     string `json:"after_size"`
		CompressRatio string `json:"compress_ratio"`
		StartTime     string `json:"start_time"`
		FinishTime    string `json:"finish_time"`
		Duration      string `json:"duration"`
	}{
		ID:            r.ID,
		FileName:      r.FileName,
		BeforeSize:    utils.HumanSize(uint64(r.BeforeSize)),
		AfterSize:     utils.HumanSize(uint64(r.AfterSize)),
		CompressRatio: fmt.Sprintf("%.2f%%", float64(r.AfterSize)/float64(r.BeforeSize)*100),
		StartTime:     r.StartTime.String(),
		FinishTime:    r.FinishTime.String(),
		Duration:      r.FinishTime.Sub(r.StartTime).String(),
	}
	return json.Marshal(s)
}
