package models

import "time"

func GetFormattedTime() string {
	layout := "2006-01-02 15:04:05"
	return time.Now().Format(layout)
}

func GetDay() string {
	layout := "2006-01-02"
	return time.Now().Format(layout)
}
