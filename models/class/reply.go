package class

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Reply struct {
	Id      int
	Content string    `orm:"type(text)"`
	Article *Article  `orm:"rel(fk)"`
	Author  *User     `orm:"rel(fk)"`
	Time    time.Time `orm:"auto_now_add"`

	Default bool
}

func (a *Reply) Create() (n int64, err error) {
	o := orm.NewOrm()
	if n, err = o.Insert(a); err != nil {
		beego.Info(err)
	}
	return
}

func (a *Reply) Gets() (rets []*Reply) {
	o := orm.NewOrm()
	qs := o.QueryTable("reply")
	if a.Article != nil {
		qs = qs.Filter("article_id", a.Author.Id)
	}
	if a.Author != nil {
		qs = qs.Filter("author_id", a.Author.Id)
	}

	qs = qs.Filter("defunct", 0)

	qs.All(&rets)

	for k := range rets {
		rets[k].Article.ReadDB()
		rets[k].Article.ReadDB()
	}

	return
}
