package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)

	category := this.Input().Get("cate")

	var err error
	topics := make([]*models.Topic, 0)

	if len(category) == 0 {
		topics, err = models.GetAllTopics(true)
	} else {
		topics, err = models.GetTopicsByCategory(category, true)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	Category, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"] = Category

}
