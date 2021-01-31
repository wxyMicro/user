/*
* @Time    : 2021-01-31 11:06
* @Author  : CoderCharm
* @File    : user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package handler

import (
	"context"
	"github.com/wxyMicro/user/domain/model"
	"github.com/wxyMicro/user/domain/service"
	user "github.com/wxyMicro/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

//注册
func (u *User) Register(ctx context.Context, UserRegisterRequest *user.UserRegisterRequest, UserRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     UserRegisterRequest.UserName,
		FirstName:    UserRegisterRequest.FirstName,
		HashPassword: UserRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	UserRegisterResponse.Message = "添加成功"
	return nil
}

//登录
func (u *User) Login(ctx context.Context, UserLoginRequest *user.UserLoginRequest, UserLoginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(UserLoginRequest.Username, UserLoginRequest.Pwd)
	if err != nil {
		return err
	}
	UserLoginResponse.IsSuccess = isOk
	return nil
}

//查询用户信息
func (u *User) GetUserInfo(ctx context.Context, UserInfoRequest *user.UserInfoRequest, UserInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(UserInfoRequest.UserName)
	if err != nil {
		return err
	}
	UserInfoResponse = UserForResponse(userInfo)
	return nil
}

//类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
