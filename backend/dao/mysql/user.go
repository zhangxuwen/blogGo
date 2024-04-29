package mysql

import (
	"blogWeb/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"

	"go.uber.org/zap"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "zhangmuyu"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (bool, error) {
	zap.L().Debug("mysql user.go at 16")
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	zap.L().Debug("mysql user.go at 24")
	if count > 0 {
		return true, ErrorUserExist
	}
	return false, nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.PassWord = encryptPassword(user.PassWord)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.PassWord)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.PassWord // 用户登录的密码
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}

	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.PassWord {
		return ErrorInvalidPassword
	}
	// zap.L().Info("password equel is oPassword")
	return
}
