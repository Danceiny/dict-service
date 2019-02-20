package main

import (
    "fmt"
    . "github.com/Danceiny/dict-service/service"
    . "github.com/Danceiny/go.utils"
    "github.com/go-redis/redis"
    "github.com/jinzhu/gorm"
    log "github.com/sirupsen/logrus"
)

func Prepare() {
    InitEnv()
    InitDB()
    InitRedis()
    ScanComponent()
    CollectPanic()
}

var (
    client *redis.Client
    db     *gorm.DB
)

var (
    repositoryServiceImplCpt *RepositoryServiceImpl
    redisImplCpt             *RedisImpl
    treeCacheServiceImplCpt  *TreeCacheServiceImpl
    treeServiceImplCpt       *TreeServiceImpl
    baseCrudServiceImplCpt   *BaseCrudServiceImpl
    baseCacheServiceImplCpt  *BaseCacheServiceImpl
)

func InitEnv() {
}

func ScanComponent() {
    repositoryServiceImplCpt = &RepositoryServiceImpl{db}
    redisImplCpt = &RedisImpl{client}
    baseCacheServiceImplCpt = &BaseCacheServiceImpl{redisImplCpt}
    baseCrudServiceImplCpt = &BaseCrudServiceImpl{repositoryServiceImplCpt, baseCacheServiceImplCpt}
    treeCacheServiceImplCpt = &TreeCacheServiceImpl{redisImplCpt, baseCacheServiceImplCpt}
    treeServiceImplCpt = &TreeServiceImpl{repositoryServiceImplCpt, baseCrudServiceImplCpt, treeCacheServiceImplCpt}
    ScanService()
}

func ScanService() {
    AreaServiceImplCpt = &AreaServiceImpl{
        repositoryServiceImplCpt,
        treeServiceImplCpt,
        baseCacheServiceImplCpt}
}

func InitDB() {
    var err error
    var connectionInfo = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
        GetEnvOrDefault("MYSQL_USERNAME", "root").(string),
        GetEnvOrDefault("MYSQL_PASSWORD", "root").(string),
        GetEnvOrDefault("MYSQL_ADDR", "127.0.0.1:3306").(string),
        GetEnvOrDefault("MYSQL_DATABASE", "dict").(string),
        GetEnvOrDefault("MYSQL_CHARSET", "utf8").(string))
    db, err = gorm.Open("mysql", connectionInfo)
    if err != nil {
        log.Fatalf("connect to mysql [%s] failed: %v", connectionInfo, err)
    }
    db.LogMode(GetEnvOrDefault("SHOW_SQL", false).(bool))
}

func InitRedis() {
    client = redis.NewClient(&redis.Options{
        Addr:     GetEnvOrDefault("REDIS_ADDR", "127.0.0.1:6379").(string),
        Password: GetEnvOrDefault("REDIS_PASSWORD", "").(string),
        DB:       GetEnvOrDefault("REDIS_DB", 0).(int),
    })
}

func CollectPanic() {
    err := recover()
    if err != nil {
        log.Errorf("Collect Panic: %v", err)
    }
}
