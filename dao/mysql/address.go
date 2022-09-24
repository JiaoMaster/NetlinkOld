package mysql

import "NetLinkOld/models"

func InsertAddress(uid int64, ad models.Address) error {
	sqlStr := "insert into `address`(userId, name, phone, address, commit) VALUES (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, uid, ad.Name, ad.Phone, ad.Address, ad.Commit)
	return err
}

func UpdateAddress(id int64, ad models.Address) error {
	sqlStr := "update address set name = ?,phone = ?,address = ?,commit = ? where id = ?"
	_, err := db.Exec(sqlStr, ad.Name, ad.Phone, ad.Address, ad.Commit, id)
	return err
}

func DelAddress(id int64) error {
	sqlStr := "delete from address where id = ?"
	_, err := db.Exec(sqlStr, id)
	return err
}

func QueAddress(id int64) (*models.Address, error) {
	sqlStr := "select id, userId, name, phone, address, commit from address where id = ?"
	ad := new(models.Address)
	err := db.Get(ad, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func QueAdList(uid int64) ([]*models.AddressList, error) {
	sqlStr := "select id, userId, name, phone, address, commit from address where userId = ?"
	ad := []*models.AddressList{}
	err := db.Select(&ad, sqlStr, uid)
	if err != nil {
		return nil, err
	}
	return ad, nil
}
