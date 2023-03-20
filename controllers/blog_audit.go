package controllers

import(
	"sanskrit_api/models"
	// "github.com/astaxie/beego"
)

type AuditBlogController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *AuditBlogController) Get() {
	uid := this.getInt64("uid",-1)
	status := this.getInt("status",-2)
	maps,id,err := models.Blogs(uid,models.BLOG_STATUS,[...]interface{}{uid,uid,uid,uid,uid,status,this.page,this.pageSize})
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
// @router / [post]
func (this *AuditBlogController) Post() {
	id := this.getInt64("id",-1)
	status := this.getInt("status",-1)
	reason := this.getString("reason",0,0)	
	_,err :=models.SqlRaw("UPDATE zd_blog SET status=?,reason=? WHERE id=?",[...]interface{}{status,reason,id})
	if err != nil{
		this.false(-1,err)
	}
	this.response(id)
}