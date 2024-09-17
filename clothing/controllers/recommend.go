package controllers

import (
	"github.com/astaxie/beego"
	"clothing/models"
)

type RecommendController struct {
	beego.Controller
}

type UserRecommendResponse struct {
	Code int `json:"code"`
	Data []models.ClothingData `json:"data"`
	Msg string `json:"msg"`
}

func (c *RecommendController) GetUserRecommendList() {
	// id := c.Ctx.Input.Param(":userID") //后续用id来获取数据
        c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")  //允许跨域
	var resp UserRecommendResponse
	data, err := models.QueryClothingRecommendItems()
	if err != nil {
		resp = 	UserRecommendResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
	} else {
		resp = 	UserRecommendResponse{
			Code: 200,
			Data: data,
			Msg: "OK",
		}
	}
	c.Data["json"]  = resp
	// c.Ctx.WriteString("User ID: " + id)
	c.ServeJSON()
}
