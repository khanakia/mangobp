package auth_nats

import (
	"fmt"

	"github.com/khanakia/mangobp/mango/natso"
	"github.com/khanakia/mangobp/pkg/auth/auth_repo"
	"github.com/nats-io/nats.go"
)

type AuthNats struct {
	Config
}

type Config struct {
	UserRepo auth_repo.UserRepo
	Natso    natso.Natso
}

type res struct {
	Code string
}

func New(config Config) {
	ec := config.Natso.GetEncodedConn()
	// c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	ec.Subscribe("auth.ping", PingHandler(ec))

	PingSubscriber(ec)

	AuthRegister(ec, config.UserRepo)
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

type ErrorResp struct {
	Msg  string `json:"msg,omitempty"`
	Code string `json:"code,omitempty"`
}

type UserRequest struct {
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Password string  `json:"password"`
}

type LoginResp struct {
	Token string `json:"token,omitempty"`
}
type Resp struct {
	Data  interface{} `json:"data,omitempty"`
	Error *ErrorResp  `json:"error,omitempty"`
}

// nats request auth.register '{"firstName":"derek"}'
func AuthRegister(ec *nats.EncodedConn, repo auth_repo.UserRepo) {
	ec.Subscribe("auth.register", func(subj, reply string, msg *UserRequest) {
		fmt.Println("MSG", msg)

		// // as msg.YourName is pointer so we need to cast it so it do not through error in case of nil in ParseName()
		// name := cast.ToString(msg.Name)
		// person := goutil.ParseName(name)

		// _, err := repo.Register(&auth_domain.User{
		// 	Email:     msg.Email,
		// 	FirstName: person.FirstName,
		// 	LastName:  person.LastName,
		// 	Password:  msg.Password,
		// })

		// if err != nil {
		// 	ec.Publish(reply, &ErrorResp{
		// 		Code: "400",
		// 		Msg:  err.Error(),
		// 	})
		// 	return
		// }

		// lr := Resp{
		// 	Error: &ErrorResp{
		// 		Msg: "Test",
		// 	},
		// }

		// ec.Publish(reply, lr)
	})
}

type UserLoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func AuthLogin(ec *nats.EncodedConn, repo auth_repo.UserRepo) {
	ec.Subscribe("auth.login", func(subj, reply string, msg *UserLoginReq) {
		fmt.Println("MSG", msg)

		// // as msg.YourName is pointer so we need to cast it so it do not through error in case of nil in ParseName()
		// name := cast.ToString(msg.Name)
		// person := goutil.ParseName(name)

		// _, err := repo.Register(&auth_domain.User{
		// 	Email:     msg.Email,
		// 	FirstName: person.FirstName,
		// 	LastName:  person.LastName,
		// 	Password:  msg.Password,
		// })

		// if err != nil {
		// 	ec.Publish(reply, &ErrorResp{
		// 		Code: "400",
		// 		Msg:  err.Error(),
		// 	})
		// 	return
		// }

		lr := Resp{
			Error: &ErrorResp{
				Msg: "Test",
			},
		}

		ec.Publish(reply, lr)
	})
}
