package controllers

import (
	"encoding/json"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
	"clothing/common"
)

func UpdateAccessToken() {
    ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancelFunc()
    req_str := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", common.AppID, common.AppSecret)
    fmt.Println(req_str)
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, req_str, nil)
    if err != nil {
        fmt.Println("new http get request failed:", err)
        return
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("http get failed:", err)
        return
    }
    defer resp.Body.Close()
    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read response body failed:", err)
        return
    }
    common.TokenLock.Lock()
    err = json.Unmarshal(respBody, &common.AccessTokenValue)
    if err != nil {
	    fmt.Println("json Unmarshal error:", err)
    } else {
    	fmt.Printf("update access_token success, token=%s, will exire %d seconds", common.AccessTokenValue.AccessToken, common.AccessTokenValue.ExpiresIn)
    }
    common.TokenLock.Unlock()
}

// 周期性的更新AccessToken的值
func PeriodicUpdateAccessToken() {
	UpdateAccessToken()  //初始化第一次执行
	ticker := time.NewTicker(time.Second * time.Duration(common.AccessTokenValue.ExpiresIn - 10))
	for _ = range ticker.C {
		UpdateAccessToken() 
        }
}
