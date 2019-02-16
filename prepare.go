package main

import (
    "fmt"
    . "github.com/Danceiny/dict-service/persistence"
    . "github.com/Danceiny/dict-service/service"
    . "github.com/Danceiny/go.utils"
    "github.com/go-redis/redis"
    "github.com/jinzhu/gorm"
    "log"
    "os"
)

var (
    client *redis.Client
    db     *gorm.DB
)

func Prepare() {
    InitEnv()
    InitDB()
    InitRedis()
    ScanComponent()
}

func InitEnv() {
    _ = os.Setenv("REDIS_ADDR", "127.0.0.1:6379")
    _ = os.Setenv("REDIS_PASSWORD", "")
}

func ScanComponent() {
    RepoCpt = &RepositoryServiceImpl{db}
    RedisImplCpt = &RedisImpl{}
    BaseCrudServiceImplCpt = &BaseCrudServiceImpl{RepoCpt}
    BaseCacheServiceImplCpt = &BaseCacheServiceImpl{RedisImplCpt}
}

func InitDB() {
    db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
        GetEnvOrDefault("MYSQL_USERNAME", "root"),
        GetEnvOrDefault("MYSQL_PASSWORD", "root"),
        GetEnvOrDefault("MYSQL_ADDR", "127.0.0.1:3306"),
        GetEnvOrDefault("MYSQL_CHARSET", "utf-8")))
    if err != nil {
        log.Fatalf("connect to mysql failed: %v", err)
    }
    db.LogMode(true)
}

func InitRedis() {
    client = redis.NewClient(&redis.Options{
        Addr:     GetEnvOrDefault("REDIS_ADDR", "127.0.0.1:6379").(string),
        Password: GetEnvOrDefault("REDIS_PASSWORD", "").(string),
        DB:       GetEnvOrDefault("REDIS_DB", 0).(int),
    })
}
