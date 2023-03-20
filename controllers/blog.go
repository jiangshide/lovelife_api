package controllers

import (
	// "github.com/astaxie/beego"
	"sanskrit_api/models"
	"errors"
	// "github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
)

type BlogController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /history [get]
func (this *BlogController) HistoryGet() {
	
	uid := this.getInt64("uid",-1)
	source := this.getInt("source",-1)
	maps,ids,err := models.SqlList(models.BLOG_HISTORY,[...]interface{}{uid,uid,uid,uid,uid,uid,source,this.page,this.pageSize})
	if err != nil || ids == 0{
		this.false(-1,err)
	}else{
		this.response(maps)
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /history [post]
func (this *BlogController) HistoryPost() {
	history := new(models.History)
	history.Uid = this.getInt64("uid",-1)
	history.ContentId = this.getInt64("contentId",-1)
	history.Num = this.getInt("num",0)
	history.Source = this.getInt("source",-1)
	if _,id,err := history.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /lrcUpdate [get]
func (this *BlogController) LrcUodate() {
	id := this.getInt64("id",0)
	name := this.getString("name",0,0)
	lrcZh := this.getString("lrcZh",0,0)
	lrcEs := this.getString("lrcEs",0,0)
	if len(lrcZh) > 0 && len(lrcEs) >0{
		if _,err :=models.SqlRaw(models.UPDATE_LRC,[...]interface{}{name,lrcZh,lrcEs,id});err != nil{
			this.false(-1,err)
		}else{
			this.response(id)
		}
	}else if len(lrcZh) > 0 && len(lrcEs) == 0{
		if _,err := models.SqlRaw(models.UPDATE_LRC_ZH,[...]interface{}{name,lrcZh,id});err != nil{
			this.false(-1,err)
		}else{
			this.response(id)
		}
	} else if len(lrcEs) > 0 && len(lrcZh) == 0{	
		if _,err := models.SqlRaw(models.UPDATE_LRC_ES,[...]interface{}{name,lrcEs,id});err != nil{
			this.false(-1,err)
		}else{
			this.response(id)
		}
	}else{
		this.false(-1,"参数错误!")
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /lrc [get]
func (this *BlogController) Lrc() {
	id := this.getInt64("id",0)
	name := this.getString("name",DEFAULT_ISNULL,DEFAULT_MIN_SIZE)
	// author := this.getString("author",0,0)//歌手
	// language := this.getInt("language",-1)//语言
	var lrcCode models.LrcCode
	models.GetLrc(id,name,&lrcCode)
	this.response(lrcCode)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *BlogController) Get() {
	uid := this.getInt64("uid",-1)
	status := this.getInt("status",2)
	maps,id,err := models.Blogs(uid,models.BLOG_STATUS,[...]interface{}{uid,uid,uid,uid,uid,status,this.page,this.pageSize})
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
// @router /Id [get]
func (this *BlogController) Id() {
	id := this.getInt("id",-1)		
	uid := this.getInt64("uid",-1)
	maps,ids,err := models.Blogs(uid,models.BLOG_ID, [...]interface{}{uid,uid,uid,uid,uid,id})
	if err != nil || ids == 0{
		this.false(-1,err)
	}else{
		this.response((*maps)[0])
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /format [get]
func (this *BlogController) Format() {
	uid := this.getInt64("uid",-1)
	title := this.getString("title",0,0)
	format := this.getInt("format",0)	

	if len(title) > 0 && format > 0{
		maps,ids,err := models.Blogs(uid,models.BLOG_NAME_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,title,title,format,this.page,this.pageSize})
		if err != nil || ids == 0{
			this.false(-1,"data is null")	
		}else{
			this.response(maps)
		}
	}else if len(title) >0 && format <= 0{
		maps,ids,err := models.Blogs(uid,models.BLOG_NAME,[...]interface{}{uid,uid,uid,uid,uid,title,title,this.page,this.pageSize})
		if err != nil || ids == 0{
			this.false(-1,"data is null")	
		}else{
			this.response(maps)
		}
	}else{
		maps,ids,err := models.Blogs(uid,models.BLOG_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,format,this.page,this.pageSize})
		if err != nil || ids == 0{
			this.false(-1,"data is null")	
		}else{
			this.response(maps)
		}
	}
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /collection [get]
func (this *BlogController) Collection() {
	uid := this.getInt64("uid",-1)			

	maps,id,err := models.Blogs(uid,models.BLOG_COLLECTION,[...]interface{}{uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
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
// @router /user [get]
func (this *BlogController) User() {
	uid := this.getInt64("uid",-1)
	fromUid := this.getInt64("fromUid",-1)
	status := this.getInt("status",2)
	if uid == fromUid{
		if status == 2{
			maps,id,err := models.Blogs(uid,models.BLOG_UID_PRIVATE,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null")	
			}
			this.response(maps)
		}else{
			maps,id,err := models.Blogs(uid,models.BLOG_STATUS,[...]interface{}{uid,uid,uid,uid,uid,uid,status,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null")	
			}
			this.response(maps)
		}
	}else{
		if status == 2{
			maps,id,err := models.Blogs(uid,models.BLOG_UID,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null")	
			}
			this.response(maps)
		}else{
			maps,id,err := models.Blogs(uid,models.BLOG_STATUS_PUBLIC,[...]interface{}{uid,uid,uid,uid,uid,status,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null")	
			}
			this.response(maps)
		}
	}
	
}


// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /city [get]
func (this *BlogController) City() {
	city := this.getString("city",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	uid := this.getInt64("uid",-1)
	mode := this.getInt("mode",0)

	var sql =models.BLOG_CITY
	if mode == 0{
		sql = models.BLOG_HOMETOWN
	}

	maps,id,err := models.Blogs(uid,sql,[...]interface{}{uid,uid,uid,uid,uid,city,this.page,this.pageSize})
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
// @router /follow [get]
func (this *BlogController) Follow() {
	uid := this.getInt64("uid",-1)

	maps,id,err := models.Blogs(uid,models.BLOG_FOLLOW,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
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
// @router /recommend [get]
func (this *BlogController) Recommend() {
	uid := this.getInt64("uid",-1)
	format := this.getInt("format",-1)

	if format != -1{
		maps,id,err := models.Blogs(uid,models.BLOG_RECOMMEND_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,format,this.page,this.pageSize})
		if err != nil || id == 0{
			maps,id,err = models.Blogs(uid,models.BLOG_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,uid,format,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null")
			}else{
				this.response(maps)
			}
			// _,err :=models.SqlRaw(mofels.BLOG_RECOMMEND_ADD,[...]interface{}{status,reason,id})
		}else{
			this.response(maps)
		}
	}else{
		maps,id,err := models.Blogs(uid,models.BLOG_RECOMMEND,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
		if err != nil || id == 0{
			maps,id,err = models.Blogs(uid,models.BLOG_FIND,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null")
			}else{
				this.response(maps)
			}
			// _,err :=models.SqlRaw(models.BLOG_RECOMMEND_ADD,[...]interface{}{status,reason,id})
		}else{
			this.response(maps)
		}
	}
	
	// models.PushAlias("18311271366")
	// models.RegisterIM("18312271344","7f09446f077a12169c88e416e8117c53","成1宜","http://zd112.oss-cn-beijing.aliyuncs.com/img/JPEG_20190420_201820.jpg")
	
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router /find [get]
func (this *BlogController) Find() {
	// status := this.getInt("status",1)
	uid := this.getInt64("uid",-1)

	maps,id,err := models.Blogs(uid,models.BLOG_FIND,[...]interface{}{uid,uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
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
// @router /praise [get]
func (this *BlogController) Praise() {
	uid := this.getInt64("uid",-1)

	maps,id,err := models.Blogs(uid,models.BLOG_PRAISE,[...]interface{}{uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
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
// @router /browse [get]
func (this *BlogController) Browse() {
	uid := this.getInt64("uid",-1)

	maps,id,err := models.Blogs(uid,models.BLOG_BROWSE,[...]interface{}{uid,uid,uid,uid,uid,uid,this.page,this.pageSize})
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
// @router /channel [get]
func (this *BlogController) Channel() {
	channelId:=this.getInt64("channelId",-1)
	uid := this.getInt64("uid",-1)
	sort:=this.getInt("sort",1)

	sql := models.BLOG_CHANNEL_NEW
	if sort == 2{
		sql = models.BLOG_CHANNEL_HOT
	}

	maps,id,err := models.Blogs(uid,sql,[...]interface{}{uid,uid,uid,uid,uid,uid,channelId,this.page,this.pageSize})
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
// @router /syncStatus [get]
func (this *BlogController) SyncStatus() {
	uid := this.getInt64("uid",0)
	name := this.getString("name",0,0)
	sufix := this.getString("sufix",0,0)

	maps,ids,err := models.Sql(models.BLOG_NAME_SUFIX, [...]interface{}{uid,uid,uid,uid,uid,uid,name,sufix})
	if ids==0 || err != nil{
		this.false(-1,err)
	}else{
		this.response(maps)
	}
} 

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /syncBlog [post]
func (this *BlogController) SyncBlog() {
	blog := new(models.Blog)
	blog.Name = this.getString("name",0,0)
	blog.Sufix = this.getString("sufix",0,0)
	blog.Width = this.getInt("width",0)
	blog.Height = this.getInt("height", 0)
	blog.Duration = this.getInt64("duration",0)
	blog.Bitrate = this.getInt("bitrate",0)
	blog.Size = this.getInt64("size",0)
	blog.Rotate = this.getInt("rotate",0)
	blog.Url = this.getString("url",0,0)
	blog.Cover = this.getString("cover",0,0)
	blog.ChannelId = this.getInt64("channelId",0)
	blog.Title = this.getString("title",0,0)
	blog.Des = this.getString("des",0,0)
	blog.City = this.getString("city",0,0)
	blog.Position = this.getString("position",0,0)
	blog.Address = this.getString("address",0,0)
	blog.Format = this.getInt("format",1)
	blog.Uid = this.getInt64("uid",0)
	blog.Latitude = this.getFloat("latitude",0.0)
	blog.Longitude = this.getFloat("longitude",0.0)
	blog.LocationType = this.getString("locationType",0,0)
	blog.AdCode = this.getString("adCode",0,0)
	blog.Source=5//本地同步
	blog.Status=2

	if _,id,err := blog.ReadOrCreates();err != nil || id==0{
		this.false(-1,err)
	}else{

		this.response(id)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /publish [post]
func (this *BlogController) Publish() {
	blog := new(models.Blog)
	blog.Uid = this.getInt64("uid",-1)
	blog.ChannelId = this.getInt64("channelId",-1)
	blog.Content = this.getString("content",0,0)
	blog.Title = this.getString("title", 0, 0)
	blog.Des = this.getString("des", 0, 0)
	blog.Latitude = this.getFloat("latitude",0.0)
	blog.Longitude = this.getFloat("longitude",0.0)
	blog.LocationType = this.getString("locationType",0,0)
	blog.City = this.getString("city",0,0)
	blog.Status = this.getInt("status",2)
	blog.Position = this.getString("position",0,0)
	blog.Address = this.getString("address",0,0)
	blog.Format = this.getInt("format",0)
	blog.Cover = this.getString("cover",0,0)

	blog.ChannelId = models.InitChannel(blog.ChannelId,blog.Uid,"")

	models.Location(blog.Latitude,blog.Longitude,blog.LocationType,blog.Uid)
	if blog.Uid == -1{
		models.AddMake(this.getString("netInfo",0,0),this.getString("device",0,0),this.getIp(),blog.Uid,0,BLOG,errors.New("uid is -1"))
		this.false(-1,"uid is -1")
	}
	if blog.ChannelId == -1{
		models.AddMake(this.getString("netInfo",0,0),this.getString("device",0,0),this.getIp(),blog.Uid,0,BLOG,errors.New("channelId is -1"))
		this.false(-1,"channelId is -1")
	}
	position := models.AddMake(this.getString("netInfo",0,0),this.getString("device",0,0),this.getIp(),blog.Uid,0,BLOG,nil)
	blog.PositionId= position.Id
	id,err := blog.Add()
	if err != nil{
		models.CourseAdd(blog.Uid,id,4,"发布了一个动态:"+blog.Title+"，好可惜,却失败了！",blog.Des,blog.Url,"")
		this.false(-1,err)
	}else{
		models.AddFile(this.getString("filesJson",0, 0),id,blog.ChannelId)
		models.AddAt(this.getString("atsJson",0,0),3,blog.Uid,id)
		models.AddBlogStyle(this.getString("styleJson",0,0),id)
		models.CourseAdd(blog.Uid,id,4,"发布了一个动态:"+blog.Title,blog.Des,blog.Url,"")
		if blog.Format == 0{
			models.AddScore(blog.Uid,id,6)//更新活跃度
		}else if blog.Format == 1{
			models.AddScore(blog.Uid,id,5)//更新活跃度
		}else if blog.Format == 2{
			models.AddScore(blog.Uid,id,4)//更新活跃度
		}else if blog.Format == 3{
			models.AddScore(blog.Uid,id,7)//更新活跃度
		}

		profile := new(models.Profile)
		profile.Id = blog.Uid
		profile.Query()

		notification := new(models.Notification)
		notification.Id = id
		notification.Uid= blog.Uid 
		notification.Action=2
		notification.Name = "来自《"+profile.Nick+"》的动态通知"
		notification.Content=profile.Nick+"创建了一条动态:"+blog.Title+",点击查看TA吧!"
		if profile.Sex == 1{
			notification.PushTag("men")
		}
		if profile.Sex == 2{
			notification.PushTag("women")
		}
		this.response(id)

		maps,ids,err := models.SqlList(models.USER_FOLLOWED,[...]interface{}{blog.Uid,blog.Uid,blog.Uid,blog.Uid,blog.Uid,1000,0})
			if err != nil || ids == 0{
				// this.false(-1,"data is null!")	
			}else{
				for _,v := range (*maps){
					if v["createBlogNotice"] == "0"{
						profile := new(models.Profile)
						profile.Id = blog.Uid 
						profile.Query()

						if profile.CreateBlogNotice == 1{
							notification := new(models.Notification)
							notification.Id = id
							notification.Uid= blog.Uid
							notification.Action=8
							notification.Name = "来自《"+profile.Nick+"》的动态通知"
							notification.Content=profile.Nick+"创建了一条动态:"+blog.Title+",点击查看TA吧!"
							uidStr := v["id"].(string)	
							if uid, err := strconv.ParseInt(uidStr, 10, 64);err == nil{
								notification.PushAlia(uid)
							}
						}
					}
				}
			}
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /update [post]
func (this *BlogController) Update() {
	id := this.getInt64("id",-1)
	status := this.getInt("status",0)
	reason := this.getString("reason",0,0)
	_,err :=models.SqlRaw("UPDATE zd_blog SET status=?,reason=? WHERE id=?",[...]interface{}{status,reason,id})
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
// @router /remove [post]
func (this *BlogController) Remove() {
	id := this.getInt64("id",-1)
	_,err :=models.SqlRaw(models.BLOG_DELETE,[...]interface{}{id})
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
// @router /praiseAdd [post]
func (this *BlogController) PraiseAdd() {
	fromId := this.getInt64("fromId",-1)
	uid := this.getInt64("uid",-1)
	blogId := this.getInt64("blogId",-1)
	status := this.getInt("status",-1)

	praise := new(models.BlogPraise)
	_,_,err:= praise.ReadOrCreates(uid,blogId,status)
	if err != nil{
		this.false(-1,err)		
	}else{

		praise:= make(map[string]interface{})
		praise["blog_id"]=blogId
		praise["status"]=1
		praiseNum,_ := models.SqlCount("zd_blog_praise",praise)
		this.response(praiseNum)

		fromProfile := new(models.Profile)
		fromProfile.Id = fromId
		fromProfile.Query()
		if status > 0 && fromProfile.PraiseNotice == 1{
			profile := new(models.Profile)
			profile.Id = uid
			profile.Query()

			notification := new(models.Notification)
			notification.Id = blogId
			notification.Uid= uid
			notification.Action=2
			notification.Name = "来自《"+profile.Nick+"》的点赞通知"
			notification.Content=profile.Nick+"给你点赞了,点击查看TA吧!"
			notification.PushAlia(fromId)
		}
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /viewAdd [post]
func (this *BlogController) ViewAdd() {	
	fromId := this.getInt64("fromId",-1)	
	uid := this.getInt64("uid",-1)
	blogId := this.getInt64("blogId",-1)
	num := this.getInt("num",0)
	
	blogView := new(models.BlogView)
	blogView.Uid = uid
	blogView.BlogId = blogId
	blogView.Num = num
	blogView.Nums = 1
	if _,id,err := blogView.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
		fromProfile := new(models.Profile)
		fromProfile.Id = fromId
		fromProfile.Query()

		if fromProfile.BroweBlogNotice == 1{
			profile := new(models.Profile)
			profile.Id = uid
			profile.Query()

			notification := new(models.Notification)	
			notification.Id = blogId
			notification.Uid= uid
			notification.Action=4
			notification.Name = "来自《"+profile.Nick+"》的浏览通知"
			notification.Content=profile.Nick+"查看了你的动态,赶快去关注TA吧!"
			notification.PushAlia(fromId)
		}
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /share [get]
func (this *BlogController) Share() {
	blogId := this.getInt64("blogId",-1)
	status := this.getInt("status",1)
	station := this.getInt("station",0)

	maps,id,err := models.SqlList("SELECT BS.create_time date,UP.id,UP.icon,UP.nick,UP.age,UP.sex,UP.city,UP.online,UP.online_id onlineId,UP.online_name onlineName,UF.id followId,UF.status follows FROM zd_blog_share BS LEFT JOIN zd_user_profile UP ON UP.id=BS.uid LEFT JOIN zd_user_follow UF ON UF.uid=UP.id WHERE BS.blog_id=? AND BS.status=? AND BS.station=? ORDER BY BS.update_time DESC LIMIT ? OFFSET ? ",[...]interface{}{blogId,status,station,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null")	
	}
	this.response(maps)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /shareAdd [post]
func (this *BlogController) ShareAdd() {
	uid := this.getInt64("uid",-1)
	blogId := this.getInt64("blogId",-1)
	status := this.getInt("status",-1)
	fromIds := this.getString("fromIds",0,0)
	station := this.getInt("station",0)
	var count interface{}
	var err error
	var arr []int64
	json.Unmarshal([]byte(fromIds),&arr)
	if len(arr) > 0{
		for _,v := range arr{
			models.SqlRaw(models.BLOG_SHARE,[...]interface{}{uid,blogId,status,v,station})
		}
		maps := make(map[string]interface{})
		maps["blog_id"]=blogId
		maps["status"]=1//成功的状态
		count,err = models.SqlCount("zd_blog_share",maps)
		_,err = models.SqlRaw(models.BLOG_UPDATE_SHARE,[...]interface{}{count,blogId})
	}else{
		count,err = models.ShareAdd(uid,blogId,0,status,station)
		if err != nil{
			this.false(-1,err)
		}
	}
	this.response(count)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /topAdd [post]
func (this *BlogController) TopAdd() {
	top := new(models.BlogTop)
	top.Uid = this.getInt64("uid",-1)
	top.BlogId = this.getInt64("blogId",-1)
	top.Status = this.getInt("status",-1)
	if _,id,err := top.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /collectionAdd [post]
func (this *BlogController) CollectionAdd() {
	collection := new(models.BlogCollection)
	collection.Uid = this.getInt64("uid",-1)
	collection.BlogId = this.getInt64("blogId",-1)
	collection.Status = this.getInt("status",-1)
	if _,id,err := collection.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /commentAdd [post]
func (this *BlogController) CommentAdd() {
	uid := this.getInt64("uid",-1)
	contentId := this.getInt64("contentId",-1)
	content := this.getString("content",0,0)
	at := this.getString("at",0,0)
	topic := this.getString("topic",0,0)
	urls := this.getString("urls",0,0)
	status := this.getInt("status",1)
	reason := this.getString("reason",0,0)
	if _,err := models.SqlRaw(models.COMMENT_ADD,[...]interface{}{uid,contentId,content,at,topic,urls,status,reason});err != nil{
		this.false(-1,err)
	}else{
		total := new(Total)
		comment:= make(map[string]interface{})
		comment["content_id"]=contentId
		comment["status"]=1
		commentNum,_ := models.SqlCount("zd_comment",comment)

		commentUid := make(map[string]interface{})
		commentUid["content_id"]=contentId
		commentUid["status"]=1//成功的状态
		commentUidNum,_ := models.SqlCount("zd_comment_uid",commentUid)
		
		total.Count = commentNum+commentUidNum
		total.Size = commentUidNum
		this.response(total)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /comment [get]
func (this *BlogController) Comment() {
	pageSize := this.getInt("pageSize", 10)
	page := this.getInt("page", 0)

	uid := this.getInt64("uid",-1)
	contentId := this.getInt64("contentId",-1)

	if maps,id,err := models.SqlList(models.COMMENT_STATUS,[...]interface{}{uid,uid,uid,contentId,pageSize,page*pageSize});err != nil || id == 0{
		this.false(-1,"the data is null")
	}else{
		this.response(maps)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /praiseCommentAdd [post]
func (this *BlogController) PraiseCommentAdd() {
	commentPraise := new(models.CommentPraise)
	commentPraise.Id = this.getInt64("id",0)
	commentPraise.Uid = this.getInt64("uid",-1)
	commentPraise.CommentId = this.getInt64("commentId",-1)
	commentPraise.Status = this.getInt("status",-1)

	if _,id,err := commentPraise.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}

type Total struct{
	Size int64 `json:"size"`
	Count int64 `json:"count"`
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /commentUidAdd [post]
func (this *BlogController) CommentUidAdd() {
	commentId := this.getInt64("commentId",-1)
	uid := this.getInt64("uid",-1)
	fromUid := this.getInt64("fromUid",-1)
	contentId := this.getInt64("contentId",-1)
	content := this.getString("content",0,0)
	at := this.getString("at",0,0)
	topic := this.getString("topic",0,0)
	urls := this.getString("urls",0,0)
	reply := this.getInt("reply",0)
	status := this.getInt("status",1)
	reason := this.getString("reason",0,0)

	if _,err := models.SqlRaw(models.COMMENT_UID_ADD,[...]interface{}{commentId,uid,fromUid,contentId,content,at,topic,urls,reply,status,reason});err != nil{
		this.false(-1,err)
	}else{
		total := new(Total)
		comment:= make(map[string]interface{})
		comment["content_id"]=contentId
		comment["status"]=1
		commentNum,_ := models.SqlCount("zd_comment",comment)

		commentUid := make(map[string]interface{})
		commentUid["comment_id"]=commentId
		commentUid["status"]=1//成功的状态
		commentUidNum,_ := models.SqlCount("zd_comment_uid",commentUid)
		
		total.Count = commentNum+commentUidNum
		total.Size = commentUidNum
		this.response(total)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /commentUid [get]
func (this *BlogController) CommentUid() {
	pageSize := this.getInt("pageSize", 10)
	page := this.getInt("page", 0)

	commentId := this.getInt64("commentId",-1)
	uid := this.getInt64("uid",-1)
	if maps,id,err := models.SqlList(models.COMMENT_UID_STATUS,[...]interface{}{uid,uid,uid,commentId,pageSize,page*pageSize});err != nil || id == 0{
		this.false(-1,err)
	}else{
		this.response(maps)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /praiseCommentUidAdd [post]
func (this *BlogController) PraiseCommentUidAdd() {
	commentUidPraise := new(models.CommentUidPraise)
	commentUidPraise.Id = this.getInt64("id",0)
	commentUidPraise.Uid = this.getInt64("uid",-1)
	commentUidPraise.CommentUidId = this.getInt64("commentUidId",-1)
	commentUidPraise.Status = this.getInt("status",-1)
	if _,id,err := commentUidPraise.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}


