package models

type User struct {
	UserId   int64  `json:"userid" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type UserSignUp struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserInMysql struct {
	UserId   int64  `json:"userid" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
