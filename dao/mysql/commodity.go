package mysql

import (
	"NetLinkOld/models"
	"strconv"
)

func InsertCom(com *models.Commodity) error {
	sqlStr := "insert into commodity(shopId, typeId, name, cover, introduction, images, stock, detailImage, sold, score, amount) values (?,?,?,?,?,?,?,?,?,?,?)"
	_, err := db.NamedExec(sqlStr, com)
	return err
}

func QueComList(page string, amount string, typeid string) ([]*models.Commodity, error) {
	sqlStr := "select typeId, name, cover, introduction, sold, score, amount from commodity where typeId = ? order by id limit ?,?"
	re := []*models.Commodity{}
	p, _ := strconv.ParseInt(page, 10, 64)
	a, _ := strconv.ParseInt(amount, 10, 64)
	err := db.Select(&re, sqlStr, typeid, (p-1)*a, a)
	return re, err
}

func QueComDetail(id int64) (*models.Commodity, error) {
	sqlStr := "select id,shopId, typeId, name, cover, introduction, images, stock, detailImage, sold, score, amount from commodity where id = ?"
	re := new(models.Commodity)
	err := db.Get(re, sqlStr, id)
	return re, err
}
