package controllers

import(
	"sanskrit_api/models"
	// "github.com/astaxie/beego"
)

// Users APi
type AutitUserController struct {
	BaseController
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router / [get]
func (this *AutitUserController) Get() {
	status := this.getInt("status",-1)
	sql := models.UserSql+",P.create_time createDate,P.update_time updateTime,U.name,U.ip,U.status,U.reason,U.online,(SELECT count(1) FROM zd_channel C WHERE C.uid=P.id) channelsNum,(SELECT count(1) FROM zd_blog B WHERE B.uid=P.id) blogsNum FROM zd_user_profile P LEFT JOIN zd_uc_user U ON P.id=U.id WHERE U.status=? ORDER BY P.update_time DESC LIMIT ? OFFSET ?"
	maps,id,err := models.SqlList(sql,[...]interface{}{status,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null!")	
	}else{
		this.response(maps)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router / [post]
func (this *AutitUserController) Post() {
	id := this.getInt64("id",-1)
	updateId := this.getInt64("updateId",-1)
	status := this.getInt("status",-1)
	_,err :=models.SqlRaw("UPDATE zd_uc_user SET status=?,update_id=? WHERE id=?",[...]interface{}{id,updateId,status})
	if err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}