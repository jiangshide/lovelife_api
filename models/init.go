package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"reflect"
	"github.com/jiangshide/GoComm/utils"
	// "sanskrit_api/socket"
	"time"
	// "fmt"
	"encoding/json"
)

func init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPsw := beego.AppConfig.String("db.psw")
	dbName := beego.AppConfig.String("db.name")
	timeZone := beego.AppConfig.String("db.timezone")
	maxConn, _ := beego.AppConfig.Int("maxConn")
	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	if dbPort == "" {
		dbPort = "3306"
	}
	// dns := dbUser + ":" + dbPsw + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	dns := dbUser + ":" + dbPsw + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?"
	
	if timeZone != "" {
		dns += "&loc=" + url.QueryEscape(timeZone)
	}
	orm.RegisterDataBase("default", "mysql", dns, maxConn, maxIdle)
	orm.RegisterModel(new(User),new(Pay),new(WeiXinInfo),new(Profile),new(Active),new(Frequency),new(Certification),new(Remarks),new(Course),new(Invite),new(Friend),new(UserTop),new(UserFollow),new(UserRecommend),new(Position),new(DeviceInfo),new(Device),new(Channel),new(ChannelNature),new(ChannelNatureUser),new(ChannelType),new(ChannelTypeUser),new(ChannelTop),new(ChannelRecommend),new(Blog),new(BlogStyle),new(BlogPraise),new(BlogRecommend),new(BlogTop),new(BlogView),new(BlogCollection),new(BlogShare),new(Format),new(File),new(Top),new(Praise),new(Comment),new(CommentPraise),new(CommentUid),new(CommentUidPraise),new(At),new(Share),new(View),new(History),new(Report),new(ReportType),new(App),new(AppCount),new(AppAds),new(Word))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	InitFormat()
	InitChannelNature()	
	InitChannelType()
	InitReportType()
	InitAppAds()
	// InitImId(1)
	// InitChannel(0,60,"徐亮")
	// InitChannel(0,61,"哄哄")
	
	// socket.StartWebsocket("192.168.3.7:8093")
	startTimer(updateUserState)
	// updateImSign()
	// importImUser()
}

func startTimer(f func()) {    //每个24小时检查一次    
	go func() {            
		for {
            f()
                now := time.Now()                // 计算下一个零点
                next := now.Add(time.Hour * 24)
                next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
                t := time.NewTimer(next.Sub(now))
                <-t.C
        }
    }()
}

func updateUserState(){
	SqlRaw("UPDATE zd_user_profile SET online=-1 WHERE update_time < now()-72000 ",[...]interface{}{})//20个小时未更新的设置为离线状态
}

func updateImSign(){
	profile := new(Profile)
	list,count := profile.List(100,0)
	for _,v := range(list){
		beego.Info("list:",v.Id," | count:",count)
		InitImId(v.Id)
	}
}

func importImUser(){
	profile := new(Profile)
	list,count := profile.List(100,0)
	for _,v := range(list){
		beego.Info("list:",v.Id," | count:",count)
		imUser := new(ImUser)
		imUser.Identifier = v.Id
		imUser.Nick = v.Nick
		imUser.FaceUrl = v.Icon
		if jsonBytes,err := json.Marshal(imUser);err == nil && v.Id == 3{
			http := new(Http)
			http.Url = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?usersig=eJwtzEELgjAYxvHvsmthr3NLJ3SoSyEdhIooumhb*mrZmJJm9N0z9fj8Hvh-yH67s17KEJ9QC8i03yhVUeENe84wKpIy7Wx8S5lHWqMkvs0AqOdywYZHNRqN6pxzTgFg0Aoff5sDpcL2hDNWMOnigR0Gy3DCngdWh8ocN7SOzy00cSbuqRudVk7L1vll9tbxdUG*P5wbM7k_&identifier=jiangshide&sdkappid=1400287594"
			http.Data = string(jsonBytes)
			http.Post()
		}
	}
}

type ImUser struct{
	Identifier int64
	Nick string
	FaceUrl string
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

/**
遗留问题: 单词分割
 */
func Field(model interface{}) (fieldName string, fieldValue interface{}) {
	if model == nil {
		return fieldName, fieldValue
	}
	var field reflect.Type
	var value reflect.Value
	field = reflect.TypeOf(model).Elem()
	value = reflect.ValueOf(model).Elem()
	size := field.NumField()

	for i := 0; i < size; i++ {
		v := value.Field(i)
		fieldName = utils.StrFirstToLower(field.Field(i).Name) //TODO 需要解决单词分割问题
		switch value.Field(i).Kind() {
		case reflect.Bool:
			fieldValue = v.Bool()
		case reflect.Int:
			if v.Int() != 0 {
				fieldValue = v.Int()
			}
		case reflect.Int64:
			if v.Int() != 0 {
				fieldValue = v.Int()
			}
		case reflect.String:
			fieldValue = v.String()
		}
		if fieldValue != nil && fieldValue != 0 {
			break
		}
	}
	return
}

func Field1(model interface{}, name string) (fieldName string, fieldValue interface{}) {
	var field reflect.Type
	var value reflect.Value
	field = reflect.TypeOf(model).Elem()
	value = reflect.ValueOf(model).Elem()
	size := field.NumField()
	for i := 0; i < size; i++ {
		fieldName = utils.StrFirstToLower(field.Field(i).Name)
		if name == fieldName {
			fieldValue = value.Field(i)
			break
		}
	}
	return
}

func SqlRaw(sql string,ids interface{})(res interface{},err error){
	res,err = orm.NewOrm().Raw(sql,ids).Exec()
	return
}

func SqlRawRow(sql string,ids interface{},model interface{})(err error){
	err = orm.NewOrm().Raw(sql,ids).QueryRow(&model)
	return
}


func SqlCount(table string,maps map[string]interface{})(count int64,err error){
	query := orm.NewOrm().QueryTable(table)
	for k,v := range maps{
		beego.Info("k:",k," | v:",v)
		query = query.Filter(k,v)
	}
	count,err = query.Count()
	return
}

func Sql(sql string,ids interface{})(res interface{},id int64,err error) {
	var maps []orm.Params
 	o := orm.NewOrm()
    id,err = o.Raw(sql,ids).Values(&maps)
   if len(maps) > 0{
   		res = maps[0]
   }
    return res,id,err
}

func SqlList(sql string,ids interface{}) (list *[]orm.Params,id int64,err error) {
 	var maps []orm.Params
 	o := orm.NewOrm()
    id,err = o.Raw(sql,ids).Values(&maps)
    return &maps,id,err
}

