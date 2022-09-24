package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
	"strconv"
)

func AddAddress(uidS string, ad models.Address) error {
	uid, _ := strconv.ParseInt(uidS, 10, 64)
	err := mysql.InsertAddress(uid, ad)
	return err
}

func DelAddress(idS string) error {
	id, _ := strconv.ParseInt(idS, 10, 64)
	return mysql.DelAddress(id)
}

func UpdateAddress(idS string, ad models.Address) error {
	id, _ := strconv.ParseInt(idS, 10, 64)
	return mysql.UpdateAddress(id, ad)
}

func GetAdList(uidS string) ([]*models.AddressList, error) {
	uid, _ := strconv.ParseInt(uidS, 10, 64)
	return mysql.QueAdList(uid)
}

func GetAd(idS string) (*models.Address, error) {
	id, _ := strconv.ParseInt(idS, 10, 64)
	return mysql.QueAddress(id)
}
