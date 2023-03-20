package controllers

import (
	// "github.com/astaxie/beego"	
	// "github.com/astaxie/beego/orm"	
	"sanskrit_api/models"
	"strconv"
)

type ChannelController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /add [post]
func (this *ChannelController) Add() {

	id := this.getInt64("id",0)
	channel := new(models.Channel)
	channel.Id = id
	channel.Name = this.getString("name",2,2)
	err := channel.Query()
	channel.Uid = this.getInt64("uid",-1)
	channel.ChannelNatureId = this.getInt("channelNatureId",-1)
	channel.Latitude = this.getFloat("latitude",0.0)
	channel.Longitude = this.getFloat("longitude",0.0)
	channel.LocationType = this.getString("locationType",0,0)
	channel.Cover = this.getString("cover",0,0)
	channel.Des = this.getString("des",2,2)
	models.Location(channel.Latitude,channel.Longitude,channel.LocationType,channel.Uid)
	if err != nil{	
		channel.Status = 2
		id,err := channel.Add()
		models.CourseAdd(channel.Uid,channel.Id,3,"创建了一个频道:"+channel.Name,channel.Des,channel.Cover,"")
		if err != nil{
			this.false(-1,err)
		}else{
			models.AddScore(channel.Uid,channel.Id,8)//更新活跃度
			this.response(id)
			maps,ids,err := models.SqlList(models.USER_FOLLOWED,[...]interface{}{channel.Uid,channel.Uid,channel.Uid,channel.Uid,channel.Uid,1000,0})
			if err != nil || ids == 0{
				// this.false(-1,"data is null!")	
			}else{
				for _,v := range (*maps){
					if v["createChannelNotice"] == "1"{
						profile := new(models.Profile)
						profile.Id = channel.Uid 
						profile.Query()

						notification := new(models.Notification)
						notification.Id = id
						notification.Uid= channel.Uid
						notification.Action=7
						notification.Name = "来自《"+profile.Nick+"》的频道通知"
						notification.Content=profile.Nick+"创建了一条频道:"+channel.Name+",点击查看TA吧!"
						uidStr := v["id"].(string)	
						if uid, err := strconv.ParseInt(uidStr, 10, 64);err == nil{
							notification.PushAlia(uid)
						}
					}
				}
			}
		}	
	}else{
		id,err := channel.Update()
		models.CourseAdd(channel.Uid,channel.Id,3,"针对频道"+channel.Name+"进行了更新",channel.Des,channel.Cover,"")
		if err != nil{
			this.false(-1,err)		
		}
		this.response(id)
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /nature name int8 [get]
func (this *ChannelController) Nature() {
	uid := this.getInt64("uid",-1)
	maps,id,err := models.SqlList(models.SqlNature(),[...]interface{}{uid})
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
// @router /type name int8 [get]
func (this *ChannelController) Type() {
	uid := this.getInt64("uid",-1)
	if uid == -1{
		maps,id,err := models.SqlList(models.CHANNEL_TYPE,[...]interface{}{})
		if err != nil || id == 0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}else{
		maps,id,err := models.SqlList(models.CHANNEL_TYPE_UID,[...]interface{}{uid})
		if err != nil || id == 0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}
	
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /typeList name int8 [get]
func (this *ChannelController) TypeList() {
	channelType := new(models.ChannelType)
	maps,total := channelType.List(this.pageSize,this.page*this.pageSize)
	if total <= 0{
		this.false(-1,"data is null")
	}
	this.response(maps)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /channelId [get]
func (this *ChannelController) ChannelId() {
	status := this.getInt("status",2)	
	id := this.getInt64("id",-1)
	num := this.getInt("num",1)
	maps,id,err := models.SqlList(models.CHANNEL_ID,[...]interface{}{status,num,id,this.page,this.pageSize})
	if err != nil || id==0{
		this.false(-1,"data is null")
	}
	this.response(maps)	
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /channelType [get]
func (this *ChannelController) ChannelType() {
	id := this.getInt64("id",-1)
	name := this.getString("name",0,0)
	uid := this.getInt64("uid",-1)
	models.AddAt(this.getString("atsJson",0,0),3,uid,id)
	if len(name) > 0 {
		maps,id,err := models.Channels(uid,models.CHANNEL_NAME,[...]interface{}{uid,uid,uid,uid,name,this.page,this.pageSize})
		if err != nil || id==0{
			this.false(-1,"data is null")
		}else{
			this.response(maps)	
		}
	}else if id == 1{
		this.Recommend()
	}else if id == 2{
		maps,id,err := models.Channels(uid,models.CHANNEL_HOT,[...]interface{}{uid,uid,uid,uid,this.page,this.pageSize})
		if err != nil || id==0{
			this.false(-1,"data is null")
		}else{
			this.response(maps)	
		}
	}else{
		maps,id,err := models.Channels(uid,models.CHANNEL_TYPE_ID,[...]interface{}{uid,uid,uid,uid,id,this.page,this.pageSize})
		if err != nil || id==0{
			this.false(-1,"data is null")
		}else{
			this.response(maps)	
		}
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /user [get]
func (this *ChannelController) User() {
	status := this.getInt("status",2)
	uid := this.getInt64("uid",-1)
	fromUid := this.getInt64("fromUid",-1)
	if uid == fromUid{
		if status == -1{
			maps,id,err := models.SqlList(models.CHANNEL_UID_ALL,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id==0{
				this.false(-1,"data is null")
			}		
			this.response(maps)	
		}else{
			maps,id,err := models.SqlList(models.CHANNEL_UID,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id==0{
				this.false(-1,"data is null")
			}		
			this.response(maps)	
		}
	}else{
		if status == -1{
			maps,id,err := models.SqlList(models.CHANNEL_UID_ALL_PUBLIC,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id==0{
				this.false(-1,"data is null")
			}		
			this.response(maps)	
		}else{
			maps,id,err := models.SqlList(models.CHANNEL_UID_PUBLIC,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id==0{
				this.false(-1,"data is null")
			}		
			this.response(maps)	
		}
	}
	
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /search [get]
func (this *ChannelController) Search() {
	key := this.getString("word",0,0)
	uid := this.getInt64("uid",-1)
	source := this.getInt("source",2)
	maps,id,err := models.SqlList(models.CHANNEL_SEARCH,[...]interface{}{uid,uid,uid,uid,key,key,key,this.page,this.pageSize})
	if err != nil || id==0{
		this.false(-1,"data is null")
	}else{
		for _,v := range (*maps){
			id,_ := strconv.ParseInt(v["id"].(string),10,64)
			word := new(models.Word)
			word.ReadOrCreates(uid,id,key,source)
		}
		this.response(maps)	
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /official [get]
func (this *ChannelController) Official() {
	uid := this.getInt64("uid",-1)
	maps,id,err := models.Channels(uid,models.CHANNEL_OFFICIAL,[...]interface{}{uid,uid,uid,uid,this.page,this.pageSize})
	if err != nil || id==0{
		this.false(-1,"data is null")		
	}else{
		this.response(maps)	
	}

	// maps,id,err := models.Channels(models.CHANNEL_OFFICIAL,[...]interface{}{status,id,this.page,this.pageSize})
	// if err != nil || id==0{
	// 	this.false(-1,"data is null")
	// }
	// this.response(maps)	
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /recommend [get]
func (this *ChannelController) Recommend() {
	uid := this.getInt64("uid",-1)
	maps,id,err := models.Channels(uid,models.CHANNEL_RECOMMEND_UID,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
	if err != nil || id==0{
		maps,id,err := models.Channels(uid,models.CHANNEL_RECOMMEND_UNFOLLOW,[...]interface{}{uid,uid,uid,uid,this.page,this.pageSize})
		if err != nil || id==0{
			maps,id,err := models.Channels(uid,models.CHANNEL_RECOMMEND_ALL,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,err)
			}else{
				this.response(maps)
			}
		}else{
			this.response(maps)	
		}
	}else{
		this.response(maps)	
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty	
// @router /follow [get]
func (this *ChannelController) Follow() {
	uid := this.getInt64("uid",-1)
	maps,id,err := models.Channels(uid,models.CHANNEL_FOLLOW,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
	if err != nil || id==0{
		this.false(-1,"data is null")
	}
	this.response(maps)	
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *ChannelController) Get() {
	status := this.getInt("status",-2)
	uid := this.getInt64("uid",-1)
	maps,id,err := models.Channels(uid,models.CHANNEL_STATUS,[...]interface{}{uid,uid,uid,uid,status,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}else{
		this.response(maps)
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /id [get]
func (this *ChannelController) Id() {
	id := this.getInt64("id",-1)
	uid := this.getInt64("uid",-1)
	blogId := this.getInt64("blogId",0)
	fromId := this.getInt64("fromId",-1)
	res,_,err := models.Sql("SELECT C.id,C.uid,C.name,C.cover,C.des,C.latitude,C.longitude,C.official,C.update_time date,P.nick,P.sex,P.age,P.icon,(SELECT url FROM zd_blog WHERE id=?) url,(SELECT COUNT(1) FROM zd_blog WHERE channel_id=C.id) blogNum,P.city,(SELECT status FROM zd_user_follow WHERE uid=C.uid AND from_id=? ) follows,(SELECT reason FROM zd_report WHERE content_id=C.id AND source=2 AND uid=? ) reportr FROM zd_channel C LEFT JOIN zd_user_profile P ON P.id=C.uid WHERE C.id=?",[...]interface{}{blogId,fromId,fromId,id})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}else{

		this.response(res)

		fromProfile := new(models.Profile)
		fromProfile.Id = uid
		fromProfile.Query()

		if fromProfile.BroweChannelNotice == 1{
			profile := new(models.Profile)
			profile.Id = fromId
			profile.Query()

			notification := new(models.Notification)
			notification.Id = id
			notification.Uid= fromId
			notification.Action=6
			notification.Name = "来自《"+profile.Nick+"》的浏览通知"
			notification.Content=profile.Nick+"查看了你的频道,赶快去关注TA吧!"
			notification.PushAlia(uid)
		}
	}
}




