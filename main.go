package main

import (
    "github.com/Danceiny/dict-service/api"
    "github.com/Danceiny/dict-service/persistence"
    log "github.com/sirupsen/logrus"
)

func main() {
    defer clean()
    if err := api.Server.Run(); err != nil {
        log.Fatal(err)
    }
    // listen and serve on 0.0.0.0:8080
}

func clean() {
    persistence.CloseDB()
}
