package tool

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis_ctx struct {
	cli         redis.Conn
	pool        *redis.Pool
	max_idle    int
	max_active  int
	host        string
	timeout     int
	idleTimeout int
}

func NewRedis(MaxIdle, maxActice, TimeOut, IdleTimeout int, Host string) Redis_ctx {
	r := Redis_ctx{
		max_idle:    maxActice,
		max_active:  maxActice,
		timeout:     TimeOut,
		idleTimeout: IdleTimeout,
		host:        Host,
	}
	return r
}
func (ctx *Redis_ctx) Init() {
	timeout := time.Duration(ctx.timeout) * time.Second
	ctx.pool = &redis.Pool{
		MaxIdle:     ctx.max_idle,
		MaxActive:   ctx.max_active,
		IdleTimeout: time.Duration(ctx.timeout) + time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			var err error
			for i := 0; i < 3; i++ {
				_, err = c.Do("PING")
				if err == nil {
					break
				}
				time.Sleep(100 * time.Millisecond)
			}
			return err
		},
		//ops :=[]redis.DialOption{

		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp",ctx.host,)
			c, err := redis.DialTimeout("tcp", ctx.host, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

}

func (ctx *Redis_ctx) SET(key string, value interface{}) error {
	var cli redis.Conn
	cli = ctx.pool.Get()
	defer cli.Close()

	startTime := time.Now()

	s, err := redis.String(cli.Do("SET", key, value))
	fmt.Println(s)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	endTime := time.Now()
	log.Println("SSET take time", endTime.Sub(startTime))
	return nil
}
func (ctx *Redis_ctx) GET(key string) (value string, err error) {
	var cli redis.Conn
	cli = ctx.pool.Get()
	defer cli.Close()

	startTime := time.Now()

	s, err := redis.String(cli.Do("GET", key))
	fmt.Println(s)
	if err != nil {
		log.Printf("%v\n", err)
		return "", err
	}
	endTime := time.Now()
	log.Println("GET take time", endTime.Sub(startTime))
	return s, nil
}

// func main() {
// 	red := Redis_ctx{
// 		max_idle:    5,
// 		max_active:  5,
// 		idleTimeout: 5,
// 		host:        "127.0.0.1:6379",
// 		timeout:     5,
// 	}
// 	red.Init()
// 	err := red.SADD("dddd", "d")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	value, err2 := red.SMEMBERS("dddd")
// 	if err2 != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(value)
// }

func (ctx *Redis_ctx) SADD(key string, value interface{}) error {
	var cli redis.Conn
	cli = ctx.pool.Get()
	defer cli.Close()
	startTime := time.Now()

	s, err := redis.String(cli.Do("SADD", key, value))
	fmt.Println(s)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	endTime := time.Now()
	log.Println("SET take time", endTime.Sub(startTime))
	return nil
}

func (ctx *Redis_ctx) SMEMBERS(key string) (value []string, err error) {
	var cli redis.Conn
	cli = ctx.pool.Get()
	defer cli.Close()

	startTime := time.Now()

	s, err := redis.Strings(cli.Do("SMEMBERS", key))
	fmt.Println(s)
	if err != nil {
		log.Printf("%v\n", err)
		return []string{}, err
	}
	endTime := time.Now()
	log.Println("SMEMBERS take time", endTime.Sub(startTime))
	return s, nil
}

var Red Redis_ctx

func init() {
	Red = NewRedis(5, 5, 5, 5, "127.0.0.1:6379")
	Red.Init()

}
