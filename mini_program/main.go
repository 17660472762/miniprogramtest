package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
)

func main() {
	log.Print("微信云托管服务启动成功")
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
	r.GET("/accessToken", func(c *gin.Context) {
		accessToken, err := getAccessToken(c)
		if err != nil {
			log.Printf("err%v", err)
		}
		c.JSON(http.StatusOK, accessToken)
	})
	r.Run()
}

type AccessToken struct {
	AccessToken string `json:"access_token" form:"access_token" url:"access_token"`
	ExpiresIn   int64  `json:"expires_in" form:"expires_in" url:"expires_in"`
}

type WechatConfig struct {
	AppID  string `json:"appid"`
	Secret string `json:"secret"`
}

type AccessTokenParam struct {
	GrantType string `json:"grant_type" form:"grant_type" url:"grant_type"`
	AppID     string `json:"appid" form:"appid" url:"appid"`
	Secret    string `json:"secret" form:"secret" url:"secret"`
}

func getAccessToken(c *gin.Context) (AccessToken, error) {
	var result AccessToken
	config := AccessTokenParam{
		GrantType: "client_credential",
		AppID:     "wxde659347f55b799b",
		Secret:    "9a022075538836978d7b3003b72bfcfd",
	}
	body, _ := query.Values(config)
	request, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token?"+body.Encode(), nil)
	if err != nil {
		return result, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(resp, &result)
	log.Printf("result:%v", result)
	return result, err
}
