package mysql

import (
	"NetLinkOld/models"
	"fmt"
	"github.com/pkg/errors"
)

func InsertOrder(order *models.Order) error {
	sqlStr := "insert into `order`(createUserId, payUserId, commodityId, payType, ifApply, ifPay, amount, unapplyReason,number,addressId) VALUES (:createUserId, :payUserId, :commodityId, :payType, :ifApply, :ifPay, :amount, :unapplyReason,:number,:addressId)"
	sqlStr2 := "select userId from UserToOld where oldId like CONCAT('%',?,'%') "
	var id int64
	err := db.Get(&id, sqlStr2, order.CreateUserId)
	if id == 0 {
		return errors.New("亲属寻找失败")
	}
	order.PayUserId = id
	_, err = db.NamedExec(sqlStr, order)
	return err
}

func QueOrderList(page int64, amount int64, id string, ch int) ([]*models.OrderList, error) {
	sqlStr := "select id,createUserId,payUserId,commodityId,payType,ifApply,ifPay,amount,unapplyReason,number,flag from `order` where payUserId = ? order by id limit  ?,?"
	if ch == 0 {
		sqlStr = "select id,createUserId,payUserId,commodityId,payType,ifApply,ifPay,amount,unapplyReason,number,flag from `order` where payUserId = ? and ifApply = 0 order by id limit  ?,?"
	}
	sqlStr2 := "select name,cover from commodity where id = ?"
	Olist := []*models.OrderList{}
	err := db.Select(&Olist, sqlStr, id, (page-1)*amount, amount)
	for i, v := range Olist {
		err = db.Get(Olist[i], sqlStr2, v.CommodityId)
		if err != nil {
			break
		}
	}
	return Olist, err
}

func QueOrderListByOld(page int64, amount int64, id string) ([]*models.OrderList, error) {
	sqlStr := "select id,createUserId,payUserId,commodityId,payType,ifApply,ifPay,amount,unapplyReason,number from `order` where createUserId = ? order by id limit  ?,?"
	sqlStr2 := "select name,cover from commodity where id = ?"
	Olist := []*models.OrderList{}
	err := db.Select(&Olist, sqlStr, id, (page-1)*amount, amount)
	for i, v := range Olist {
		err = db.Get(Olist[i], sqlStr2, v.CommodityId)
		if err != nil {
			break
		}
	}
	return Olist, err
}

func QueOrderDetail(id string) (*models.OrderDetail, error) {
	sqlStr := "select id, createUserId, payUserId, commodityId, payType, ifApply, ifPay, amount, unapplyReason,number,addressId from `order` where id = ?"
	sqlStr2 := "select name,cover from commodity where id = ?"
	ODetail := new(models.OrderDetail)
	err := db.Get(ODetail, sqlStr, id)
	fmt.Println(ODetail)
	err = db.Get(ODetail, sqlStr2, ODetail.CommodityId)
	return ODetail, err
}

func PayOrder(id string, payType int, adId string) error {
	tx, err := db.Beginx()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	if err != nil {
		return err
	}
	var re struct {
		ComId int8 `db:"commodityId"`
		IfPay int8 `json:"ifPay" db:"ifPay"`
	}
	sqlstr1 := "select commodityId,ifPay from `order` where id = ?"
	err = db.Get(&re, sqlstr1, id)
	if re.IfPay != 0 {
		return errors.New("重复付款")
	}
	if err != nil {
		return err
	}
	sqlstr2 := "update `commodity` set stock = stock - 1,sold = sold + 1 where id = ? and stock > 0"
	rs, err := tx.Exec(sqlstr2, re.ComId)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("PayOrder sql err")
	}
	sqlStr3 := "update `order` set ifApply = ?,ifPay = ?,payType = ?,addressId = ? where id = ?"
	_, err = tx.Exec(sqlStr3, 2, 1, payType, adId, id)
	return err
}

func UnapplyOrder(id string, UnapplyReason string) error {
	sqlStr := "update `order` set ifApply = ?,unapplyReason = ? where id = ?"
	_, err := db.Exec(sqlStr, 1, UnapplyReason, id)
	return err
}

func DeleteOrder(id string) {

}
