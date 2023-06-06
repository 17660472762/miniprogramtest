package message

import (
	"fmt"
	"net/http"

	"github.com/17660472762/miniprogramtest/pkg/codec"
	"github.com/17660472762/miniprogramtest/pkg/results"
	"github.com/gin-gonic/gin"
)

func (s *Service) Receive(ctx *gin.Context) {
	var req codec.ReceiveTextMsg
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("error%v", results.NewError(results.CodeBadRequest, "", "参数错误", err))
	}
	ctx.JSON(http.StatusOK, "")
}
