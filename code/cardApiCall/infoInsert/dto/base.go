package dto

import "time"

type AbstractSelectResult struct {
	ID              int64     `json:"id"`
	Dataowner       string    `db:"dataowner" json:"dataowner"`
	RegistDate      time.Time `db:"regist_date" json:"registDate"`
	EnableStartDate time.Time `db:"enable_start_date" json:"enableStartDate"`
	EnableEndDate   time.Time `db:"enable_end_date" json:"enableEndDate"`
	Version         int64     `db:"version" json:"version"`
}
