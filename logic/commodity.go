package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
)

func CreateCommodity(com *models.Commodity) error {
	err := mysql.InsertCom(com)
	return err
}

func GetCommodityList(page string, amount string, typeid string) ([]*models.Commodity, error) {
	r, err := mysql.QueComList(page, amount, typeid)
	return r, err
}

func GetCommodityDetail(id int64) (*models.Commodity, error) {
	r, err := mysql.QueComDetail(id)
	return r, err
}
