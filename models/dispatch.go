package models

import "time"

type Dispatch struct {
	Id         int64     `json:"id" db:"id"`
	OrderId    int64     `json:"orderId" db:"orderId"`
	ExpressId  int64     `json:"expressId" db:"expressId"`
	ShopId     int64     `json:"shopId" db:"shopId"`
	Remarks    string    `json:"remarks" db:"remarks"`
	CreateTime time.Time `json:"createTime" db:"createTime"`
}
type DispatchDetail struct {
	Id         int64     `json:"id" db:"id"`
	OrderId    int64     `json:"orderId" db:"orderId"`
	ExpressId  int64     `json:"expressId" db:"expressId"`
	ShopId     int64     `json:"shopId" db:"shopId"`
	Remarks    string    `json:"remarks" db:"remarks"`
	CreateTime time.Time `json:"createTime" db:"createTime"`
	Order      *Order    `json:"order"`
	Address    *Address  `json:"address"`
}
