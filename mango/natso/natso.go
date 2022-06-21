package natso

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/khanakia/mangobp/mango/cli"
	"github.com/nats-io/nats.go"
)

const (
	NATS_SUBJ  = "foo"
	NATS_QUEUE = "foo"
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

	opts := []nats.Option{nats.Name("mango app")}
	// opts = setupConnOptions(opts)

	// nc, err := nats.Connect(nats.DefaultURL, nats.ErrorHandler(func(_ *nats.Conn, sub *nats.Subscription, err error) {
	// 	fmt.Println("DSFHSDJFHKJDS", err, sub.Subject)
	// }))

	nc, err := nats.Connect(nats.DefaultURL, opts...)
	fmt.Println("setNats", err)
	if err != nil {
		return
	}
	pkg.natsConn = nc

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	fmt.Println(err)
	if err != nil {
		return
	}
	pkg.natsEncodedConn = ec

	// just a test subscriber
	// _, _ = nc.QueueSubscribe(NATS_SUBJ, NATS_QUEUE, reply) // errors handled below
	// _ = nc.Flush()
	// if err := nc.LastError(); err != nil {
	// 	fmt.Println("SFHSDKJHFJKSDHKJFHDS")
	// 	log.Fatal(err)
	// }

	// log.Printf("Listening on [%s]", NATS_SUBJ)
}

// No need to wait if we are running gin server
func (pkg Natso) Wait() {
	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Block
	log.Println()
	log.Printf("Draining...")
	_ = pkg.natsConn.Drain()
	log.Fatalf("Exiting")
}

// Worker function for the backend
func reply(msg *nats.Msg) {
	fmt.Println("Received: ", string(msg.Data))
	// fn := string(msg.Data)
	// content, err := ioutil.ReadFile(fn)
	// if err != nil {
	// 	log.Printf("error reading %s: %s", fn, err)
	// }
	// if len(content) >= 1024*1024 {
	// 	// Max. NATS message size
	// 	content = content[:1024*1024]
	// }
	_ = msg.Respond([]byte("Response: Foo Bar"))
}

func (pkg Natso) GetConn() *nats.Conn {
	return pkg.natsConn
}

func (pkg Natso) GetEncodedConn() *nats.EncodedConn {
	return pkg.natsEncodedConn
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
