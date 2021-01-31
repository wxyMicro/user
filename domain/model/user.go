/*
* @Time    : 2021-01-31 10:58
* @Author  : CoderCharm
* @File    : user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package model

type User struct {
	// 主键id
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	// 用户名称
	UserName  string `gomr:"unique_index;not_null"`
	FirstName string

	// 哈希后的密码
	HashPassword string
}
