package logdb_nats_client

import (
	"github.com/khanakia/mangobp/mango/logdb/logdb_domain"
	"github.com/nats-io/nats.go"
)

// type Config struct {
// 	Natso natso.Natso
// }

// type LogDbNatsClient struct {
// 	Config
// 	ec *nats.EncodedConn
// }

// func (pkg LogDbNatsClient) Version() string {
// 	return "0.01"
// }

// func New(config Config) LogDbNatsClient {
// 	pkg := LogDbNatsClient{Config: config, ec: config.Natso.GetEncodedConn()}
// 	return pkg
// }

// // ttl - in seconds
// func (a LogDbNatsClient) Send(args logdb_domain.CreateArgs) (bool, error) {
// 	err := a.ec.Publish(logdb_domain.NATS_LOGDB_LOG, args)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

func Send(args logdb_domain.CreateArgs, ec *nats.EncodedConn) (bool, error) {
	err := ec.Publish(logdb_domain.NATS_LOGDB_LOG, args)
	if err != nil {
		return false, err
	}
	return true, nil
}
