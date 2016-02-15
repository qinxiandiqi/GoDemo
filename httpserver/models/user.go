package model
import "time"

//The user info.
type User struct {
	UserId int
	UserName string
	Credits int
	Password string
	LastVisit time.Time
	LastIp string
}
