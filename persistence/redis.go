package persistence

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	client       *redis.Client
	RedisImplCpt *RedisImpl
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.13:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisImplCpt = &RedisImpl{}
}

type Pipeliner = redis.Pipeliner

type Redis interface {
	HGet(key string, field string) []byte
	Get(key string) []byte
	Set(key string, bytes []byte, expiration time.Duration)
	HDel(key string, fields ...string) int64
	HSet(key string, field string, bytes []byte)
	Pipeline() Pipeliner
	Pipelined(func(pipe Pipeliner) error) []string
	Del(key ...string) int64
	HMSet(key string, m map[string]interface{})
	HMGet(key string, fields ...string) []interface{}
	MGet(keys ...string) []interface{}
	MSet(m map[string]interface{})
}
type RedisImpl struct {
}

func (impl *RedisImpl) HMGet(key string, fields ...string) []interface{} {
	if ret, err := client.HMGet(key, fields...).Result(); err != nil {
		log.Warning(err)
		return nil
	} else {
		return ret
	}
}

func (impl *RedisImpl) MGet(keys ...string) []interface{} {
	if ret, err := client.MGet(keys...).Result(); err != nil {
		log.Warning(err)
		return nil
	} else {
		return ret
	}
}
func (impl *RedisImpl) HMSet(key string, m map[string]interface{}) {
	_, err := client.HMSet(key, m).Result()
	if err != nil {
		log.Warning(err)
	}
}
func (impl *RedisImpl) MSet(m map[string]interface{}) {
	var args = make([]interface{}, len(m))
	for k, v := range m {
		args = append(args, k, v)
	}
	_, err := client.MSet(m).Result()
	if err != nil {
		log.Warning(err)
	}
}
func (impl *RedisImpl) Pipelined(f func(pipe Pipeliner) error) []string {
	cmds, err := client.Pipelined(f)
	if err != nil {
		log.Warning(err)
		return nil
	} else {
		ret := make([]string, len(cmds))
		for i, cmd := range cmds {
			ret[i] = cmd.(*redis.StringCmd).Val()
		}
		return ret
	}
}

func (impl *RedisImpl) Pipeline() Pipeliner {
	return client.Pipeline()
}

func (impl *RedisImpl) HGet(key string, field string) []byte {
	if ret, err := client.HGet(key, field).Bytes(); err != nil {
		log.Warningf("redis HGet error: %v", err)
		return nil
	} else {
		return ret
	}
}
func (impl *RedisImpl) Get(key string) []byte {
	if ret, err := client.Get(key).Bytes(); err != nil {
		log.Warningf("redis Get error: %v", err)
		return nil
	} else {
		return ret
	}
}
func (impl *RedisImpl) Set(key string, bytes []byte, expiration time.Duration) {
	if _, err := client.Set(key, bytes, expiration).Result(); err != nil {
		log.Warningf("redis Set error: %v", err)
	}
}
func (impl *RedisImpl) HSet(key string, field string, bytes []byte) {
	log.Infof("【HSet】key: %s, field: %s, bytes: %s", key, field, bytes)
	if _, err := client.HSet(key, field, bytes).Result(); err != nil {
		log.Warningf("redis HSet error: %v", err)
	}
}
func (impl *RedisImpl) HDel(key string, fields ...string) int64 {
	if ret, err := client.HDel(key, fields...).Result(); err != nil {
		log.Warningf("HGet error: %v", err)
		return 0
	} else {
		return ret
	}
}

func (impl *RedisImpl) Del(key ...string) int64 {
	if ret, err := client.Del(key...).Result(); err != nil {
		log.Warningf("Del error: %v", err)
		return 0
	} else {
		return ret
	}
}
