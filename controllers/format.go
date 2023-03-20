package controllers

import (
	// "github.com/astaxie/beego"
	"sanskrit_api/models"
	// "strconv"
)

type FormatController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /jsd [post]
func (this *FormatController) Jsd() {
	// ids,err := models.SqlRaw(models.REPORT_TYPE,[...]interface{}{})
	// if err != nil{
	// 	this.false(-1,err)
	// }
	// this.response(ids)
	
}	

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /addInit [get]
func (this *FormatController) AddInit() {
	var list = [...]string{"图文","语音","视频","文字","Web","VR"}
	for _,v := range list{
		format := new(models.Format)
		format.Name = v
		format.Des = v
		format.Uid = 0
		format.Add()
	}
	this.response(list)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *FormatController) Get() {
	maps,id,err := models.SqlList(models.SqlFormat(""),[...]interface{}{20,0})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}
	var user models.User
	http := new(models.Http)
	http.Url = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN"
	http.Model = &user
	err = http.Get()
	this.response(maps)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /update [post]
func (this *FormatController) Update() {
	des := this.getString("des",0,6)
	id := this.getInt64("id",-1)
	if id == -1{
		this.false(-1,"id is -1")
	}
	_,err := models.SqlRaw(models.FORMAT_UPDATE,[...]interface{}{des,id})
	if err != nil{
		this.false(-1,err)
	}
	this.response(id)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router / [post]
func (this *FormatController) Post() {
	format := new(models.Format)
	format.Name = this.getString("name",0,2)
	format.Des = this.getString("des",0,0)
	format.Uid = this.getInt64("uid",-1)
	if format.Uid == -1{
		this.false(-1,"uid is -1")
	}
	id,err := format.Add()
	if err != nil{
		this.false(-1,err)
	}
	this.response(id)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /word [get]
func (this *FormatController) Word() {
	uid := this.getInt64("uid",0)
	source := this.getInt("source",0)
	if uid > 0{
		maps,id,err := models.SqlList(models.WORD_QUERY,[...]interface{}{uid,source,this.page,this.pageSize})
		if err != nil || id==0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}else{
		maps,id,err := models.SqlList(models.WORD_HISTORY,[...]interface{}{source,this.page,this.pageSize})
		if err != nil || id==0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /banner [get]
func (this *FormatController) Banner() {
	name := this.getString("name",0,0)
	uid := this.getInt64("uid",-1)
	word := new(models.Word)
	if len(name) > 0{
		search := word.Search(name)
		if len(search) > 0{
			this.response(search)
			word := new(models.Word)
			word.ReadOrCreates(uid,0,name,0)
		}else{
			this.false(-1,"data is null")
		}
	}else{
		banner := word.Banner()
		if len(banner) > 0{
			this.response(banner)
		}else{
			this.false(-1,"data is null")
		}
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /del [get]
func (this *FormatController) Del() {
	uid := this.getInt64("uid",-1)
	name := this.getString("name",0,0)
	if len(name) > 0{
		if _,err := models.SqlRaw("DELETE FROM zd_hostory WHERE uid=? AND name=? ",[...]interface{}{uid,name});err != nil{
			this.false(-1,err)
		}else{
			this.response(1)
		}
	}else{
		if _,err := models.SqlRaw("DELETE FROM zd_hostory WHERE uid=? ",[...]interface{}{uid});err != nil{
			this.false(-1,err)
		}else{
			this.response(1)
		}
	}
}

