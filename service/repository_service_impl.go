package service

import (
	"fmt"
	. "github.com/Danceiny/dict-service/persistence/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	RepoCpt *RepositoryServiceImpl
)

func CloseDB() {
	err := RepoCpt.db.Close()
	if err != nil {
		log.Fatalf("close db error: %v", err)
	}
}

type RepositoryServiceImpl struct {
	db *gorm.DB
}

func (repo *RepositoryServiceImpl) Add(entity EntityIfc) {
	if repo.db.NewRecord(entity) {
		now := time.Now().Unix()
		entity.SetCreatedTime(now)
		entity.SetUpdatedTime(now)
		repo.db.Create(entity)
	}
}

func (repo *RepositoryServiceImpl) Get(t DictTypeEnum, bid BID, simple bool, withTrashed bool) EntityIfc {
	entity := repo.GetEntity(t, bid, withTrashed)
	if entity != nil && !simple {
		bytes := repo.loadAttr(t, bid)
		entity.SetAttr(bytes)
	}
	return entity
}
func (repo *RepositoryServiceImpl) loadAttr(t DictTypeEnum, bid BID) (val []byte) {
	rows, _ := repo.db.Raw(fmt.Sprintf("SELECT attr FROM %s WHERE %s = ?", t.GetTableName(), t.GetBidColName()), bid).Rows()
	rows.Next()
	_ = rows.Scan(&val)
	return val
}
func (repo *RepositoryServiceImpl) GetEntity(t DictTypeEnum, bid BID, withTrashed bool) EntityIfc {
	var sql string
	if withTrashed {
		sql = fmt.Sprintf("%s = ?", t.GetBidColName())
	} else {
		sql = fmt.Sprintf("%s = ? AND deleted_time = 0", t.GetBidColName())
	}
	if CATEGORY == t {
		var entity CategoryEntity
		if repo.db.Where(sql, bid).First(&entity).RecordNotFound() {
			return nil
		} else {
			return &entity
		}
	} else if AREA == t {
		var entity AreaEntity
		if repo.db.Where(sql, bid).First(&entity).RecordNotFound() {
			return nil
		} else {
			return &entity
		}
	} else if CAR == t {
		var entity CarEntity
		if repo.db.Where(sql, bid).First(&entity).RecordNotFound() {
			return nil
		} else {
			return &entity
		}
	}
	return nil
}

func (*RepositoryServiceImpl) GetCids(t DictTypeEnum, parentBid BID) []BID {
	return nil
}
