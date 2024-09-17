package common

import (
	"sync"
)


var AppID = "wx312c36de9d0cd635"
var AppSecret = "1342255806694beefc6965527da2cd2f"

type DesignerCertificationStruct struct {
	UserID string `json:"userid"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	ClothingWorks []string `json:"clothing_works"`
}

type AccessTokenStruct struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}
var AccessTokenValue AccessTokenStruct
var TokenLock sync.RWMutex

