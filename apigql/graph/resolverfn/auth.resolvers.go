package resolverfn

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/khanakia/mangobp/apigql/graph/model"
	"github.com/khanakia/mangobp/pkg/auth/auth_domain"
	"github.com/khanakia/mangobp/pkg/auth/auth_fn"
	"github.com/spf13/cast"
	"github.com/ubgo/goutil"
)

func (r *mutationResolver) AuthRegister(ctx context.Context, input model.RegisterInput) (*model.LoginResponse, error) {
	db := r.GormDB.DB

	yourName := cast.ToString(input.YourName)

	person := goutil.ParseName(yourName)

	usern, err := auth_fn.Register(&auth_domain.User{
		Email:     input.Email,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Company:   cast.ToString(input.Company),
		Password:  input.Password,
	}, db)

	if err != nil {
		return nil, err
	}

	// YTD
	// auth_fn.SendRegisterEmail(*usern, db)

	token, errToken := auth_fn.CreateToken(*usern, db)
	if errToken != nil {
		return nil, errors.New("server error")
	}

	return &model.LoginResponse{
		Token: token,
	}, nil
}

func (r *mutationResolver) AuthLogin(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	// var response nats_util.Resp
	// ec := r.Natso.GetEncodedConn()
	// err := ec.Request("auth.login", UserLoginReq{UserName: input.UserName, Password: input.Password}, &response, 10*time.Second)
	// fmt.Println(err)
	// fmt.Println(response)
	// if response.Error != nil {
	// 	return nil, errors.New(response.Error.Msg)
	// }

	// return nil, errors.New("test")

	// r.LogDbNatsClient.Send(logdb_domain.CreateArgs{
	// 	CreatedAt: time.Now().Add(-time.Hour * 24),
	// 	Flag:      "info",
	// 	Name:      "auth_login",
	// 	Request:   &input,
	// 	ClientIP:  middleware.GetClientIP(ctx),
	// })

	// token, err := auth_fn.Login(input.UserName, input.Password, r.GormDB.DB)

	// if err != nil {
	// 	return nil, err
	// }

	// return &model.LoginResponse{
	// 	Token: token,
	// }, nil

	panic("sdfs")
}

func (r *mutationResolver) AuthForgotPassword(ctx context.Context, userName string) (bool, error) {
	// db := r.GormDB.DB

	// err := auth_fn.ForgotPassword(auth_fn.ForgotPasswordRequest{
	// 	UserName: userName,
	// }, r.CacheNatsClient, db)

	// if err != nil {
	// 	return false, err
	// }

	// return true, nil
	panic("sdfs")
}

func (r *mutationResolver) AuthResetPassword(ctx context.Context, token string, password string) (*model.LoginResponse, error) {
	// db := r.GormDB.DB

	// user, err := auth_fn.ResetPassword(auth_fn.ResetPasswordRequest{
	// 	Token:    token,
	// 	Password: password,
	// }, r.CacheNatsClient, db)

	// if err != nil {
	// 	return nil, err
	// }

	// token, errToken := auth_fn.CreateToken(user, db)
	// if errToken != nil {
	// 	return nil, errors.New("server error")
	// }

	// return &model.LoginResponse{
	// 	Token: token,
	// }, nil
	panic("sdfs")
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type UserLoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
