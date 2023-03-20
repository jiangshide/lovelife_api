package models

import (
"github.com/astaxie/beego"
 "encoding/base64"
 "github.com/jiangshide/GoComm/utils"
 "strconv"
 "encoding/json"
 )

type ImUserList struct{
	Total int `json:"total"`
	Start int `json:"start"`
	Count int `json:"count"`
	Users []UserIM `json:"users"`
}

type UserIM struct{
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Mtime string `json:"mtime"`
	Ctime string `json:"ctime"`
}

type Notification struct{	
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	Action int `json:"action"`//类型:1~通告,2~点赞,3~粉丝,4~浏览,5~创建频道,6~发布动态
	Name string `json:"name"`
	Content string `json:"content"`
}

var Author = "Basic "+base64.StdEncoding.EncodeToString([]byte("cac7d5d277914d18c854d9c6:f042009ba021b23558f549ea"))
var JmsgUrl = "https://api.im.jpush.cn/v1/users/"
var JpushUrl = "https://api.jpush.cn/v3/push"

func RegisterIM(username,nickname,avatar string){	
	data := `[{"username":"`+username+`","password":"`+utils.Md5ToStr(username)+`","nickname":"`+nickname+`","avatar":"`+avatar+`"}]`
	Delete(username)
	http := new(Http)
	http.Url = JmsgUrl
	http.Author = Author
	http.Data = string(data)
	if response,err := http.Post();err != ""{
		beego.Info("RegisterIM~response:",response," | err:",err)
	}
}

func Delete(username string)string{
	http := new(Http)
	http.Url = JmsgUrl+username
	http.Author = Author
	_,err := http.Delete()
	return err
}

func UserList(model ImUserList){
	// /v1/users/?start={start}&count={count}
	http := new(Http)
	http.Author = Author
	http.Url = JmsgUrl+"?start=0&count=100"
	http.Model = model
	if err := http.Get();err != nil{
		beego.Info("UserList~err:",err)
	}
}

 func (this *Notification)PushAlia(alias int64){
 	// author := "Basic "+base64.StdEncoding.EncodeToString([]byte("cac7d5d277914d18c854d9c6:f042009ba021b23558f549ea"))
 	// data :=`{"platform":["android", "ios"],"audience":{"alias" : [ `+"2"+`]},"notification": {
  //       "android": {
  //           "alert": "Hi, JPush!",
  //           "title": "Send to Android",
  //           "builder_id": 1,
  //           "large_icon": "http://www.jiguang.cn/largeIcon.jpg",
  //           "intent": {
  //               "url": "intent:#Intent;component=com.jiguang.push/com.example.jpushdemo.SettingActivity;end",
  //           },
  //           "extras": {	
  //               "newsid": 321
  //           }
  //       }	
  //   }}`
 	jsonStr,err := json.Marshal(this)
 	beego.Info(err," | jsonStr:",string(jsonStr))

    data :=`{"platform":["android", "ios"],"audience":{"alias" : [ `+strconv.FormatInt(alias,10)+`]},"notification":{"android": {
            "alert": "`+this.Content+`",
            "title": "`+this.Name+`",
            "extras": `+string(jsonStr)+`
        }}}`
    http := new(Http)
	http.Url = JpushUrl
	http.Author = Author
	http.Data = data
	if response,err := http.Post();err != ""{
		beego.Info("PushAlia~response:",response," | err:",err)
	}
 } 


func (this *Notification) PushTag(tag string){
	jsonStr,err := json.Marshal(this)
 	beego.Info(err," | jsonStr:",string(jsonStr))

    data :=`{"platform":["android", "ios"],"audience":{"tag" : [ `+tag+`]},"notification":{"android": {
            "alert": "`+this.Content+`",
            "title": "`+this.Name+`",
            "extras": `+string(jsonStr)+`
        }}}`
    http := new(Http)
	http.Url = JpushUrl
	http.Author = Author
	http.Data = data
	if response,err := http.Post();err != ""{
		beego.Info("PushTag~response:",response," | err:",err)
	}
}

func PushAll(){
 	// author := "Basic "+base64.StdEncoding.EncodeToString([]byte("cac7d5d277914d18c854d9c6:f042009ba021b23558f549ea"))
 	data :=`{"platform":["android", "ios"],"audience":"all","notification":{"alert":"Hi,JPush~the test!!!"}}`
 	http := new(Http)
	http.Url = JpushUrl
	http.Author = Author
	http.Data = data
	if response,err := http.Post();err != ""{
		beego.Info("PushAll~response:",response," | err:",err)
	}
 } 

 
