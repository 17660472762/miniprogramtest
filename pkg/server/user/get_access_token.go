package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/17660472762/miniprogramtest/pkg/codec"
	"github.com/17660472762/miniprogramtest/pkg/results"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
)

func (s *Service) GetAccessToken(c *gin.Context) (interface{}, results.Error) {
	var result codec.AccessToken
	body, _ := query.Values(s.wechatconfig)
	request, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token?"+body.Encode(), nil)
	if err != nil {
		return result, results.NewError(results.CodeBadRequest, "", "请求失败", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return result, results.NewError(results.CodeBadRequest, "", "请求失败", err)
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, results.NewError(results.CodeBadRequest, "", "请求失败", err)
	}
	err = json.Unmarshal(resp, &result)
	log.Printf("result:%v", result)
	return result, nil
}
