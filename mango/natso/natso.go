package natso

import (
	"fmt"

	"github.com/khanakia/mangobp/mango/cli"
	"github.com/nats-io/nats.go"
)

type Natso struct {
	Config
	natsConn        *nats.Conn
	natsEncodedConn *nats.EncodedConn
}

func (pkg Natso) Version() string {
	return "0.01"
}

type Config struct {
	Cli cli.Cli
}

func New(config Config) Natso {
	pkg := Natso{
		Config: config,
	}

	setNats(&pkg)

	return pkg
}

func setNats(pkg *Natso) {
	nc, err := nats.Connect(nats.DefaultURL, nats.ErrorHandler(func(_ *nats.Conn, sub *nats.Subscription, err error) {
		fmt.Println("DSFHSDJFHKJDS", err, sub.Subject)

	}))
	fmt.Println(err)
	if err != nil {
		return
	}
	pkg.natsConn = nc
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	fmt.Println(err)
	if err != nil {
		return
	}
	pkg.natsEncodedConn = c
}

func (pkg Natso) GetConn() *nats.Conn {
	return pkg.natsConn
}

func (pkg Natso) GetEncodedConn() *nats.EncodedConn {
	return pkg.natsEncodedConn
}
