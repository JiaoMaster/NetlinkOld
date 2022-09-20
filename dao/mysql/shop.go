package mysql

import (
	"NetLinkOld/models"
	"strconv"
)

func InsertShop(shop *models.Shop) error {
	sqlStr := "insert into shop(typeId, name, image, Introduction) values (:typeId,:name,:image,:Introduction)"
	_, err := db.Exec(sqlStr, shop.TypeId, shop.Name, shop.Image, shop.Introduction)
	return err
}

func QueShopList(page string, amount string, typeid string) ([]*models.ShopList, error) {
	sqlStr := "select typeId,name,image from shop where typeId = ? order by id limit ?,?"
	re := []*models.ShopList{}
	p, _ := strconv.ParseInt(page, 10, 64)
	a, _ := strconv.ParseInt(amount, 10, 64)
	err := db.Select(&re, sqlStr, typeid, (p-1)*a, a)
	return re, err
}

func QueShopDetail(id int64) (*models.Shop, error) {
	sqlStr := "select id,typeId,name,image,introduction from shop where id = ?"
	re := new(models.Shop)
	err := db.Get(re, sqlStr, id)
	return re, err
}
