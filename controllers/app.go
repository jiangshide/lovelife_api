package controllers

import (
	"sanskrit_api/models"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego"
	"fmt"
)

type AppController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true			"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *AppController) Get() {
	name := this.getString("name", DEFAULT_ISNULL, 1)
	
	// var profile models.profile

	// qb,err := orm.NewQueryBuilder("mysql")
	// qb.Select("id","name","create_time","update_time").From("zd_user_profile")
	// o := orm.NewOrm()
	// o.Raw(qb.String()).QueryRow(&profile)
	// beego.Info("-----profile:",profile)

	maps,id,err := models.SqlList(models.UPDATE_NAME,[...]interface{}{name,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}
	this.response(maps)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [post]
func (this *AppController) Post() {
	uid := this.getInt64("uid",0)
	name := this.getString("name", DEFAULT_ISNULL, 1)
	pkg := this.getString("pkg",DEFAULT_ISNULL,1)
	platform := this.getString("platform", DEFAULT_ISNULL, 1)
	channel := this.getString("channel", DEFAULT_ISNULL, 1)
	version := this.getString("version", DEFAULT_ISNULL, 1)
	code := this.getInt("code",0)
	evn:=this.getInt("env",0)
	duration := this.getInt("duration",0)
	times := this.getInt("times",0)
	des := this.getString("des",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	url := this.getString("url",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	status := this.getInt("status",0)

	ids,err :=models.SqlRaw(models.APP_UPDATE,[...]interface{}{uid,name,pkg,platform,channel,version,code,evn,duration,times,des,url,status})
	if err != nil{
		this.false(-1,err)
	}
	this.response(ids)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /update [get]
func (this *AppController) Update() {
	update := new(models.App)
	update.Name = this.getString("name", DEFAULT_ISNULL, 1)
	update.Platform = this.getString("platform", DEFAULT_ISNULL, 1)
	update.Version = this.getString("version", DEFAULT_ISNULL, 1)
	update.Code = this.getInt("code",0)

	 var maps []orm.Params
	if _, err := update.Query().Filter("name", update.Name).Filter("platform", update.Platform).Filter("version", update.Version).Filter("code",update.Code).Values(&maps, "id","name","des","version","code","duration","times","status", "url"); err != nil {
		this.false(-1, err)
	}
	this.response(maps[0])
}

// @Title Post
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /status [post]
func (this *AppController) Status() {
	status := this.getInt("status",-1000)
	id := this.getInt64("id",-1)

	var err error
	if status == -1000 || id == -1{
		this.false(-1,fmt.Sprintf("status is %d and id is %d",status,id))
	}else if status == -2{
		_,err =models.SqlRaw(models.UPDATE_DEL,[...]interface{}{id})
	}else {
		_,err =models.SqlRaw(models.UPDATE_STATUS,[...]interface{}{status,id})
	}	
	if err != nil{
		this.false(-1,err)
	}
	this.response(id)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true			"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /ads [get]
func (this *AppController) Ads() {
	maps,id,err := models.SqlList("SELECT id,name,cover,url FROM zd_app_ads LIMIT 1 OFFSET 0 ",[...]interface{}{})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}
	this.response((*maps)[0])
}
