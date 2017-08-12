package service

import (
	"database/sql"
	"echodemo/model"
	"echodemo/util"
	"github.com/donnie4w/go-logger/logger"
)

/**
根据用户ID 查询用户信息
 */
func GetUserById(uid int64) (*model.User, error) {
	user := model.User{}
	selectSql := "select * from t_user where id = ?"
	if err := db.Get(&user, selectSql, uid); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err)
		return nil, util.ERROR_DATABASE
	}
	return &user, nil
}


/**
根据用户账号查询用户信息
 */
func GetUserByAccount(account string) (*model.User, error) {
	user := model.User{}
	selectSql := "select * from t_user where account = ?"
	if err := db.Get(&user, selectSql, account); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err)
		return nil, util.ERROR_DATABASE
	}
	return &user, nil
}

/**
插入用户数据
 */
func InserUser(user model.User) (int64, error) {
	insertSql := "INSERT INTO t_user (username,account,password, phone_num,createTime) VALUES (?, ?, ?, ?, ?)"
	result := db.MustExec(insertSql, user.Username, user.Account, user.Password, user.PhoneNum, user.CreateTime)
	id, err := result.LastInsertId()
	if err != nil {
		logger.Debug(err)
		return 0, nil
	}
	return id, nil
}

/**
修改用户密码
 */
func UpdateUserPassword(user *model.User) (int64, error) {
	updateSql := "UPDATE t_user SET password = ?  WHERE id = ?"
	result := db.MustExec(updateSql, user.Password, user.Id)
	return result.RowsAffected()
}

/**
删除用户
 */
func DeleteUser(userId int64) (int64, error) {
	deleteSql := "DELETE FROM t_user WHERE id = ?"
	 result,err := db.Exec( deleteSql, userId)
	if err != nil {
		logger.Debug(err)
		return 0, err
	}
	return result.RowsAffected()
}
