package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"	
 // "encoding/json"
 )

type Pay struct{
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	Source int `json:"source"`//渠道来源:0~wechat,1~支付宝,2~qq钱包...
	Appid string `json:"appid"`//开放平台审核通过的应用APPID
	Mchid string `json:"mchid"`//支付分配的商户号
	Apikey string `json:"apiKey"`//支付平台密钥
	PackageValue string `json:"packageValue"`//Sign=WXPay
	NonceStr string `json:"nonceStr"`//随机字符串，不长于32位
	Sign string `json:"sign"`//签名，详见签名生成算法注意：签名方式一定要与统一下单接口使用的一致
	PrepayId string `json:"prepayId"`//回的支付交易会话ID
	Date int64 `json:"date"`//时间戳
	TradeType string `json:"tradeType"`//交易类型
	Status int `json:"status"`//状态:1~成功,-1失败
	Reason string `json:"reason"`//原由
}

func (this *Pay) TableName() string {
	return TableName("user_pay")
}

func (this *Pay) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}