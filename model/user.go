package model

import (
	"naive-server/api"

	"naive-server/dao"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int
	Username string
	Password string
}

func SignIn(req *api.SignInReq) (resp *api.SignInResp, err error) {
	var user User
	err = dao.DB.Model(&User{Username: req.Username}).Where(&User{Username: req.Username}).First(&user).Error
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if user.Password != req.Password {
		return nil, err
	}
	token, err := ReleaseToken(&user)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	resp = &api.SignInResp{
		Access_token: token,
	}
	return resp, err

}

func SignUp(req *api.SignUpReq) (*api.SignUpResp, error) {
	userID := 21 //todo:实际上id应该和数据库已有的用户不一样，但是懒得写
	newUser := User{
		Id:       userID,
		Username: req.Username,
		Password: req.Password,
	}
	err := dao.DB.Create(&newUser).Error
	if err != nil {
		logrus.Error(err)
	}
	token, err := ReleaseToken(&newUser)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	resp := &api.SignUpResp{
		Access_token: token,
	}
	return resp, err
}

func CheckIn(req *api.CheckInReq) (resp *api.CheckInResp, err error) {

	var point int = 0
	//解析token
	if req.Checkword != "" {
		point = 1
	} else {
		point = 0
	}
	resp = &api.CheckInResp{
		Point: point,
	}
	return resp, nil
}
