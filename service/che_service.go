package service

import (
    . "github.com/Danceiny/dict-service/common"
)

type CheService interface {
    BatchQuery(ids []BID, simple bool, p int, c int, onlyId bool) []*CheVO
}
