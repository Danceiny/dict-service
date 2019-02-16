package main

import (
    log "github.com/sirupsen/logrus"
)

func clean() {
    err := db.Close()
    if err != nil {
        log.Fatalf("close DB error: %v", err)
    }
}
