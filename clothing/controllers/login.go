package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"io"
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"clothing/common"
)

type LoginController struct {
	beego.Controller
}

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Code int `json:"code"`
	Data []WechatGetPhoneNumberResponse `json:"data"`
	Msg string `json:"msg"`
}

type WatermarkStruct struct {
	Timestamp int `json:"timestamp"`
	Appid string `json:"appid"`
}

type PhoneInfoStruct struct {
	PhoneNumber string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode string `json:"countryCode"`
	Watermark WatermarkStruct `json:"watermark"`
}

type WechatGetPhoneNumberResponse struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	PhoneInfo PhoneInfoStruct `json:"phone_info"`
}

func (c *LoginController) PostLoginInfo() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")  //允许跨域
	var resp LoginResponse
	body := c.Ctx.Input.RequestBody
	fmt.Println("received: ", body)
	var loginData LoginRequest
	err := json.Unmarshal(body, &loginData)
	if err != nil {
		resp = LoginResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	fmt.Println("json: ", loginData)
	phoneResp, err := GetPhoneByCode(loginData.Code)
	code := phoneResp.ErrCode //默认
	if (code == 0) {
		code = 200  //成功
	}
	var data []WechatGetPhoneNumberResponse
	if (phoneResp != nil) {
		data = append(data, *phoneResp)
	}
	c.Data["json"]  = LoginResponse{
		Code: code,
		Msg: phoneResp.ErrMsg,
		Data: data,
	}
	c.ServeJSON()
	return
}

func GetPhoneByCode(code string) (*WechatGetPhoneNumberResponse, error) {
	common.TokenLock.Lock()
	// 初始化时token为空的情况处理
	if (len(common.AccessTokenValue.AccessToken) == 0) {
		UpdateAccessToken()
	}
	req_str := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", common.AccessTokenValue.AccessToken)
	fmt.Println(req_str)
	common.TokenLock.Unlock()
	loginReq := LoginRequest{
		Code: code,
	}
	dataByte, err := json.Marshal(loginReq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	bodyReader := bytes.NewReader(dataByte)
    	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
    	defer cancelFunc()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, req_str, bodyReader)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	defer resp.Body.Close()
	var phoneResp WechatGetPhoneNumberResponse
	err = json.Unmarshal(body, &phoneResp)
	return &phoneResp, err
}
