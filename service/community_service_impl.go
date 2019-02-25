package service

import (
    . "github.com/Danceiny/dict-service/common"
)

var CommunityServiceImplCpt *CommunityServiceImpl

type CommunityServiceImpl struct {
}

func (*CommunityServiceImpl) BatchQuery(ids []BID, simple bool, p int, c int, onlyId bool) []*CommunityVO {
    return nil
}
