package controllers

import (
	// "github.com/astaxie/beego"
	"sanskrit_api/models"
	// "strconv"
)

type ReportController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /type [get]
func (this *ReportController) Type() {
	maps,id,err := models.SqlList(models.REPORT_TYPE,[...]interface{}{})
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
// @router / [get]
func (this *ReportController) Get() {
	types := this.getInt("type",-1)
	source := this.getInt("source",-1)
	status := this.getInt("status",-1)
	reportSql :="SELECT R.id,R.uid,R.content_id contentId,R.type,R.source,R.status,R.reason,R.update_time date,R.create_time createDate,RT.name,UP.nick,UP.icon FROM zd_report R LEFT JOIN zd_report_type RT ON RT.id=R.type LEFT JOIN zd_user_profile UP ON UP.id=R.uid "
	if types != -1 && source != -1 && status != -1{
		maps,id,err := models.SqlList(reportSql+"WHERE R.type=? AND R.source=? AND R.status=? ORDER BY R.update_time LIMIT ? OFFSET ?",[...]interface{}{types,source,status,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}else if types == -1 && source != -1 && status != -1{
		maps,id,err := models.SqlList(reportSql+"WHERE R.source=? AND R.status=? ORDER BY R.update_time LIMIT ? OFFSET ?",[...]interface{}{source,status,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}else if types != -1 && source != -1 && status == -1{
		maps,id,err := models.SqlList(reportSql+"WHERE R.type=? AND R.source=? ORDER BY R.update_time LIMIT ? OFFSET ?",[...]interface{}{types,source,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}else if types != -1 && source == -1 && status != -1{
		maps,id,err := models.SqlList(reportSql+"WHERE R.type=? AND R.status=? ORDER BY R.update_time LIMIT ? OFFSET ?",[...]interface{}{types,status,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null")
		}
		this.response(maps)
	}else{
		maps,id,err := models.SqlList(reportSql+"WHERE R.source=? ORDER BY R.update_time LIMIT ? OFFSET ?",[...]interface{}{source,this.page,this.pageSize})
		if err != nil || id == 0{
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
// @router /add [post]
func (this *ReportController) Add() {
	uid := this.getInt64("uid",-1)
	contentId := this.getInt64("contentId",-1)
	types := this.getInt("type",0)
	source := this.getInt("source",0)
	status := this.getInt("status",0)
	reason := this.getString("reason",0,0)
	if status == 2{
		if source == 1{
			models.SqlRaw("UPDATE zd_uc_user SET status=-5,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}else if source ==2{
			models.SqlRaw("UPDATE zd_channel SET status=-5,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}else if source == 3{
			models.SqlRaw("UPDATE zd_blog SET status=-5,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}else if source == 4{
			models.SqlRaw("UPDATE zd_blog_comment SET status=-5,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}
	}else{
		if source == 1{
			models.SqlRaw("UPDATE zd_uc_user SET status=1,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}else if source ==2{
			models.SqlRaw("UPDATE zd_channel SET status=2,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}else if source == 3{
			models.SqlRaw("UPDATE zd_blog SET status=2,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}else if source == 4{
			models.SqlRaw("UPDATE zd_blog_comment SET status=1,reason=? WHERE id=?",[...]interface{}{reason,contentId})
		}
	}

	report := new(models.Report)
	_,id,err:= report.ReadOrCreates(uid,contentId,types,source,status,reason)
	if err != nil{
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
// @router /feedbackList [get]
func (this *ReportController) FeedbackList() {
	uid := this.getInt64("uid",-1)
	status := this.getInt("status",-1)
	sqlBase := "SELECT FB.id,FB.uid,FB.content_id contentId,FB.content,FB.url,FB.contact,FB.source,FB.create_time date,UP.icon,UP.nick FROM zd_feedback FB LEFT JOIN zd_user_profile UP ON UP.id=FB.uid WHERE status=? "
	if uid != -1{
		if maps,ids,err := models.SqlList(sqlBase+",uid=? ORDER BY FB.create_time LIMIT ? OFFSET ? ",[...]interface{}{status,uid,this.page,this.pageSize});err != nil || ids == 0{
			this.false(-1,err)
		}else{
			this.response(maps)
		}
	}else{
		if maps,ids,err := models.SqlList(sqlBase+"ORDER BY FB.create_time LIMIT ? OFFSET ? ",[...]interface{}{status,this.page,this.pageSize});err != nil || ids == 0{
			this.false(-1,err)
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
// @router /feedback [post]
func (this *ReportController) Feedback() {
	uid := this.getInt64("uid",-1)
	contentId := this.getInt64("contentId",0)
	content := this.getString("content",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	url := this.getString("url",0,0)
	contact := this.getString("contact",0,0)
	source := this.getInt("source",0)
	if _,err := models.SqlRaw("INSERT INTO zd_feedback(uid,content_id,content,url,contact,source) VALUES(?,?,?,?,?,?)",[...]interface{}{uid,contentId,content,url,contact,source});err != nil{
		this.false(-1,err)
	}else{
		this.response(uid)
	}
}

