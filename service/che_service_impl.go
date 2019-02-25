package service

import (
    . "github.com/Danceiny/dict-service/common"
)

var (
    CheServiceImplCpt *CheServiceImpl
)

type CheServiceImpl struct {
}

func (*CheServiceImpl) BatchQuery(ids []BID, simple bool, p int, c int, onlyId bool) []*CheVO {
    return nil
}
