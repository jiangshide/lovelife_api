package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	"sanskrit_api/models"
	// "strconv"
)

// Users APi
type AutitChannelController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *AutitChannelController) Get() {
	uid := this.getInt64("uid",-1)
	status := this.getInt("status",-2)
	maps,id,err := models.SqlList(models.CHANNEL_STATUS,[...]interface{}{uid,uid,uid,uid,status,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}else{	
		this.response(maps)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router / [post]
func (this *AutitChannelController) Post() {
	id := this.getInt64("id",-1)
	status := this.getInt("status",-1)
	official:=this.getInt("official",-1)
	reason := this.getString("reason",0,0)
	_,err := models.SqlRaw("UPDATE zd_channel SET status=?,official=?,reason=? where id=?",[...]interface{}{status,official,reason,id})
	if err != nil{
		this.false(-1,err)
	}
	this.response(true)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /type [post]
func (this *AutitChannelController) Type() {
	id := this.getInt64("id",-1)
	uid := this.getInt64("uid",-1)
	channelTypeId := this.getInt("channelTypeId",-1)
	_,err := models.SqlRaw("UPDATE zd_channel SET channel_type_id=? where id=?",[...]interface{}{channelTypeId,id})
	if err != nil{
		this.false(-1,err)
	}
	channleTypeUser := new(models.ChannelTypeUser)
	channleTypeUser.Uid = uid
	channleTypeUser.ChannelId = id
	channleTypeUser.ChannelTypeId = channelTypeId 
	_,err = channleTypeUser.Add()

	if err != nil{
		this.false(-1,err)
	}
	this.response(true)
}




