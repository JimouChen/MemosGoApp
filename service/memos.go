package service

import (
	"memos/dao/mysql"
	"time"
)

func AddTask(content string,
	taskType int,
	predictTime time.Time,
	finishedTime time.Time) error {
	// write db
	if err := mysql.AddTask(content, taskType, predictTime, finishedTime); err != nil {

	}

	return nil
}
