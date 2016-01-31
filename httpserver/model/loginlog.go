package model
import "time"

//The login log of user
type LoginLog struct {
	LoginLogId int
	UserId int
	Ip string
	LoginDatetime time.Time
}