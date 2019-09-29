package model

import (
	"time"
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
