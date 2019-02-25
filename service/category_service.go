package service

import (
    . "github.com/Danceiny/dict-service/common"
)

type CategoryService interface {
    BatchQuery(ids []BID, simple bool, p int, c int, onlyId bool) []*CategoryVO
}
