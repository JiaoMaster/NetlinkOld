package models

type Address struct {
	Id      int64  `json:"id,omitempty" db:"id"`
	UserId  int64  `json:"userId,omitempty,string" db:"userId"`
	Name    string `json:"name,omitempty" db:"name"`
	Phone   string `json:"phone,omitempty" db:"phone"`
	Address string `json:"address,omitempty" db:"address"`
	Commit  string `json:"commit,omitempty" db:"commit"`
}

type AddressList struct {
	Id      int64  `json:"id,omitempty" db:"id"`
	UserId  int64  `json:"userId,omitempty,string" db:"userId"`
	Name    string `json:"name,omitempty" db:"name"`
	Phone   string `json:"phone,omitempty" db:"phone"`
	Address string `json:"address,omitempty" db:"address"`
	Commit  string `json:"commit,omitempty" db:"commit"`
}
