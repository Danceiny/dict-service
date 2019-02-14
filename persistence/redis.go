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

type Redis interface {
    HGet(key string, field string) []byte
    Get(key string) []byte
    HDel(key string, field string) int64
}
type RedisImpl struct {
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
func (impl *RedisImpl) HDel(key string, field string) int64 {
    if ret, err := client.HDel(key, field).Result(); err != nil {
        log.Warningf("HGet error: %v", err)
        return 0
    } else {
        return ret
    }
}
