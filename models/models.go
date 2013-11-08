package models

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME         = "data/beeblog.db"
	_SQLLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id           int64
	Uid          int64
	Title        string
	Content      string `orm:"size(300)"`
	Attachment   string
	Created      time.Time `orm:"index"`
	Updated      time.Time `orm:"index"`
	Views        int64
	Author       string    `orm:"size(50)"`
	ReplyTime    time.Time `orm:"index"`
	ReplyCount   int64
	ReplyLastUId int64
}

func RegisterDB() {
	if IsDir(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	/* 注册 */
	orm.RegisterModel(new(Category), new(Topic))

	/* 注册驱动 */
	orm.RegisterDriver(_SQLLITE3_DRIVER, orm.DR_Sqlite)

	/* 注册默认驱动 必须要 default */
	orm.RegisterDataBase("default", _SQLLITE3_DRIVER, _DB_NAME)
}

func IsDir(Dir string) bool {
	f, e := os.Stat(Dir)
	if e != nil {
		return false
	}

	return f.IsDir()
}

func AddCategory(name string) error {
	/* 初始化 */
	o := orm.NewOrm()

	/* 这是上面的结构 */
	cate := &Category{Title: name}

	/* 查询数据 */
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	fmt.Println("%s,%s  --- addcategory", err, name)

	/* 查询成功,说明有数据 */

	if err == nil {
		beego.Error(err, name)
		return err
	}

	_, err = o.Insert(cate)
	fmt.Println("the err = [%s]", err)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}

	_, err = o.Delete(cate)
	return err

}

func DeleteTopic(id string) error {

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: tid}

	_, err = o.Delete(topic)
	return err
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		beego.Error(err)
	}
	return err
}

func ModifyTopic(tid, title, content string) error {
	o := orm.NewOrm()
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
		return err
	}

	topic := &Topic{Id: id}
	beego.Debug(topic)

	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}

	return err
}

func GetAllTopics(IsDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error

	if IsDesc {
		_, err = qs.OrderBy("created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	/* 初始化 能不能有其他方法 */
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++

	_, err = o.Update(topic)
	return topic, err
}
