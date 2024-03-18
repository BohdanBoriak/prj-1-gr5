package domain

import "time"

type User struct {
	Id       uint64
	NickName string
	Time     time.Duration
}
