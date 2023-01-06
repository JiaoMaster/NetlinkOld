package mysql

import (
	"NetLinkOld/models"
)

func CreateDispatch(Ds *models.Dispatch) error {
	sqlStr := "insert into dispatch(orderId, expressId, shopId,remarks) VALUES (?,?,?,?)"
	_, err := db.Exec(sqlStr, Ds.OrderId, Ds.ExpressId, Ds.ShopId, Ds.Remarks)
	return err
}

func DelDispatch(id int64) error {
	sqlStr := "delete from dispatch where id = ?"
	_, err := db.Exec(sqlStr, id)
	return err
}

func UpdateDispatch(New *models.Dispatch) error {
	sqlStr := "update dispatch set expressId = ?,remarks = ?"
	_, err := db.Exec(sqlStr, New.ExpressId, New.Remarks)
	return err
}

func GetDispatchDetail(id int64) (*models.DispatchDetail, error) {
	DsD := new(models.DispatchDetail)
	sqlStr1 := "select id, orderId, expressId, shopId, remarks from dispatch where id = ?"
	err := db.Get(DsD, sqlStr1, id)
	if err != nil {
		return nil, err
	}
	sqlStr2 := "select id, createUserId, payUserId, commodityId, payType, ifApply, ifPay, amount, unapplyReason, number, addressId from `order` where id = ?"
	err = db.Get(*DsD.Order, sqlStr2, DsD.OrderId)
	if err != nil {
		return nil, err
	}
	sqlStr3 := "select id, userId, name, phone, address, commit from address where  id = ?"
	err = db.Get(*DsD.Address, sqlStr3, DsD.Order.AddressId)
	if err != nil {
		return nil, err
	}
	return DsD, nil
}

func GetDispatchList(ShopId int64, SortType int8, Sortid int64) (*[]*models.Dispatch, error) {
	DL := new([]*models.Dispatch)
	var sqlStr string

	if SortType == 0 {
		sqlStr = "select id, orderId, expressId, shopId, remarks from dispatch where shopId = ? and orderId = ? order by createTime desc "
	} else if SortType == 1 {
		sqlStr = "select id, orderId, expressId, shopId, remarks from dispatch where shopId = ? and expressId = ? order by createTime desc "
	} else if SortType == 3 {
		sqlStr = "select id, orderId, expressId, shopId, remarks from dispatch where shopId = ? and orderId = ? order by createTime desc "
	}
	err := db.Select(DL, sqlStr, ShopId, Sortid)
	if err != nil {
		return nil, err
	}
	return DL, nil
}
