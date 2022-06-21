package cache_natsapi

import (
	"fmt"

	"github.com/khanakia/mangobp/mango/nats_util"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/nats-io/nats.go"
	"github.com/ubgo/gofm/cache"
)

// This packaeg is receive request so this will behave more like an api and will be included in cache microservices

const (
	NATS_CACHE_PUT   = "cache.put"
	NATS_CACHE_GET   = "cache.get"
	NATS_CACHE_DEL   = "cache.del"
	NATS_CACHE_FLUSH = "cache.flush"
)

type Config struct {
	Natso natso.Natso
	Cache cache.Cache
}

type CacheNatsApi struct {
	Config
}

type res struct {
	Code string
}

func New(config Config) {
	ec := config.Natso.GetEncodedConn()
	PutSubs(ec, config.Cache)
	GetSubs(ec, config.Cache)
	DelSubs(ec, config.Cache)
	FlushSubs(ec, config.Cache)
}

type CachePutReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Ttl   int    `json:"ttl"`
}

func PutSubs(ec *nats.EncodedConn, cache cache.Cache) {
	ec.Subscribe(NATS_CACHE_PUT, func(subj, reply string, msg CachePutReq) {
		fmt.Println(msg)
		cache.Put(msg.Key, msg.Value, msg.Ttl)
	})
}

type CacheGetReq struct {
	Key string `json:"key"`
}

func GetSubs(ec *nats.EncodedConn, cache cache.Cache) {
	ec.Subscribe(NATS_CACHE_GET, func(subj, reply string, msg CacheGetReq) {
		ec.Publish(reply, nats_util.CreateRespWithData(cache.Get(msg.Key)))
	})
}

type CacheDelReq struct {
	Key string `json:"key"`
}

func DelSubs(ec *nats.EncodedConn, cache cache.Cache) {
	ec.Subscribe(NATS_CACHE_DEL, func(subj, reply string, msg CacheDelReq) {
		cache.Del(msg.Key)
	})
}

func FlushSubs(ec *nats.EncodedConn, cache cache.Cache) {
	ec.Subscribe(NATS_CACHE_FLUSH, func(subj, reply string, msg interface{}) {
		cache.Flush()
	})
}
