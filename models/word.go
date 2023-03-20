package models

import ("github.com/astaxie/beego/orm"
//  "github.com/astaxie/beego"	
 )


type Word struct {
	Id int64	`json:"id"`
	Uid   int64  `json:"uid"`//用户ID
	ContentId int64 `json:contentId`//内容来源ID
	Name string `json:"name"`//关键字名称
	Source int `json:"source"`//内容类型:1~用户,2~频道,3~动态,4~评论	
}

func (this *Word) TableName() string {		
	return TableName("word")
}

func (this *Word) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Word) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Word) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Word) ReadOrCreates(uid,contentId int64,name string,source int)(created bool, id int64, err error){
	word:=Word{Uid:uid,ContentId:contentId,Name:name,Source:source}
	created,id,err = orm.NewOrm().ReadOrCreate(&word,"uid","content_id","name","source")
	id,err = word.Update()
	return
}

var WORD_QUERY = "SELECT name,source FROM zd_word WHERE uid=? AND source=? order by update_time DESC limit ? offset ?"
var WORD_HISTORY = "SELECT name,source FROM zd_word WHERE source=? order by update_time DESC limit ? offset ?"


func (this *Word) Banner()(words []interface{}){
	blogs,_,_ := SqlList("SELECT title,name FROM zd_blog WHERE status=2 order by update_time DESC limit 5 offset 0",[...]interface{}{})
	for _,v := range *blogs{
		word := new(Word)
		word.Source = 3
		word.Name=v["title"].(string)
		if word.Name == ""{
			word.Name = v["name"].(string)
		}
		words = append(words,word)
	}
	channels,_,_ := SqlList("SELECT name FROM zd_channel WHERE status=2 order by update_time DESC limit 5 offset 0",[...]interface{}{})
	for _,v := range *channels{
		word := new(Word)
		word.Source = 2	
		word.Name=v["name"].(string)
		words = append(words,word)
	}
	users,_,_ := SqlList("SELECT UP.nick FROM zd_user_profile UP LEFT JOIN zd_uc_user U ON UP.id = U.id WHERE U.status=2 order by UP.update_time DESC limit 5 offset 0",[...]interface{}{})
	for _,v := range *users{
		word := new(Word)
		word.Source = 1	
		word.Name=v["nick"].(string)
		words = append(words,word)
	}
	return
}

func (this *Word) Search(name string)(words []interface{}){
	blogs,_,_ := SqlList("SELECT title FROM zd_blog WHERE status=2 AND instr(title,?) >0 order by update_time DESC limit 5 offset 0",[...]interface{}{name})
	for _,v := range *blogs{
		word := new(Word)
		word.Source = 3
		word.Name=v["title"].(string)
		words = append(words,word)
	}
	channels,_,_ := SqlList("SELECT name FROM zd_channel WHERE status=2 AND instr(name,?) >0  order by update_time DESC limit 5 offset 0",[...]interface{}{name})
	for _,v := range *channels{
		word := new(Word)
		word.Source = 2	
		word.Name=v["name"].(string)
		words = append(words,word)
	}
	users,_,_ := SqlList("SELECT UP.nick FROM zd_user_profile UP LEFT JOIN zd_uc_user U ON UP.id = U.id WHERE U.status=2 AND instr(UP.nick,?) >0 order by UP.update_time DESC limit 5 offset 0",[...]interface{}{name})
	for _,v := range *users{
		word := new(Word)
		word.Source = 1	
		word.Name=v["nick"].(string)
		words = append(words,word)
	}
	return
}

