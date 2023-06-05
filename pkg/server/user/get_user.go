package user

import (
	"net/http"

	"github.com/17660472762/miniprogramtest/pkg/results"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetUser(c *gin.Context)(interface{}, results.Error) {
	var User struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
	User.Name = "aaa"
	User.Phone = "12345678910"
	c.JSON(http.StatusOK, User)
	return nil ,nil
}
