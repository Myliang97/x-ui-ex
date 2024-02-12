package service

import (
	"errors"
	"x-ui/database"
	"x-ui/database/model"
	"x-ui/logger"

	"gorm.io/gorm"
)

type UserService struct {
}

func (s *UserService) GetFirstUser() (*model.V2rayUser, error) {
	db := database.GetDB()

	user := &model.V2rayUser{}
	err := db.Model(model.V2rayUser{}).
		First(user).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CheckUser(username string, password string) *model.V2rayUser {
	db := database.GetDB()

	user := &model.V2rayUser{}
	err := db.Model(model.V2rayUser{}).
		Where("username = ? and password = ?", username, password).
		First(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		logger.Warning("check user err:", err)
		return nil
	}
	return user
}

func (s *UserService) UpdateUser(id int, username string, password string) error {
	db := database.GetDB()
	return db.Model(model.V2rayUser{}).
		Where("id = ?", id).
		Update("username", username).
		Update("password", password).
		Error
}

func (s *UserService) UpdateFirstUser(username string, password string) error {
	if username == "" {
		return errors.New("username can not be empty")
	} else if password == "" {
		return errors.New("password can not be empty")
	}
	db := database.GetDB()
	user := &model.V2rayUser{}
	err := db.Model(model.V2rayUser{}).First(user).Error
	if database.IsNotFound(err) {
		user.Username = username
		user.Password = password
		return db.Model(model.V2rayUser{}).Create(user).Error
	} else if err != nil {
		return err
	}
	user.Username = username
	user.Password = password
	return db.Save(user).Error
}
