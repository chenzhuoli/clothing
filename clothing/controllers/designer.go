package controllers

import (
	"github.com/astaxie/beego"
	"clothing/models"
	"clothing/common"
	"encoding/json"
	"fmt"
)

type DesignerController struct {
	beego.Controller
}

type CertificationResponse struct {
	Code int `json:"code"`
	Data []int `json:"data"`
	Msg string `json:"msg"`
}

func (c *DesignerController) PostDesignerCertificationInfo() {
        c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")  //允许跨域
	var resp CertificationResponse
	body := c.Ctx.Input.RequestBody
	fmt.Println("received: ", body)
	var certification_data common.DesignerCertificationStruct
	err := json.Unmarshal(body, &certification_data)
	if err != nil {
		resp = CertificationResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	fmt.Println("json: ", certification_data)
	err = models.InsertDesignerCertificationData(certification_data)
	if err != nil {
		resp = CertificationResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
	} else {
		resp = CertificationResponse{
			Code: 200,
			Data: nil,
			Msg: "OK",
		}
	}
	c.Data["json"]  = resp
	c.ServeJSON()
}

func (c *DesignerController) PostDesignerClothing() {
        c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")  //允许跨域
	var resp CertificationResponse
	body := c.Ctx.Input.RequestBody
	fmt.Println("received: ", body)
	var total_data common.DesignerUploadClothingTotal
	err := json.Unmarshal(body, &total_data)
	if err != nil {
		resp = CertificationResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	fmt.Println("json: ", total_data)
	/*
	err = models.InsertDesignerClothingData(total_data)
	if err != nil {
		resp = CertificationResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	err = InsertESIndex(total_data)
	if err != nil {
		resp = CertificationResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	*/
	resp = CertificationResponse{
		Code: 200,
		Data: nil,
		Msg: "OK",
	}
	c.Data["json"]  = resp
	c.ServeJSON()
}
