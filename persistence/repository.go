package persistence

import (
    "fmt"
    . "github.com/Danceiny/dict-service/persistence/entity"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    log "github.com/sirupsen/logrus"
)

var (
    db   *gorm.DB
    Repo *Repository
)

func init() {
    var err error
    // db, err = gorm.Open("mysql", "chaoge:chaoge@tcp(192.168.1.40:3306)/dict?charset=utf8&parseTime=True&loc=Local")
    db, err = gorm.Open("mysql", "root:1996@tcp(192.168.1.13:3306)/dict?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatalf("connect to mysql failed: %v", err)
    }
    db.LogMode(true)
    Repo = &Repository{}
}

func CloseDB() {
    err := db.Close()
    if err != nil {
        log.Fatalf("close db error: %v", err)
    }
}

type Repository struct {
}

func (repo *Repository) Get(t DictTypeEnum, bid BID, simple bool, withTrashed bool) Entity {
    entity := repo.GetEntity(t, bid, withTrashed)
    if entity != nil && !simple {
        bytes := repo.loadAttr(t, bid)
        entity.SetAttr(bytes)
    }
    return entity
}
func (*Repository) loadAttr(t DictTypeEnum, bid BID) (val []byte) {
    rows, _ := db.Raw(fmt.Sprintf("SELECT attr FROM %s WHERE %s = ?", t.GetTableName(), t.GetBidColName()), bid).Rows()
    rows.Next()
    _ = rows.Scan(&val)
    return val
}
func (*Repository) GetEntity(t DictTypeEnum, bid BID, withTrashed bool) Entity {
    var sql string
    if withTrashed {
        sql = fmt.Sprintf("%s = ?", t.GetBidColName())
    } else {
        sql = fmt.Sprintf("%s = ? AND deleted_time = 0", t.GetBidColName())
    }
    if CATEGORY == t {
        var entity CategoryEntity
        if db.Where(sql, bid).First(&entity).RecordNotFound() {
            return nil
        } else {
            return &entity
        }
    } else if AREA == t {
        var entity AreaEntity
        if db.Where(sql, bid).First(&entity).RecordNotFound() {
            return nil
        } else {
            return &entity
        }
    } else if CAR == t {
        var entity CarEntity
        if db.Where(sql, bid).First(&entity).RecordNotFound() {
            return nil
        } else {
            return &entity
        }
    }
    return nil
}
