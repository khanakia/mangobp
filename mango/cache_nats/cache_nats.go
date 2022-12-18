package cache_nats

import (
	"time"

	"github.com/khanakia/mangobp/mango/nats_util"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/nats-io/nats.go"
)

// this is cache client more like RDBMS but as we will be using NATS for communication so this pkg will be called in all the client side microservice which want to consume the cache pkg.

const (
	NATS_CACHE_PUT   = "cache.put"
	NATS_CACHE_GET   = "cache.get"
	NATS_CACHE_DEL   = "cache.del"
	NATS_CACHE_FLUSH = "cache.flush"
)

type Config struct {
	Natso natso.Natso
}

type CacheNats struct {
	Config
	ec *nats.EncodedConn
}

func (pkg CacheNats) Version() string {
	return "0.01"
}

func New(config Config) *CacheNats {
	pkg := CacheNats{Config: config, ec: config.Natso.GetEncodedConn()}
	return &pkg
}

type CachePutReq struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	Ttl   int         `json:"ttl"`
}

// ttl - in seconds
func (a CacheNats) Put(key string, val interface{}, ttl int) (bool, error) {
	err := a.ec.Publish(NATS_CACHE_PUT, CachePutReq{
		Key:   key,
		Value: val,
		Ttl:   ttl,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

type CacheGetReq struct {
	Key string `json:"key"`
}

func (a CacheNats) Get(key string) interface{} {
	var response nats_util.Resp
	err := a.ec.Request(NATS_CACHE_GET, CacheGetReq{Key: key}, &response, 10*time.Millisecond)
	if err != nil {
		return nil
	}
	return response.Data
}

func (a CacheNats) Del(key string) {
	a.ec.Publish(NATS_CACHE_DEL, CacheGetReq{
		Key: key,
	})
}

func (a CacheNats) Flush() {
	a.ec.Publish(NATS_CACHE_FLUSH, "{}")
}
