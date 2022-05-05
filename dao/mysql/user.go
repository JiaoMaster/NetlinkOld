package mysql

import (
	"NetLinkOld/models"
	"NetLinkOld/pkg/jwt"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const secret = "jiaomaster"

func CheckUserExist(UserName string) (bool, error) {
	//1.根据用户名与库中用户名匹配
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, UserName); err != nil {
		fmt.Println(count)
		return false, err
	}
	return count > 0, nil
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Register(user *models.UserInMysql) (ok bool, err error) {
	//1.检查账号是否重复
	username := user.Username
	ok, err = CheckUserExist(username)
	if ok {
		zap.L().Debug("CheckUserExist(userId) fail...", zap.String("DeBug", "账号存在"))
		return false, errors.New("账号存在")
	}
	if !ok {
		zap.L().Debug("CheckUserExist(userId) !ok...", zap.Error(err))
	}
	//2.密码加密
	userPassword := encryptPassword(user.Password)
	user.Password = userPassword
	//3.数据入库
	sqlStr := `insert into user (user_id,username,password,email) values(?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password, user.Username)
	if err != nil {
		return false, err
	}
	return ok, err
}

func Login(user *models.UserInMysql) (string, error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select user_id, username, password from user where username=?`
	err := db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return "", errors.New("用户不存在")
	}
	if err != nil {
		// 查询数据库失败
		return "", errors.New("查询数据库失败")
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return "", errors.New("密码错误")
	}
	token, err := jwt.GenToken(user.UserId, user.Username)
	return token, nil
}

func GetUserInfo(username string) (userinfo *models.User, err error) {
	userinfo = new(models.User)
	sqlStr := `select user_id, username, email from user where username = ?`
	err = db.Get(userinfo, sqlStr, username)
	if err != nil {
		zap.L().Error("db.Get(userinfo,sqlStr,username) err", zap.Error(err))
		return nil, err
	}
	return userinfo, nil
}

func GetUsername(id int64) (username string, err error) {
	sqlStr := `select username from user where user_id = ?`
	type user struct {
		Username string `json:"username" db:"username"`
	}
	u := new(user)
	err = db.Get(u, sqlStr, id)
	if err != nil {
		zap.L().Error("db.Get(userinfo,sqlStr,username) err", zap.Error(err))
		return "nil", err
	}
	return u.Username, nil
}

func PutUserInfo(user *models.User) error {
	fmt.Println(user)
	sqlStr := "update user set email=? where username = ?"
	ret, err := db.Exec(sqlStr, user.Email, user.Username)
	if err != nil {
		zap.L().Error("update failed, err:", zap.Error(err))
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		zap.L().Error("get RowsAffected failed, err:", zap.Error(err))
		return err
	}
	zap.L().Debug("update success...", zap.Int64(" affected rows:", n))
	return nil
}

func PutUserLocation(UserLocation *models.UserLocation) error {
	fmt.Println(UserLocation)
	sqlStr := "update user set location=? where username = ?"
	ret, err := db.Exec(sqlStr, UserLocation.Location, UserLocation.Username)
	if err != nil {
		zap.L().Error("update failed, err:", zap.Error(err))
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		zap.L().Error("get RowsAffected failed, err:", zap.Error(err))
		return err
	}
	zap.L().Debug("update success...", zap.Int64(" affected rows:", n))
	return nil
}

func GetUserLocation(UserLocation *models.UserLocation) (Location float64, err error) {
	sqlStr := `select location from user where username = ?`

	err = db.Get(UserLocation, sqlStr, UserLocation.Username)
	if err != nil {
		zap.L().Error("db.Get(userinfo,sqlStr,username) err", zap.Error(err))
		return 0, err
	}
	Location = UserLocation.Location
	return Location, nil
}
