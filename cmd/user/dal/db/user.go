package db

import (
	"context"

	"HuaTug.com/kitex_gen/favorites"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type User struct {
	//gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

const TableName = "users"

/* func (u *User) TableName() string {
	return TableName
} */

/*
User 结构体有一个 TableName 方法，返回的是 "users"。所以当你在 CreateUser 函数中调用 DB.WithContext(ctx).Create(users).Error 时，
GORM 会自动调用 users（即 *User 类型）的 TableName 方法，得到 "users"，然后在 "users" 这个表上执行插入操作。
*/

func CreateUser(ctx context.Context, user *users.User) error {
	if err := DB.WithContext(ctx).Model(&users.User{}).Create(user); err != nil {
		//ToDo:分层处理err
		return errors.Wrapf(err.Error, "CreateUser failed,err: %v", err)
	}
	return nil
}

func CheckUser(ctx context.Context, username, password string) (users.User, error, bool) {

	var user users.User
	var count int64
	if err := DB.WithContext(ctx).Model(&users.User{}).Where("user_name=?", username).Count(&count).Find(&user); err != nil {
		hlog.Info(err)
	}
	if count == 0 {
		logrus.Info("正在创建新用户")
		return user, nil, true
	}
	if err, flag := utils.VerifyPassword(password, user.Password); !flag {
		return user, errors.Wrapf(err, "Password Wrong,err:%v", err), true
	}
	return user, nil, false
}

func CheckUserExistById(ctx context.Context, userId int64) (bool, error) {
	var user users.User
	if err := DB.WithContext(ctx).Where("id=?", userId).Find(&user).Error; err != nil {
		return false, errors.Wrapf(err, "User not exist,err:%v", err)
	}
	if user == (users.User{}) {
		return false, nil
	}
	return true, nil
}
func DeleteUser(ctx context.Context, userId int64) error {
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Delete(&users.User{}); err != nil {
		return errors.Wrapf(err.Error, "Delete user failed,err: %v", err)
	}
	return nil
}

func UpdateUser(ctx context.Context, user *users.User) error {
	if err := DB.WithContext(ctx).Where("user_id=?", user.UserId).Updates(user); err != nil {
		return errors.Wrapf(err.Error, "Update user failed,err: %v", err)
	}
	return nil
}

func QueryUser(ctx context.Context, keyword *string, page, pageSize int64) ([]*users.User, int64, error) {
	db := DB.WithContext(ctx).Model(users.User{}).WithContext(context.Background())
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where("user_name like ?", "%"+*keyword+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, total, errors.Wrapf(err, "QueryUser count failed,err:%v", err)
	}
	var res []*users.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res); err != nil {
		return nil, total, errors.Wrapf(err.Error, "Limit failed,err:%v", err)
	}
	return res, total, nil
}

func GetUser(ctx context.Context, userId int64) (*users.User, error) {
	var user *users.User
	if err := DB.WithContext(ctx).Model(users.User{}).Where("user_id=?", userId).Find(&user); err != nil {
		logrus.Info(err)
		return user, errors.Wrapf(err.Error, "GetUser failed,err:%v", err)
	}
	return user, nil
}

func UserExist(ctx context.Context, uid int64) ([]*favorites.User, error) {
	var user []*favorites.User
	if err := DB.WithContext(ctx).Model(&users.User{}).Where("user_id=?", uid).Find(&user); err != nil {
		return nil, errors.Wrapf(err.Error, "User not exist,err:%v", err)
	}
	return user, nil
}
