package main

import (
    . "github.com/Danceiny/dict-service/web"
    log "github.com/sirupsen/logrus"
)

func main() {
    defer clean()
    Prepare()
    if err := Server.Run(); err != nil {
        log.Fatal(err)
    }
    // listen and serve on 0.0.0.0:8080
}
