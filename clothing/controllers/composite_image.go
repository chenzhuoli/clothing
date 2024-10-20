package controllers

import (
	"github.com/astaxie/beego"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"io"
	"fmt"
	"time"

)

type CompositeImageController struct {
	beego.Controller
}

type CompositeImageRequest struct {
	ClothingPicture string `json:"clothing_picture"`
	UserPicture string `json:"user_picture"`
}

type CompositeImageResponse struct {
	Code int `json:"code"`
	Data []string `json:"data"`
	Msg string `json:"msg"`
}

func (c *CompositeImageController) GenerateCompositeImage() {
        c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")  //允许跨域
	body := c.Ctx.Input.RequestBody
	fmt.Println("received: ", body)
	var req CompositeImageRequest
	var resp CompositeImageResponse
	err := json.Unmarshal(body, &req)
	if err != nil {
		resp = CompositeImageResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	fmt.Println("json: ", req)
	resp = GetCompositeImage(req)
	c.Data["json"]  = resp
	c.ServeJSON()
}

func GetCompositeImage(req CompositeImageRequest) CompositeImageResponse {
	return CompositeImageResponse{
		Code: 200,
		Data: []string{"http://101.42.109.110:8081/7,02d1134eb8"},
		Msg: "OK",
	}
	// TODO: 改成调用算法服务
	req_str := "www.baidu.com"  //TODO: 改成算法服务	
	dataByte, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return CompositeImageResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
	}
	bodyReader := bytes.NewReader(dataByte)
    	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
    	defer cancelFunc()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, req_str, bodyReader)
	if err != nil {
		fmt.Println(err)
		return CompositeImageResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return CompositeImageResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return CompositeImageResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
	}
	fmt.Println(string(body))
	defer resp.Body.Close()
	var compositeImageResp CompositeImageResponse
	err = json.Unmarshal(body, &compositeImageResp)
	return compositeImageResp
}
