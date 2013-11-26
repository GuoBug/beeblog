package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplNames = "topic.html"

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
	}
	/* 解析表单 */
	title := this.Input().Get("title")
	category := this.Input().Get("category")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")

	var err error

	if len(tid) == 0 {
		err = models.AddTopic(title, content, category)
		if models.CheckCategory(category) {
			beego.Debug("had")
			models.UpdateCategory(category)
		} else {
			beego.Debug("Not ! ")
			models.AddCategory(category)
		}
	} else {
		err = models.ModifyTopic(tid, title, content, category)
		if models.CheckCategory(category) {
			beego.Debug("had")
		} else {
			beego.Debug("Not ! ")
			models.AddCategory(category)
		}
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplNames = "topic_add.html"

}

func (this *TopicController) Modify() {
	this.TplNames = "topic_modify.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
	}
	this.Data["Topic"] = topic
}

func (this *TopicController) Del() {
	if !checkAccount(this.Ctx) {
		beego.Error("delete topic")
		this.Redirect("/login", 302)
		return
	}
	id := this.Input().Get("id")
	err := models.DeleteTopic(id)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) View() {
	this.TplNames = "topic_view.html"
	this.Data["IsTopic"] = true
	/* 自动路由 */
	topic, err := models.GetTopic(this.Ctx.Input.Params("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
}
