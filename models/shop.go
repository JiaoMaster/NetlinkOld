package models

// ShopType 商铺类型模型
type ShopType struct {
	Id    int64  `json:"id,omitempty,string" db:"id"`
	Name  string `json:"name,omitempty" db:"name"`
	Image string `json:"image,omitempty" db:"image"`
}

// CommodityType 商品类型模型
type CommodityType struct {
	Id    int64  `json:"id,omitempty,string" db:"id"`
	Name  string `json:"name,omitempty" db:"name"`
	Image string `json:"image,omitempty" db:"image"`
}

// Shop 店铺模型
type Shop struct {
	Id           int64  `json:"id,string" db:"id"`
	TypeId       int64  `json:"typeId,string" db:"typeId"`
	Name         string `json:"name" db:"name"`
	Image        string `json:"image" db:"image"`
	Introduction string `json:"introduction" db:"introduction"`
}

// Commodity 商品模型
type Commodity struct {
	Id           int64  `json:"id,omitempty,string" db:"id"`              // 订单id
	ShopId       int64  `json:"shopId,omitempty,string" db:"shopId"`      // 所属商铺id
	TypeId       int64  `json:"typeId,string" db:"typeId"`                // 商品类型id
	Name         string `json:"name,omitempty" db:"name"`                 // 商品名称
	Cover        string `json:"cover,omitempty" db:"cover"`               // 封面图片
	Introduction string `json:"introduction,omitempty" db:"introduction"` // 商品简介
	Images       string `json:"images,omitempty" db:"images"`             // 商品介绍图片列表
	Stock        int64  `json:"stock,omitempty" db:"stock"`               // 库存
	Sold         int64  `db:"sold" json:"sold,omitempty"`                 //	销量
	Score        int64  `json:"score,omitempty" db:"score"`               // 分数
	Amount       int64  `json:"amount,omitempty" db:"amount"`             // 支付金额
	DetailImage  string `json:"detailImage,omitempty" db:"detailImage"`   // 详细介绍图片
}

// Order 订单模型
type Order struct {
	Id            int64  `json:"id,omitempty,string" db:"id"`                     // 订单id
	CreateUserId  int64  `json:"createUserId,omitempty,string" db:"createUserId"` // 创建老人id
	PayUserId     int64  `json:"payUserId,omitempty,string" db:"payUserId"`       // 支付家属id
	CommodityId   int64  `json:"commodityId,omitempty,string" db:"commodityId"`   // 商品id
	PayType       int8   `json:"payType,omitempty" db:"payType"`                  // 支付类型 1：银行卡 2 ：支付宝 3：微信
	IfApply       int8   `json:"ifApply,omitempty" db:"ifApply"`                  // 是否同意 0：未确定 1：不同意 3：已同意
	IfPay         int8   `json:"ifPay,omitempty" db:"ifPay"`                      // 支付状态 0：未支付 1：已支付
	Number        int64  `json:"number,omitempty" db:"number"`
	Amount        int64  `json:"amount,omitempty" db:"amount"` // 支付金额
	AddressId     int64  `json:"addressId,omitempty" db:"addressId"`
	UnapplyReason string `json:"unapplyReason,omitempty" db:"unapplyReason"`
}

type ShopList struct {
	TypeId int64  `json:"typeId,string" db:"typeId"`
	Name   string `json:"name" db:"name"`
	Image  string `json:"image" db:"image"`
}

type OrderList struct {
	Id            int64  `json:"id,omitempty,string" db:"id"`
	CreateUserId  int64  `json:"createUserId,omitempty,string" db:"createUserId"` // 创建老人id
	PayUserId     int64  `json:"payUserId,omitempty,string" db:"payUserId"`       // 支付家属id
	CommodityId   int64  `json:"commodityId,omitempty,string" db:"commodityId"`   // 商品id
	PayType       int8   `json:"payType" db:"payType"`                            // 支付类型 1：银行卡 2 ：支付宝 3：微信
	IfApply       int8   `json:"ifApply" db:"ifApply"`                            // 是否同意 0：未确定 1：不同意 3：已同意
	IfPay         int8   `json:"ifPay" db:"ifPay"`                                // 支付状态 0：未支付 1：已支付
	Amount        int64  `json:"amount" db:"amount"`                              // 支付金额
	Name          string `json:"name,omitempty" db:"name"`                        // 商品名称
	Cover         string `json:"cover,omitempty" db:"cover"`
	Number        int64  `json:"number" db:"number"`
	UnapplyReason string `json:"unapplyReason,omitempty" db:"unapplyReason"`
}

type OrderDetail struct {
	Id            int64  `json:"id,omitempty,string" db:"id"`                     // 订单id
	CreateUserId  int64  `json:"createUserId,omitempty,string" db:"createUserId"` // 创建老人id
	PayUserId     int64  `json:"payUserId,omitempty,string" db:"payUserId"`       // 支付家属id
	CommodityId   int64  `json:"commodityId,omitempty,string" db:"commodityId"`   // 商品id
	PayType       int8   `json:"payType" db:"payType"`                            // 支付类型 1：银行卡 2 ：支付宝 3：微信
	IfApply       int8   `json:"ifApply" db:"ifApply"`                            // 是否同意 0：未确定 1：不同意 3：已同意
	IfPay         int8   `json:"ifPay" db:"ifPay"`                                // 支付状态 0：未支付 1：已支付
	Amount        int64  `json:"amount" db:"amount"`                              // 支付金额
	Name          string `json:"name,omitempty" db:"name"`                        // 商品名称
	Cover         string `json:"cover,omitempty" db:"cover"`                      // 封面图片
	Number        int64  `json:"number" db:"number"`
	AddressId     int64  `json:"addressId,omitempty" db:"addressId"`
	UnapplyReason string `json:"unapplyReason,omitempty" db:"unapplyReason"`
}
