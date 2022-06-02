package models

import "time"

type ServerLog struct {
	DateAndTime time.Time
	Username    string
	Operation   string
	Size        int
}
