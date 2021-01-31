/*
* @Time    : 2021-01-31 11:08
* @Author  : CoderCharm
* @File    : user_repository.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package repository

import (
	"github.com/wxyMicro/user/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	//根据用户ID查找用户信息
	FindUserByID(int64) (*model.User, error)
	//创建用户
	CreateUser(*model.User) (int64, error)
	//根据用户id删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*model.User) error
	//查找所有用户
	FindAll() ([]model.User, error)
}

//创建 UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error {
	//return u.mysqlDb.Migrator().CreateTable(&model.User{})
	return u.mysqlDb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
}

//根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

//根据用户ID查找用户信息
func (u UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

//创建用户
func (u UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

//根据用户id删除用户
func (u UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

//更新用户信息
func (u UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Updates(&user).Error
}

//查找所有用户
func (u UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
