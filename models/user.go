package models

type User struct {
	UserId   int64  `json:"userid,string" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	NickName string `json:"nickName" db:"nickName"`
}

type UserSignUp struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserLocation struct {
	UserId string  `json:"userId,omitempty" db:"user_id"`
	X      float64 `json:"x" db:"x"`
	Y      float64 `json:"y" db:"y"`
}

type UserInMysql struct {
	UserId   int64  `json:"userid" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	NickName string `json:"nickName" db:"nickName"`
	Location string `json:"location" db:"location"`
}

type UserToOld struct {
	UserId int64 `json:"userId,string" db:"userId"`
	OldId  int64 `json:"oldId,string" db:"oldId"`
}
