package entity

import (
    . "github.com/Danceiny/dict-service/common"
)

type CarLevel int

const (
    _ CarLevel = iota
    KIND
    BRAND
    SERIES
    MODEL
)

type CarEntity struct {
    BaseEntity
    Name  string
    Level NodeLevel
}
