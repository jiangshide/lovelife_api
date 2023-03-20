package models

import ("github.com/astaxie/beego/orm"
 "github.com/astaxie/beego"	
 "encoding/json"
 )

type At struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ContentId   int64  `json:"contentId"`//内容来源ID
	Source int `json:"Source"`//内容类型:1~用户,2~频道,3~动态,4~评论
	FromId int64 `json:"fromId"`//来自用户ID
}

var ats []int64

func (this *At) TableName() string {
	return TableName("at")
}

func (this *At) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

func AddAt(atsJson string,source int,uid,contentId int64){
	json.Unmarshal([]byte(atsJson),&ats)
	for _,v := range ats	{
		at := new(At)
		at.FromId = v
		at.Uid = uid
		at.ContentId = contentId
		at.Source = source
		if _,err := at.Add();err != nil{
			beego.Info("add~at~err:",err)
		}
	}
}
