package service

import (
	. "github.com/Danceiny/dict-service/common"
)

type AreaService interface {
	GetArea(id INT, p, c int) *AreaVO
}
