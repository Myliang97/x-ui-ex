package session

import (
	"encoding/gob"
	"x-ui/database/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	loginUser = "LOGIN_USER"
)

func init() {
	gob.Register(model.V2rayUser{})
}

func SetLoginUser(c *gin.Context, user *model.V2rayUser) error {
	s := sessions.Default(c)
	s.Set(loginUser, user)
	return s.Save()
}

func GetLoginUser(c *gin.Context) *model.V2rayUser {
	s := sessions.Default(c)
	obj := s.Get(loginUser)
	if obj == nil {
		return nil
	}
	user := obj.(model.V2rayUser)
	return &user
}

func IsLogin(c *gin.Context) bool {
	return GetLoginUser(c) != nil
}

func ClearSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	s.Save()
}
