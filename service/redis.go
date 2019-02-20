package service

import (
    "github.com/go-redis/redis"
    log "github.com/sirupsen/logrus"
    "time"
)

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
    Client *redis.Client
}

func (impl *RedisImpl) HMGet(key string, fields ...string) []interface{} {
    if fields == nil {
        return nil
    } else if len(fields) == 0 {
        return make([]interface{}, 0, 0)
    }
    if ret, err := impl.Client.HMGet(key, fields...).Result(); err != nil {
        log.Warningf("HMGet err: %v, fields: %v", err, fields)
        return nil
    } else {
        return ret
    }
}

func (impl *RedisImpl) MGet(keys ...string) []interface{} {
    if keys == nil {
        return nil
    } else if len(keys) == 0 {
        return make([]interface{}, 0, 0)
    }
    if ret, err := impl.Client.MGet(keys...).Result(); err != nil {
        log.Warningf("MGet err: %v", err)
        return nil
    } else {
        return ret
    }
}

func (impl *RedisImpl) HMSet(key string, m map[string]interface{}) {
    _, err := impl.Client.HMSet(key, m).Result()
    if err != nil {
        log.Warning(err)
    }
}

func (impl *RedisImpl) MSet(m map[string]interface{}) {
    var args = make([]interface{}, len(m))
    for k, v := range m {
        args = append(args, k, v)
    }
    _, err := impl.Client.MSet(m).Result()
    if err != nil {
        log.Warningf("MSet err: %v", err)
    }
}

func (impl *RedisImpl) Pipelined(f func(pipe Pipeliner) error) []string {
    cmds, err := impl.Client.Pipelined(f)
    if err != nil {
        log.Warningf("pipelined error: %v", err)
    }
    ret := make([]string, len(cmds))
    for i, cmd := range cmds {
        ret[i] = cmd.(*redis.StringCmd).Val()
    }
    return ret
}

func (impl *RedisImpl) Pipeline() Pipeliner {
    return impl.Client.Pipeline()
}

func (impl *RedisImpl) HGet(key string, field string) []byte {
    if ret, err := impl.Client.HGet(key, field).Bytes(); err != nil {
        log.Warningf("redis HGet error: %v", err)
        return nil
    } else {
        return ret
    }
}

func (impl *RedisImpl) Get(key string) []byte {
    if ret, err := impl.Client.Get(key).Bytes(); err != nil {
        log.Warningf("redis Get error: %v", err)
        return nil
    } else {
        return ret
    }
}

func (impl *RedisImpl) Set(key string, bytes []byte, expiration time.Duration) {
    if _, err := impl.Client.Set(key, bytes, expiration).Result(); err != nil {
        log.Warningf("redis Set error: %v", err)
    }
}

func (impl *RedisImpl) HSet(key string, field string, bytes []byte) {
    log.Debugf("【HSet】key: %s, field: %s, bytes: %s", key, field, bytes)
    if _, err := impl.Client.HSet(key, field, bytes).Result(); err != nil {
        log.Warningf("redis HSet error: %v", err)
    }
}

func (impl *RedisImpl) HDel(key string, fields ...string) int64 {
    if fields == nil {
        return 0
    }
    if ret, err := impl.Client.HDel(key, fields...).Result(); err != nil {
        log.Warningf("HGet error: %v", err)
        return 0
    } else {
        return ret
    }
}

func (impl *RedisImpl) Del(keys ...string) int64 {
    if keys == nil {
        return 0
    }
    if ret, err := impl.Client.Del(keys...).Result(); err != nil {
        log.Warningf("Del error: %v", err)
        return 0
    } else {
        return ret
    }
}
