package wrap

import (
	"net/http"

	"github.com/17660472762/miniprogramtest/pkg/results"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler func(*gin.Context) (interface{}, results.Error)

func Wrap(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h(c)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    err.Code(),
				"message": err.Message(),
			})
			log.Error().Str("api", c.Request.URL.Path).Err(err).Send()
			return
		}
		c.JSON(http.StatusOK, Ok(resp))
	}
}

func Ok(data interface{}) interface{} {
	return map[string]interface{}{
		"code": results.CodeOk,
		"data": data,
	}
}
