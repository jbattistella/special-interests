package database

type QC_Prompts struct {
	ID int `gorm:"primary_key"`
	Group string `json:"group"`
	Cause string    `json:"cause"`
}