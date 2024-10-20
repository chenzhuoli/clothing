package main

import (
	_ "clothing/routers"
	"clothing/controllers"
	"github.com/astaxie/beego"
)

func main() {
	go controllers.PeriodicUpdateAccessToken() 
	beego.Router("/recommend/:userID", &controllers.RecommendController{}, "get:GetUserRecommendList")
	beego.Router("/designer_certification", &controllers.DesignerController{}, "post:PostDesignerCertificationInfo")
	beego.Router("/login", &controllers.LoginController{}, "post:PostLoginInfo")
	beego.Router("/designer_upload_clothing", &controllers.DesignerController{}, "post:PostDesignerClothing")
	beego.Router("/search", &controllers.SearchController{}, "get:GetSearchList")
	beego.Router("/composite_image", &controllers.CompositeImageController{}, "post:GenerateCompositeImage")
	beego.Run()
}

