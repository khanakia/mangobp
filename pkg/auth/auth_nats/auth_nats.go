package auth_nats

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/khanakia/mangobp/mango/nats_util"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/khanakia/mangobp/pkg/auth/auth_domain"
	"github.com/khanakia/mangobp/pkg/auth/auth_fn"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cast"
	"github.com/ubgo/gofm/cache"
	"github.com/ubgo/goutil"
	"gorm.io/gorm"
)

const (
	NATS_AUTH_REGISTER_SUBJECT    = "auth.register"
	NATS_AUTH_LOGIN_SUBJECT       = "auth.login"
	NATS_AUTH_FORGOT_PASS_SUBJECT = "auth.forgot_pass"
	NATS_AUTH_RESET_PASS_SUBJECT  = "auth.reset_pass"
)

type Config struct {
	// UserRepo auth_repo.UserRepo
	Natso natso.Natso
	DB    *gorm.DB
	Cache cache.Store
}

type AuthNats struct {
	Config
}

func New(config Config) {
	ec := config.Natso.GetEncodedConn()
	// c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	fmt.Println(ec)
	ec.Subscribe("auth.ping", PingHandler(ec))

	PingSubscriber(ec)

	RegisterSubs(ec, config)
	LoginSubs(ec, config)
	ForgotPassSubs(ec, config)
	ResetPassSubs(ec, config)
}

type res struct {
	Code string
}

// Usual way to define Subsciber with Handler separately - Pros: Can be used with multiple subscribers loosely coupled
func PingHandler(ec *nats.EncodedConn) nats.Handler {
	return func(subj, reply string, msg string) {
		fmt.Println("MSG", msg)
		ec.Publish(reply, &res{
			Code: "200",
		})
	}
}

// Alternate Way to define subscriber - Cons: Tighly Coupled
func PingSubscriber(ec *nats.EncodedConn) {
	ec.Subscribe("auth.ping2", func(subj, reply string, msg string) {
		fmt.Println("MSG", msg)
		ec.Publish(reply, &res{
			Code: "200",
		})
	})
}

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token,omitempty"`
}

// nats request auth.register '{"firstName":"derek"}'
func RegisterSubs(ec *nats.EncodedConn, config Config) {
	ec.Subscribe(NATS_AUTH_REGISTER_SUBJECT, func(subj, reply string, msg UserRequest) {
		fmt.Println("MSG", msg.Name)
		db := config.DB
		// // as msg.YourName is pointer so we need to cast it so it do not throw error in case of nil in ParseName()
		name := cast.ToString(msg.Name)
		person := goutil.ParseName(name)

		usern, err := auth_fn.Register(&auth_domain.User{
			Email:     msg.Email,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Password:  msg.Password,
		}, db)

		if err != nil {
			// ec.Publish(reply, &ErrorResp{
			// 	Code: "400",
			// 	Msg:  err.Error(),
			// })

			ec.Publish(reply, nats_util.CreateRespWithErr(err.Error(), "400", ""))
			return
		}

		// YTD SEND EMAIL

		token, errToken := auth_fn.CreateToken(*usern, db)
		if errToken != nil {
			// ec.Publish(reply, &ErrorResp{
			// 	Code: "token_error",
			// 	Msg:  "server error",
			// })
			ec.Publish(reply, nats_util.CreateRespWithErr(err.Error(), "token_error", "400"))
			return
		}

		lr := &LoginResp{
			Token: token,
		}

		ec.Publish(reply, nats_util.CreateRespWithData(lr))
	})
}

type UserLoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func LoginSubs(ec *nats.EncodedConn, config Config) {
	ec.Subscribe(NATS_AUTH_LOGIN_SUBJECT, func(subj, reply string, msg UserLoginReq) {

		// fmt.Println("MSG", cast.ToString(msg.UserName))

		db := config.DB

		token, err := auth_fn.Login(msg.UserName, msg.Password, db)
		if err != nil {
			ec.Publish(reply, nats_util.CreateRespWithErr(err.Error(), "login_error", "400"))
		}

		ec.Publish(reply, nats_util.CreateRespWithData(gin.H{
			"token": token,
		}))
	})
}

func ForgotPassSubs(ec *nats.EncodedConn, config Config) {
	ec.Subscribe(NATS_AUTH_FORGOT_PASS_SUBJECT, func(subj, reply string, msg UserLoginReq) {
		fmt.Println("MSG", cast.ToString(msg.UserName))
		cnc := config.Cache
		// cnc.Put("test", "dfsd", 10000)
		// cnc.Flush()
		db := config.DB

		err := auth_fn.ForgotPassword(auth_fn.ForgotPasswordRequest{
			UserName: msg.UserName,
		}, cnc, db)
		if err != nil {
			ec.Publish(reply, nats_util.CreateRespWithErr(err.Error(), "login_error", "400"))
		}
	})
}

func ResetPassSubs(ec *nats.EncodedConn, config Config) {
	ec.Subscribe(NATS_AUTH_RESET_PASS_SUBJECT, func(subj, reply string, msg auth_fn.ResetPasswordRequest) {

		cnc := config.Cache
		// cnc.Put("test", "dfsd", 10000)
		// cnc.Flush()
		db := config.DB

		user, err := auth_fn.ResetPassword(msg, cnc, db)
		if err != nil {
			ec.Publish(reply, nats_util.CreateRespWithErr(err.Error(), "", "400"))
		}

		token, errToken := auth_fn.CreateToken(user, db)
		if errToken != nil {
			ec.Publish(reply, nats_util.CreateRespWithErr(err.Error(), "token_error", "400"))
			return
		}

		lr := &LoginResp{
			Token: token,
		}

		ec.Publish(reply, nats_util.CreateRespWithData(lr))
	})
}
