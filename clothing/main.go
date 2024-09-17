package main

import (
	_ "clothing/routers"
	"clothing/controllers"
	"github.com/astaxie/beego"
)

func main() {
	go controllers.PeriodicUpdateAccessToken() 
	beego.Router("/recommend/:userID", &controllers.RecommendController{}, "get:GetUserRecommendList")
	beego.Router("/designer_certification", &controllers.DesignerCertificationController{}, "post:PostDesignerCertificationInfo")
	beego.Router("/login", &controllers.LoginController{}, "post:PostLoginInfo")
	beego.Run()
}

