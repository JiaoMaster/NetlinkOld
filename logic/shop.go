package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
)

func CreateShop(shop *models.Shop) (int64, error) {
	id, err := mysql.InsertShop(shop)
	return id, err
}

func GetShopList(page string, amount string, typeid string) ([]*models.ShopList, error) {
	r, err := mysql.QueShopList(page, amount, typeid)
	return r, err
}

func GetShopDetail(id int64) (*models.Shop, error) {
	s, err := mysql.QueShopDetail(id)
	return s, err
}
