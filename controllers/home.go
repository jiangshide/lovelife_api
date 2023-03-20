package controllers

import (
	// "github.com/astaxie/beego"
	// "sanskrit_api/models"
)

// Operations about Users
type HomeController struct {
	BaseController
}

type Home struct {
	// Banner []*models.Banner `json:"banner"`
	// Nation []*models.Nation `json:"nation"`
	Name       string `json:"name"`
}

type Home1 struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Token interface{} `json:"token"`
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *HomeController) Get() {
	// home := new(Home)
	// banner := new(models.Banner)
	// home.Banner, _ = banner.List(this.pageSize, this.offSet)

	// nation := new(models.Nation)
	// home.Nation, _ = nation.List(-1, -1)

	// this.response(home)
	// home.name = "this is the jankey!"
	home1:=new(Home1)
	home1.Id  = 1
	home1.Name = "sallyshi"
	home1.Email = "sallyshi@huaxing.com"
	this.response(home1)
	// this.false(-1000,"the is error!")
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router / [post]
func (this *HomeController) Post() {
	// name := this.getString("name", 0, 0)
	// nice := this.getString("nice", 0, 0)
}
