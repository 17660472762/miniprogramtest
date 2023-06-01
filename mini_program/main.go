package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()
	//获取信息
	r.GET("/getUser", func(c *gin.Context) {
		var User struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
		}
		User.Name = "aaa"
		User.Phone = "12345678910"
		c.JSON(http.StatusOK, User)
	})
	//接入:1234端口
	r.Run(":1234")
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type WechatConfig struct {
	AppID  string `json:"appid"`
	Secret string `json:"secret"`
}

func getAccessToken(c *gin.Context) (AccessToken, error) {

	return AccessToken{}, nil
}
