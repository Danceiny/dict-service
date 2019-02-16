package service

import (
    "fmt"
    . "github.com/Danceiny/dict-service/persistence/entity"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)

type RepositoryServiceImpl struct {
    DB *gorm.DB
}

func (repo *RepositoryServiceImpl) Add(entity EntityIfc) {
    if repo.DB.NewRecord(entity) {
        now := time.Now().Unix()
        entity.SetCreatedTime(now)
        entity.SetUpdatedTime(now)
        repo.DB.Create(entity)
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
    rows, _ := repo.DB.Raw(fmt.Sprintf("SELECT attr FROM %s WHERE %s = ?", t.GetTableName(), t.GetBidColName()), bid).Rows()
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
    var e = t.InitEmpty()
    if repo.DB.Where(sql, bid).First(e).RecordNotFound() {
        return nil
    } else {
        return e
    }
}

func (*RepositoryServiceImpl) GetCids(t DictTypeEnum, parentBid BID) []BID {
    return nil
}
