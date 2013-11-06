package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["IsLoginPage"] = true
	/* 取url 中的值 */
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		/* -1 代表删除 cookie */
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Ctx.Redirect(301, "/")
		return
	}

	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	/* 获取表单 */
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd") {
		iMaxAge := 0
		if autoLogin {
			iMaxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, iMaxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, iMaxAge, "/")
	}

	this.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value

	return uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd")
}
