package user

import "time"

type User struct {
	Email [255]byte
	Date  time.Time
}
