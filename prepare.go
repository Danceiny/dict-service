package main

import (
    "fmt"
    . "github.com/Danceiny/dict-service/service"
    . "github.com/Danceiny/dict-service/web"
    . "github.com/Danceiny/go.utils"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis"
    "github.com/jinzhu/gorm"
    log "github.com/sirupsen/logrus"
)

func Prepare() {
    InitEnv()
    InitDB()
    InitRedis()
    ScanComponent()
    InitServer()
    CollectPanic()
}

var (
    client      *redis.Client
    db          *gorm.DB
    environment string
    port        string
)

var (
    repositoryServiceImplCpt *RepositoryServiceImpl
    redisImplCpt             *RedisImpl
    idFirewall               *IdFirewallServiceImpl
    treeCacheServiceImplCpt  *TreeCacheServiceImpl
    treeServiceImplCpt       *TreeServiceImpl
    baseCrudServiceImplCpt   *BaseCrudServiceImpl
    baseCacheServiceImplCpt  *BaseCacheServiceImpl
)

func InitServer() {
    var mode string
    if environment == "local" {
        mode = gin.DebugMode
    } else if environment == "dev" || environment == "test" || environment == "qa" {
        mode = gin.TestMode
    } else {
        mode = gin.ReleaseMode
    }
    gin.SetMode(mode)
    Server = gin.Default()
    Routing()
}

func InitEnv() {
    environment = GetEnvOrDefault("env", "local").(string)
    port = GetEnvOrDefault("port", "8080").(string)
}

func ScanComponent() {
    repositoryServiceImplCpt = &RepositoryServiceImpl{db}
    redisImplCpt = &RedisImpl{client}
    idFirewall = &IdFirewallServiceImpl{redisImplCpt}
    baseCacheServiceImplCpt = &BaseCacheServiceImpl{
        redisImplCpt,
    }
    baseCrudServiceImplCpt = &BaseCrudServiceImpl{
        repositoryServiceImplCpt,
        baseCacheServiceImplCpt,
        idFirewall,
    }
    treeCacheServiceImplCpt = &TreeCacheServiceImpl{
        redisImplCpt,
        baseCacheServiceImplCpt,
    }
    treeServiceImplCpt = &TreeServiceImpl{
        repositoryServiceImplCpt,
        baseCacheServiceImplCpt,
        baseCrudServiceImplCpt,
        treeCacheServiceImplCpt}
    ScanService()
}

func ScanService() {
    AreaServiceImplCpt = &AreaServiceImpl{
        repositoryServiceImplCpt,
        treeServiceImplCpt,
        baseCacheServiceImplCpt}
    CategoryServiceImplCpt = &CategoryServiceImpl{}
    CheServiceImplCpt = &CheServiceImpl{}
    CommunityServiceImplCpt = &CommunityServiceImpl{}
    CommonServiceImplCpt = &CommonServiceImpl{
        idFirewall,
        AreaServiceImplCpt,
        CategoryServiceImplCpt,
        CheServiceImplCpt,
        CommunityServiceImplCpt,
    }
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
