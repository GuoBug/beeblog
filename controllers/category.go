package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsCategory"] = true
	this.TplNames = "category.html"

	op := this.Input().Get("op")
	beego.Debug(op)

	switch op {
	case "add":
		name := this.Input().Get("category")
		beego.Debug(name)
		if len(name) == 0 {
			break
		}
		/* 需要在models 里面处理 */
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return

	case "del":
		id := this.Input().Get("id")
		beego.Debug(id)
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	}

	/* 取值分类中的所以列 Categories 在页面中*/
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
