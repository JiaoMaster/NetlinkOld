package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
	"strconv"
)

func CreateOrder(order *models.Order) error {
	err := mysql.InsertOrder(order)
	return err
}

func GetOrderList(pageStr string, amountStr string, id string, ch int) ([]*models.OrderList, error) {
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	amonut, _ := strconv.ParseInt(amountStr, 10, 64)
	Olist, err := mysql.QueOrderList(page, amonut, id, ch)
	return Olist, err
}

func GetOrderListByOld(pageStr string, amountStr string, id string) ([]*models.OrderList, error) {
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	amonut, _ := strconv.ParseInt(amountStr, 10, 64)
	Olist, err := mysql.QueOrderListByOld(page, amonut, id)
	return Olist, err
}

func GetOrderDetail(id string) (*models.OrderDetail, error) {
	data, err := mysql.QueOrderDetail(id)
	return data, err
}

func PayOrder(id string, payType int, adId string) error {
	err := mysql.PayOrder(id, payType, adId)
	return err
}

func CancelOrder() {

}

func UnapplyOrder(id string, UnapplyReason string) error {
	err := mysql.UnapplyOrder(id, UnapplyReason)
	return err
}
