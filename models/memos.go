package models

import "time"

type MemosParam struct {
	Content      string    `json:"content"`
	Type         int       `json:"type"`
	PredictTime  time.Time `json:"predictTime"`
	FinishedTime time.Time `json:"finishedTime"`
}
