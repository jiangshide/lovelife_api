package controllers

import (
"github.com/astaxie/beego"
// "github.com/astaxie/beego/orm"
	"sanskrit_api/models"
	"github.com/jiangshide/GoComm/utils"
	"strings"		
	"errors"
	"fmt"	
	"encoding/json"	
	"sanskrit_api/util"
	"time"
	// "runtime"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
      terror  "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors" 
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
        sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

var validateMap = make(map[string]string)

// Users APi
type UserController struct {
	BaseController
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router / [get]	
func (this *UserController) List() {
	status := this.getInt("status",2)
	name :=  this.getString("name",0,0)
	uid := this.getInt64("uid",0)
	if len(name) > 0 {
		maps,id,err := models.SqlList(models.USER_NAME,[...]interface{}{uid,uid,uid,uid,name,name,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null")	
		}else{
			this.response(maps)	
		}
	}else{
		maps,id,err := models.SqlList(models.USER_STATUS,[...]interface{}{uid,uid,uid,uid,status,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null")	
		}else{
			this.response(maps)	
		}
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /userExit [get]
func (this *UserController) UserExit() {
	user := new(models.User)
	user.Name = this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	err := user.GetByName()
	if err != nil{
		// this.false(-1,err)
		this.response(false)
	}
	this.response(true)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /reg [post]
func (this *UserController) Reg() {
	user := new(models.User)
	user.Name = this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	password := this.getString("password", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	validateCode := this.getString("validateCode", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	country := this.getString("country",0,0)
	icon := this.getString("icon",0,0) 
	sex := this.getInt("sex",0)

	latitude := this.getFloat("latitude",0.0)
	longitude := this.getFloat("longitude",0.0)
	locationType := this.getString("locationType",0,0)
	adCode := this.getString("adCode",0,0)


	netInfo := this.getString("netInfo",0,0)
	device:= this.getString("device",0,0)
	ip := this.getIp()

	if(!strings.EqualFold(validateCode,validateMap[adCode+user.Name])){
		models.AddMake(netInfo,device,ip,0,0,REG,errors.New("验证码错误"))
		this.false(VALIDATE_FAIL,msg[VALIDATE_FAIL])	
	}
	user.Salt = utils.GetRandomString(10)
	user.Password = utils.Md5ToStr(password + user.Salt)
	user.Status=2
	user.Ip = ip
	id, err := user.Add();
	if err != nil {
		models.AddMake(netInfo,device,ip,0,0,REG,err)
		this.false(DB_INSERT_FALSE, err)
	}

	position := models.AddMake(netInfo,device,ip,id,0,REG,nil)
	profile:=new(models.Profile)
	profile.Id = id
	profile.UnionId = user.Name
	profile.Online=1
	profile.Nick = utils.GetRandomName()
	profile.Intro=DEFAULT_INTRO
	profile.Icon = icon
	profile.Sex = sex
	profile.Latitude = latitude
	profile.Longitude = longitude
	profile.LocationType = locationType
	profile.Country = country
	profile.City = position.City
	profile.AdCode  = adCode
	profile.FollowNotice = 1
	_,err = profile.Add()
	if err != nil{
		models.AddMake(netInfo,device,ip,id,0,REG,err)
		this.false(-1,err)		
	}

	models.InitFrequency(id,0,time.Now().Unix(),2,1)
	models.InitFrequency(id,0,time.Now().Unix(),6,1)
	models.InitUserChannelNature(id,netInfo,device,ip,REG)
	models.InitChannel(profile.ChannelId,profile.Id,profile.Nick)
	
	models.CourseAdd(id,id,1,"世界，我来了!",DEFAULT_INTRO,icon,"")
	models.InitImId(id)
	if res,ids,err := models.Sql(models.USER_ID, [...]interface{}{id,id,id,id,id});err != nil || ids == 0{
		this.false(-1,err)
	}else{
		models.RegisterIM(user.Name,profile.Nick,icon)
		this.response(res)
	}

	// if err = profile.Query();err != nil{
	// 	this.false(-1,err)
	// }else{
	// 	models.RegisterIM(user.Name,profile.Nick,icon)
	// 	beego.Info("-----------profile:",profile)
	// 	this.response(profile)	
	// }
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /codeLogin [post]
func (this *UserController) CodeLogin() {
	adCode := this.getString("adCode",0,0)
	name := this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	validateCode := this.getString("validateCode", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	if(!strings.EqualFold(validateCode,validateMap[adCode+name])){
		this.false(VALIDATE_FAIL,msg[VALIDATE_FAIL])
	}else{		
		user := new(models.User)
		user.Name = name
		if err := user.GetByName();err != nil{
			this.false(USER_NOT_EXIST,msg[USER_NOT_EXIST])//第一次验证码登录需要设置密码
		}else{	
			models.InitImId(user.Id)
			if res,_,err := models.Sql(models.USER_ID, [...]interface{}{user.Id,user.Id,user.Id,user.Id,user.Id});err != nil || res == nil{
				this.false(USER_SET_INFO,msg[USER_SET_INFO])
			}else{
				models.SqlRaw("UPDATE zd_user_profile SET online=? WHERE id=?",[]interface{}{1});
				this.response(res)
			}
		}
	}
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (this *UserController) Login() {
	// this.checkToken()
	user := new(models.User)
	user.Name = this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	password := this.getString("password", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	latitude := this.getFloat("latitude",0.0)
	longitude := this.getFloat("longitude",0.0)
	locationType := this.getString("locationType",0,0)

	netInfo := this.getString("netInfo",0,0)
	device := this.getString("device",0,0)
	ip := this.getIp()

	if err := user.Query("name",user.Name); err != nil {
		models.AddMake(netInfo,device,ip,0,0,LOGIN,err)
		this.false(USER_NOT_EXIST,msg[USER_NOT_EXIST])
	}
	if user.Status == USER_FORBIDDEN {
		this.false(USER_FORBIDDEN, nil)
	}else if user.Status == -1 ||user.Status == -2 || user.Status == -4{
		this.false(USER_EXCEPTION, nil)
	} else if user.Password != utils.Md5ToStr(password+user.Salt) {
		this.false(USER_PASSWORD_ERR, nil)
	}
	models.Location(latitude,longitude,locationType,user.Id)
	models.InitImId(user.Id)
	if res,_,err := models.Sql(models.USER_ID, [...]interface{}{user.Id,user.Id,user.Id,user.Id,user.Id});err != nil{
		models.AddMake(netInfo,device,ip,user.Id,0,LOGIN,err)
		this.false(USER_SET_INFO,msg[USER_SET_INFO])
	}else{
		models.AddMake(netInfo,device,ip,user.Id,user.Id,LOGIN,nil)	
		models.SqlRaw("UPDATE zd_user_profile SET online=? WHERE id=?",[]interface{}{1});
		this.response(res)
	}
}

// @Title WeChat
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for WeChat"
// @Param	password		query 	string	true		"The password for WeChat"
// @Success 200 {string} WeChat success
// @Failure 403 user not exist
// @router /weChat [post]
func (this *UserController) WeChat() {
	code := this.getString("code", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	latitude := this.getFloat("latitude",0.0)
	longitude := this.getFloat("longitude",0.0)
	locationType := this.getString("locationType",0,0)
	adCode := this.getString("adCode",0,0)

	netInfo := this.getString("netInfo",0,0)
	device := this.getString("device",0,0)
	ip := this.getIp()

	var weixinToken models.WeiXinToken
	tokenUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx1bd7f51f4f97f248&secret=12bf4ced0f3cfc045dea802b835c555c&code="+code+"&grant_type=authorization_code"
	http := new(models.Http)
	http.Url = tokenUrl
	http.Model = &weixinToken
	err := http.Get()
	if err != nil{
		models.AddMake(netInfo,device,ip,0,0,REG,err)
		this.false(-1,err)
	}
	var weiXin models.WeiXin
	weixinInfoUrl := "https://api.weixin.qq.com/sns/userinfo?access_token="+weixinToken.AccessToken+"&openid="+weixinToken.Openid
	http.Url = weixinInfoUrl
	http.Model = &weiXin
	err = http.Get()
	if err != nil{		
		models.AddMake(netInfo,device,ip,0,0,REG,err)
		this.false(-1,err)
	}
	weixinInfo := new(models.WeiXinInfo)
	weixinInfo.Openid = weiXin.Openid
	err = weixinInfo.Query()
	if err != nil{
		user := new(models.User)
		// user.Name = weiXin.Unionid
		user.Status = 2
		user.Ip = ip
		id,err := user.Add()
		if err != nil{
			models.AddMake(netInfo,device,ip,0,0,REG,err)
			this.false(DB_INSERT_FALSE, err)
		}
		weixinInfo.Uid = id
		weixinInfo.Nick = weiXin.NickName
		weixinInfo.Sex = weiXin.Sex
		weixinInfo.Language = weiXin.Language
		weixinInfo.City = weiXin.City
		weixinInfo.Province = weiXin.Province
		weixinInfo.Country = weiXin.Country
		weixinInfo.Icon = weiXin.Headimgurl
		weixinInfo.Privilege = fmt.Sprintf("%S",weiXin.Privilege)
		weixinInfo.Unionid = weiXin.Unionid
		weixinInfo.Add()

		profile:=new(models.Profile)
		profile.Id = id
		profile.UnionId=weixinInfo.Unionid
		profile.Online=1
		profile.Nick = weixinInfo.Nick
		profile.Intro=DEFAULT_INTRO
		profile.Icon = weixinInfo.Icon
		profile.Sex = weixinInfo.Sex
		profile.Latitude = latitude
		profile.Longitude = longitude
		profile.LocationType = locationType
		profile.Country = weixinInfo.Country
		profile.Province = weixinInfo.Province
		profile.City = weixinInfo.City
		profile.AdCode = adCode
		profile.FollowNotice = 1
		if len(profile.City) == 0{
			profile.City = profile.Province
		}
		_,err = profile.Add()
		if err != nil{
			models.AddMake(netInfo,device,ip,id,0,REG,err)
			this.false(-1,err)
		}
		models.InitFrequency(id,0,time.Now().Unix(),2,1)
		models.InitFrequency(id,0,time.Now().Unix(),6,1)
		models.InitUserChannelNature(id,netInfo,device,ip,REG)
		models.AddMake(netInfo,device,ip,id,0,REG,nil)
		models.InitChannel(profile.ChannelId,profile.Id,profile.Nick)

		models.CourseAdd(id,id,1,"世界，我来了!",DEFAULT_INTRO,profile.Icon,"")
		models.InitImId(id)
		if res,_,err := models.Sql(models.USER_ID, [...]interface{}{id,id,id,id,id});err != nil{
			this.false(-1,err)
		}else{
			models.RegisterIM(weiXin.Unionid,profile.Nick,profile.Icon)
			// models.InitChannel(res["channelId"].(int64),id)
			this.response(res)	
		}
	}else{
		user := new(models.User)
		// user.Name = weiXin.Unionid
		user.Status = 2
		user.Ip = ip
		user.Update()
		if weiXin.NickName != weixinInfo.Nick || weiXin.Headimgurl != weixinInfo.Icon || weiXin.Sex != weixinInfo.Sex || fmt.Sprintf("%S",weiXin.Privilege) != weixinInfo.Privilege{
			weixinInfo.Nick = weiXin.NickName
			weixinInfo.Sex = weiXin.Sex
			weixinInfo.Language = weiXin.Language
			weixinInfo.City = weiXin.City
			weixinInfo.Province = weiXin.Province
			weixinInfo.Country = weiXin.Country
			weixinInfo.Icon = weiXin.Headimgurl
			weixinInfo.Privilege = fmt.Sprintf("%S",weiXin.Privilege)
			weixinInfo.Unionid = weiXin.Unionid
			weixinInfo.Update()

			profile:=new(models.Profile)
			profile.Id = weixinInfo.Uid
			profile.UnionId=weixinInfo.Unionid
			profile.Online=1
			profile.Nick = weixinInfo.Nick
			profile.Icon = weixinInfo.Icon
			profile.Sex = weixinInfo.Sex
			profile.Latitude = latitude
			profile.Longitude = longitude
			profile.LocationType = locationType
			profile.Country = weixinInfo.Country
			profile.Province = weixinInfo.Province
			profile.City = weixinInfo.City
			profile.AdCode = adCode
			if len(profile.City) == 0{
				profile.City = profile.Province
			}
			profile.Update()	
			models.InitChannel(profile.ChannelId,profile.Id,profile.Nick)
		}
		models.InitImId(weixinInfo.Uid)
		if res,_,err := models.Sql(models.USER_ID, [...]interface{}{weixinInfo.Uid,weixinInfo.Uid,weixinInfo.Uid,weixinInfo.Uid,weixinInfo.Uid});err != nil{
			models.AddMake(netInfo,device,ip,weixinInfo.Uid,0,LOGIN,err)
			this.false(-1,err)
		}else{
			models.Location(latitude,longitude,locationType,weixinInfo.Uid)
			models.CourseAdd(weixinInfo.Uid,weixinInfo.Uid,2,"美好的一天从这里开始...",weixinInfo.Nick,weixinInfo.Icon,"")
			this.response(res)
		}
	}
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /bind [post]
func (this *UserController) Bind() {
	// this.checkToken()
	id := this.getInt64("id",-1)
	name := this.getString("name",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	password := this.getString("password",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	validateCode := this.getString("validateCode", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	adCode := this.getString("adCode",0,0)
	if(!strings.EqualFold(validateCode,validateMap[adCode+name])){
		this.false(VALIDATE_FAIL,msg[VALIDATE_FAIL])
	}else{	
		netInfo := this.getString("netInfo",0,0)
		device := this.getString("device",0,0)
		ip := this.getIp()
		user := new(models.User)
		user.Id = id
		if err := user.Query("id",user.Id);err != nil{
			models.AddMake(netInfo,device,ip,id,0,BIND,err)
			this.false(-1,err)
		}
		user.Name = name
		user.Salt = utils.GetRandomString(10)
		user.Password = utils.Md5ToStr(password + user.Salt)
		user.Status=2
		
		if _,err := user.Update();err != nil{
			models.AddMake(netInfo,device,ip,id,0,BIND,err)
			this.false(-1,err)
		}

		// models.InitImId(id)
		if res,_,err := models.Sql(models.USER_ID, [...]interface{}{id,id,id,id,id});err != nil{
			this.false(-1,err)
		}else{
			this.response(res)
		}
	} 
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /validate [post]
func (this *UserController) Validate() {
	user := new(models.User)
	user.Name = this.getString("username", DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	code := user.GenValidateCode(6)
	validateMap[user.Name]=code

	credential := common.NewCredential(
                "AKIDyGPMsE9apkW3RnLD4mcD1conXz9Dv3IK",
                "4sOfkYobliU1GFzSeZOwDCtqHRM1SeBc",
        )
        cpf := profile.NewClientProfile()
        cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
        client, _ := sms.NewClient(credential, "", cpf)
        
        request := sms.NewSendSmsRequest()

        // params := "{\"PhoneNumberSet\":[\"+8618311271399\"],\"TemplateID\":\"718159\",\"Sign\":\"梵山科技\",\"TemplateParamSet\":[\"23423\",\"10\"],\"SmsSdkAppid\":\"1400425697\"}"
        params := "{\"PhoneNumberSet\":[\""+user.Name+"\"],\"TemplateID\":\"718159\",\"Sign\":\"梵山科技\",\"TemplateParamSet\":[\""+code+"\",\"10\"],\"SmsSdkAppid\":\"1400425697\"}"
        beego.Info("params:",params)
        if err := request.FromJsonString(params);err != nil{
        	beego.Info(err)
        	this.false(-1,err)
        	return
        }
        _, err := client.SendSms(request)
        if _, ok := err.(*terror.TencentCloudSDKError); ok || err != nil{
                beego.Info("An API error has returned: %s", err)
                this.false(-1,"An API error has returned")
        }else{
        	this.response(true)
        }
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /profiles [get]
func (this *UserController) Profiles() {
	uid := this.getInt64("uid",-1)
	fromId := this.getInt64("fromId",-1)
	if res,ids,err := models.Sql(models.USER_ID, [...]interface{}{uid,uid,uid,fromId,fromId});err != nil || ids==0{
		this.false(-1,"data is null")
	}else{
		this.response(res)
		if uid != fromId{
			profile := new(models.Profile)
			profile.Id = uid
			profile.Query()
			if profile.BroweHomeNotice == 1{
				notification := new(models.Notification)
				notification.Id = uid
				notification.Uid= fromId
				notification.Action=5
				notification.Name = "来自《"+profile.Nick+"》的浏览通知"
				notification.Content=profile.Nick+"查看了你的主页,赶快去关注TA吧!"
				notification.PushAlia(fromId)
			}
		}	
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /profile [post]
func (this *UserController) Profile() {
	p := new(models.Profile)
	err := json.Unmarshal([]byte(this.getString("profile",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)),&p)
	types := this.getInt("type",0)
	if types >0{
		frequency := new(models.Frequency)
		frequency.Uid = p.Id
		frequency.Type = types
		err :=frequency.Query()
		if err != nil{
			this.false(-1,err)
		}else{
			currTime := time.Now().Unix()
			diffTime := currTime-frequency.Time-frequency.Frequency
			if frequency.Number > 0 && diffTime > 0{
				profile := new(models.Profile)
				if len(p.Birthday) > 0{
					y, m, d := util.GetTimeFromStrDate(p.Birthday) 
					p.Age = util.GetAge(y)
					p.Constellation = util.GetConstellation(m,d)
					p.Zodiac = util.GetZodiac(y)
				} 
				profile.Id = p.Id
				if err = profile.Query();err != nil{
					if _,err := p.Add();err != nil{
						this.false(-1,err)
					}
					this.response(p)
				} else{
					if _,err := p.Update();err != nil{
						this.false(-1,err)
					}
					models.SqlRaw(models.Frequency_UPDATE,[...]interface{}{p.Id,types,1,2592000,time.Now().Unix(),frequency.Id})
					this.response(p)
				}
			}else{
				this.false(-3,"可修改次数已用完!")	
			}
		}
	}else{
		profile := new(models.Profile)
		if len(p.Birthday) > 0{
			y, m, d := util.GetTimeFromStrDate(p.Birthday) 
			p.Age = util.GetAge(y)
			p.Constellation = util.GetConstellation(m,d)
			p.Zodiac = util.GetZodiac(y)
		} 
		profile.Id = p.Id
		if err = profile.Query();err != nil{
			if _,err := p.Add();err != nil{
				this.false(-1,err)
			}
			this.response(p)
		} else{
			if _,err := p.Update();err != nil{
				this.false(-1,err)
			}
			this.response(p)
		}
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /certification [get]
func (this *UserController) Certification() {
	uid := this.getInt64("uid",-1)
	certification := new(models.Certification)
	certification.Uid = uid
	err := certification.Query()
	if err != nil{
		this.false(-1,err)
	}else{
		this.response(certification)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /certificationList [get]
func (this *UserController) CertificationList() {
	status := this.getInt("status",-1)
	maps,id,err := models.SqlList(models.CERTIFICATION_STATUS,[...]interface{}{status,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}
	this.response(maps)
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /certificationUpdate [post]
func (this *UserController) CertificationUpdate() {
	certification := new(models.Certification)
	json.Unmarshal([]byte(this.getString("certification",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)),&certification)
	id:=certification.Id
	uid:=certification.Uid
	name := certification.Name
	idCard := certification.IdCard
	idCardPicFront := certification.IdCardPicFront
	idCardPicBehind := certification.IdCardPicBehind
	url := certification.Url
	status := certification.Status
	reason := certification.Reason
	_,err := models.SqlRaw(models.CERTIFICATION_UPDATE,[...]interface{}{id,uid,name,idCard,idCardPicFront,idCardPicBehind,url,status,reason})
	if err != nil{
		this.false(-1,err)
	}
	this.response(status)
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /remarks [post]
func (this *UserController) Remarks() {
	uid := this.getInt64("uid",-1)
	fromId := this.getInt64("fromId",-1)
	name := this.getString("name",DEFAULT_ISNULL,DEFAULT_MIN_SIZE)
	url := this.getString("url",0,0)

	remarks := new(models.Remarks)
	if _,id,err := remarks.ReadOrCreates(uid,fromId,name,url);err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /course [get]
func (this *UserController) Course() {
	uid := this.getInt64("uid",-1)
	sql := "SELECT id,uid,name,des,cover,url,source,source_id sourceId,update_time date FROM zd_user_course WHERE uid=? order by create_time limit ? offset ?"
	maps,id,err := models.SqlList(sql,[...]interface{}{uid,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null!")	
	}else{
		this.response(maps)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /friendAdd [post]
func (this *UserController) FriendAdd() {
	uid:=this.getInt64("uid",-1)
	fromId:=this.getInt64("fromId",-1)
	status:=this.getInt("status",-1)
	reason:=this.getString("reason",0,0)
	url:=this.getString("url",0,0)

	friend := new(models.Friend)
	if _,id,err := friend.ReadOrCreates(uid,fromId,status,reason,url);err != nil{
		this.false(-1,err)
	}else{
		this.response(id)
	}
}
			
// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /friend [get]
func (this *UserController) Friend() {
	fromId := this.getInt64("fromId",-1)
	status:=this.getInt("status",-1)

	maps,id,err := models.SqlList(models.FRIEND_LIST,[...]interface{}{fromId,status,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null!")	
	}else{
		this.response(maps)
	}
	
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /follow [get]
func (this *UserController) Follow() {
	action := this.getInt("action",-1)
	uid := this.getInt64("uid",-1)
	if action == 1{
		maps,id,err := models.SqlList(models.USER_FOLLOWED,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
		if err != nil || id == 0{
			this.false(-1,"data is null!")	
		}else{
			this.response(maps)
		}
	}else{
		maps,id,err := models.SqlList(models.USER_RECOMMEND,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
		if err != nil || id == 0{
			// this.false(-1,"data is null!")	
			maps,id,err := models.SqlList(models.USER_FOLLOW,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
			if err != nil || id == 0{
				this.false(-1,"data is null!")	
			}else{
				this.response(maps)
			}
		}else{
			this.response(maps)
		}
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /followAdd [post]
func (this *UserController) FollowAdd() {
	follow := new(models.UserFollow)
	follow.Uid = this.getInt64("uid",-1)
	follow.FromId = this.getInt64("fromId",-1)
	follow.Status = this.getInt("status",-1)
	if _,_,err := follow.ReadOrCreates();err != nil{
		this.false(-1,err)
	}else{

		this.response(follow.Status)
		

		fromProfile := new(models.Profile)
		fromProfile.Id = follow.Uid
		fromProfile.Query()

		if follow.Status > 0 && fromProfile.FollowNotice == 1{
			profile := new(models.Profile)
			profile.Id = follow.FromId
			profile.Query()

			notification := new(models.Notification)
			notification.Id = profile.Id
			notification.Uid= follow.FromId
			notification.Action=3
			notification.Name = "来自《"+profile.Nick+"》的关注通知"
			notification.Content=profile.Nick+"关注了你,点击查看TA吧!"
			notification.PushAlia(follow.Uid)
		}
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /fans [get]
func (this *UserController) Fans() {
	uid := this.getInt64("uid",-1)
	maps,id,err := models.SqlList(models.USER_FANS,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null!")	
	}else{
		this.response(maps)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /praise [get]
func (this *UserController) Praise() {
	blogId := this.getInt64("blogId",-1)
	status := this.getInt("status",1)
	maps,id,err := models.SqlList("SELECT BV.create_time date,UP.id,UP.icon,UP.nick,UP.age,UP.sex,UP.city,UP.online,UP.online_id onlineId,UP.online_name onlineName,UF.id followId,UF.status follows FROM zd_blog_praise BV LEFT JOIN zd_user_profile UP ON UP.id=BV.uid LEFT JOIN zd_user_follow UF ON UF.uid=UP.id WHERE BV.blog_id=? AND BV.status=? ORDER BY BV.update_time DESC LIMIT ? OFFSET ? ",[...]interface{}{blogId,status,this.page,this.pageSize})
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
// @router /view [get]
func (this *UserController) View() {
	blogId := this.getInt64("blogId",-1)
	maps,id,err := models.SqlList("SELECT BV.create_time date,UP.id,UP.icon,UP.nick,UP.age,UP.sex,UP.city,UP.online,UP.online_id onlineId,UP.online_name onlineName,UF.id followId,UF.status follows FROM zd_blog_view BV LEFT JOIN zd_user_profile UP ON UP.id=BV.uid LEFT JOIN zd_user_follow UF ON UF.uid=UP.id WHERE BV.blog_id=? ORDER BY BV.update_time DESC LIMIT ? OFFSET ? ",[...]interface{}{blogId,this.page,this.pageSize})
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
// @router /activeAdd [post]
func (this *UserController) ActiveAdd() {
	uid := this.getInt64("uid",-1)
	contentId := this.getInt64("contentId",0)
	from := this.getInt("from",-1)
	if num,err := models.AddScore(uid,contentId,from);num > 0{
		this.response(num)
	}else{
		this.false(-1,err)
	}
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Upload.Id
// @Failure 403 body is empty
// @router /active [get]
func (this *UserController) Active() {
	uid := this.getInt64("uid",-1)
	from := this.getInt("from",-1)
	maps,id,err := models.SqlList("SELECT num,update_time date FROM zd_user_active WHERE uid=? AND source=? ORDER BY update_time DESC LIMIT ? OFFSET ? ",[...]interface{}{uid,from,this.page,this.pageSize})
	if err != nil || id == 0{
		this.false(-1,"data is null")	
	}
	this.response(maps)
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (this *UserController) Logout() {
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {

}

// // @Title Get
// // @Description get user by uid
// // @Param	uid		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.User
// // @Failure 403 :userId is empty
// // @router /:userId [get]
// func (this *UserController) Get() {
// 	profile := new(models.Profile)
// 	profile.Id = this.getInt64("uid",-1)
// 	if err := profile.Query();err != nil{
// 		this.false(-1,"data is null")
// 	}else{
// 		this.response(profile)
// 	}
// }

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (this *UserController) Delete() {
	user := new(models.User)
	user.Id = this.getInt64("userId", 0)
	if _, err := user.Del(); err != nil {
		this.false(DB_DELETE_FALSE, err)
	}
	this.Delete()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /online [post]
func (this *UserController) Online() {
	uid:=this.getInt64("uid",-1)
	online:=this.getInt("online",-1)
	id := this.getInt64("id",-1)
	name:=this.getString("name",0,0)

	if _,err := models.SqlRaw("UPDATE zd_user_profile SET online=?,online_id=?,online_name=? WHERE id=?",[]interface{}{online,id,name,uid});err != nil{
		this.false(-1,err)
	}else{
		this.response(true)
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /onlineList [get]
func (this *UserController) OnlineList() {
	status:=this.getInt("status",2)
	online:=this.getInt("online",1)

	maps,id,err := models.SqlList(models.USER_ONLINE,[...]interface{}{status,online})
	if err != nil || id == 0{
		this.false(-1,"data is null")
	}
	this.response(maps)
}
