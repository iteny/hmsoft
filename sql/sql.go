package sql

//用户表
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Status   int    `json:"status`
}

//登录日志
type LoginLog struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Time      int64  `json:"time"`
	Ip        string `json:"ip"`
	Useragent string `json:"useragent"`
	Uid       int    `json:"uid"`
}
