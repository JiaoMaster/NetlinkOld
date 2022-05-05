package logic

import (
	"NetLinkOld/dao/mysql"
	"NetLinkOld/models"
	"NetLinkOld/pkg/uuid"
	"fmt"
	"github.com/pkg/errors"
)

func Register(user *models.UserSignUp) error {
	//生成userID
	userId, err := uuid.Getuuid()
	newuser := new(models.UserInMysql)
	newuser.UserId = userId
	newuser.Username = user.Username
	newuser.Password = user.Password
	//入库
	ok, err := mysql.Register(newuser)
	if !ok {
		return err
	}
	//生成token
	return err
}

func Login(data *models.UserSignUp) (string, error) {
	var user *models.UserInMysql
	user = &models.UserInMysql{
		Username: data.Username,
		Password: data.Password,
	}
	//操作数据库校验登陆
	token, err := mysql.Login(user)
	if err != nil {
		if err.Error() == "密码错误" {
			return "", errors.New("密码错误")
		}
		return "", err
	}
	//生成token
	fmt.Println(data.Username)
	if err != nil {
		return "", err
	}
	return token, err
}

func GetUserInfo(username string) (userinfo *models.User, err error) {
	//传入username进行查库操作
	userinfo, err = mysql.GetUserInfo(username)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}

func PutUserInfo(user *models.User) error {
	//对新的UserInfo进行入库操作
	err := mysql.PutUserInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func PutUserLocation(UserLocation *models.UserLocation) error {
	//对新的UserInfo进行入库操作
	err := mysql.PutUserLocation(UserLocation)
	if err != nil {
		return err
	}
	return nil
}

func GetUserLocation(UserLocation *models.UserLocation) (float64, error) {
	//对新的UserInfo进行入库操作
	location, err := mysql.GetUserLocation(UserLocation)
	if err != nil {
		return 0, err
	}
	return location, nil
}
