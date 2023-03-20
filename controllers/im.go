package controllers

import (
"sanskrit_api/models"
"github.com/astaxie/beego"
// "fmt"
)

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /imUserList [get]
func (this *UserController) ImUserList() {
	// pageSize := this.getInt("pageSize", 20)
	// page := this.getInt("page", 0)
	var imUserList models.ImUserList
	http := new(models.Http)
	http.Author = models.Author
	http.Url = models.JmsgUrl+"?start="+"0"+"&count="+"300"
	http.Model = &imUserList
	err := http.Get()
	if err != nil{
		beego.Info("im~err:",err," | imUserList:",imUserList)
		this.false(-1,err)
	}else{
		if imUserList.Total == 0{
			this.false(-1,"total is 0")
		}else{
			this.response(imUserList.Users)
		}
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /imUserDelete [get]
func (this *UserController) ImUserDelete() {
	userName := this.getString("userName",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	
	err := models.Delete(userName)
	beego.Info("err:",err)
	if len(err) > 5{
		beego.Info("im~delete~err:",err)
		this.false(-1,err)
	}else{
		this.response(true)
	}
	
}