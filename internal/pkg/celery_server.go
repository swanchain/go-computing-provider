package pkg

import (
	"sync"
	"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	"github.com/swanchain/go-computing-provider/conf"
)

var RedisPool *redis.Pool
var CeleryServ *CeleryService
var celeryOnce sync.Once

type CeleryService struct {
	cli *gocelery.CeleryClient
}

func newRedisPool(url string, password string) *redis.Pool {
	RedisPool = &redis.Pool{
		MaxIdle:     5,                 // maximum number of idle connections in the pool
		MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			var conn redis.Conn
			var err error
			if password != "" {
				conn, err = redis.DialURL(url, redis.DialPassword(password))
			} else {
				conn, err = redis.DialURL(url)
			}
			if err != nil {
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return RedisPool
}

func GetRedisClient() redis.Conn {
	newRedisPool(conf.GetConfig().API.RedisUrl, conf.GetConfig().API.RedisPassword)
	return RedisPool.Get()
}

func NewCeleryService() *CeleryService {
	celeryOnce.Do(
		func() {
			RedisPool := newRedisPool(conf.GetConfig().API.RedisUrl, conf.GetConfig().API.RedisPassword)
			celeryClient, err := gocelery.NewCeleryClient(
				gocelery.NewRedisBroker(RedisPool),
				gocelery.NewRedisBackend(RedisPool),
				10)
			if err != nil {
				ulog.Fatalf("Failed init celery service, error: %+v", err)
			}
			CeleryServ = &CeleryService{
				cli: celeryClient,
			}
		})

	return CeleryServ
}

func (s *CeleryService) RegisterTask(taskName string, task interface{}) {
	s.cli.Register(taskName, task)
}

func (s *CeleryService) DelayTask(taskName string, params ...interface{}) (*gocelery.AsyncResult, error) {
	return s.cli.Delay(taskName, params...)
}

func (s *CeleryService) Start() {
	s.cli.StartWorker()
}

func (s *CeleryService) Stop() {
	s.cli.StopWorker()
}
