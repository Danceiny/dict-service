package service

import (
    . "github.com/Danceiny/dict-service/common"
)

type CommunityService interface {
    BatchQuery(ids []BID, simple bool, p int, c int, onlyId bool) []*CommunityVO
}
