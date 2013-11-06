package models

import (
	"os"
	"path"
	"time"

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
