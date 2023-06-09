package models

import "github.com/astaxie/beego/orm"

type Upload struct {
	Id         int64  `json:"id"`
	Cover      string `json:"icon"` //封面
	Name       string `json:"name"`
	Descript   string `json:"descript"`
	Type       int    `json:"type"`
	Format     string `json:"format"`
	Size       int64  `json:"size"`   //文件尺寸
	Width      int    `json:"width"`  //图片宽度
	Height     int    `json:"height"` //图片高度
	Length     int    `json:"length"` //视频时长
	Path       string `json:"path"`
	Url        string `json:"url"`   //访问地址
	Times      int    `json:"times"` //上传次数
	CreateId   int64  `json:"createId"`
	UpdateId   int64  `json:"updateId"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Views      int    `json:"views"`
}

func (this *Upload) TableName() string {
	return TableName("upload")
}

func (this *Upload) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Upload) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Upload) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Upload) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Upload) QueryAll() (list []*Nation, err error) {
	orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).All(&list)
	return
}

func (this *Upload) List(pageSize, offSet int) (list []*Nation, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}
