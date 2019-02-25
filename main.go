package main

import (
    . "github.com/Danceiny/dict-service/web"
    log "github.com/sirupsen/logrus"
)

func main() {
    defer clean()
    Prepare()
    if err := Server.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
