package message

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/17660472762/miniprogramtest/pkg/codec"
	"github.com/17660472762/miniprogramtest/pkg/results"
	"github.com/gin-gonic/gin"
)

func (s *Service) Receive(ctx *gin.Context) {
	var req codec.ReceiveTextMsg
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("error%+v", results.NewError(results.CodeBadRequest, "", "参数错误", err))
	}
	openid := ctx.Request.Header.Get("X-WX-OPENID")
	log.Printf("openid:%v", openid)
	log.Printf("text%+v", req)
	process(ctx, &req, openid)

	ctx.JSON(http.StatusOK, codec.TransferMsg{
		ToUserName:   openid,
		FromUserName: req.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "transfer_customer_service",
	})
}

func process(ctx context.Context, req *codec.ReceiveTextMsg, openid string) {
	switch req.MsgType {
	case codec.Text:
		{
			data, err := json.Marshal(codec.TextMsg{
				ToUser:  openid,
				MsgType: codec.Text,
				Text: struct {
					Content string "json:\"content\" form:\"content\" url:\"content\""
				}{Content: fmt.Sprintf("收到消息:%v", req.Content)},
			})
			if err != nil {
				log.Printf("send err:%+v", err)
			}
			request, err := http.NewRequest(http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/custom/send?", bytes.NewReader(data))
			request.Header.Set("Content-Type", "application/json")

			if err != nil {
				log.Printf("send request err:%+v", err)
			}
			var result codec.SendResult
			client := &http.Client{}
			response, err := client.Do(request)
			log.Printf("response:%+v", response)
			if err != nil {
				log.Printf("send request err:%+v", err)
			}
			defer response.Body.Close()
			resp, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("send request err:%+v", err)
			}
			log.Printf("resp:%+v", string(resp))

			err = json.Unmarshal(resp, &result)
			if err != nil {
				log.Printf("send request err:%+v", err)
			}
			log.Printf("result:%+v", result)
		}
	case codec.Image:
		{
			// TODO:
		}
	case codec.Link:
		{
			// TODO:
		}
	case codec.MiniProgramPage:
		{
			// TODO:
		}
	}

}
