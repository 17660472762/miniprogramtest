package message

import (
	"net/http"

	"github.com/17660472762/miniprogramtest/pkg/codec"
	"github.com/17660472762/miniprogramtest/pkg/results"
	"github.com/gin-gonic/gin"
)

func (s *Service) UniformSend(ctx *gin.Context) (interface{}, results.Error) {
	var req codec.ReceiveTextMsg
	if err := ctx.BindJSON(&req); err != nil {
		return nil, results.NewError(results.CodeBadRequest, "", "参数错误", err)
	}
	ctx.JSON(http.StatusOK, req)
	return nil, nil
}
