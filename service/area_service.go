package service

import (
	. "github.com/Danceiny/dict-service/common"
)

type AreaService interface {
	GetArea(id NodeId, p, c int) *AreaVO
}
