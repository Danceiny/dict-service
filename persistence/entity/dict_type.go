package entity

type DictTypeEnum int

const (
    _ DictTypeEnum = iota
    CATEGORY
    AREA
    CAR
)

func (t *DictTypeEnum) GetBidColName() string {
    if CATEGORY == *t {
        return "bid"
    } else {
        return "id"
    }
}

func (t *DictTypeEnum) GetTableName() string {
    switch *t {
    case CATEGORY:
        return "dict_category"
    case AREA:
        return "dict_area"
    default:
        return ""
    }
}

func (t *DictTypeEnum) UseHashCache() bool {
    return (*t) == CATEGORY || (*t) == AREA
}

func (t *DictTypeEnum) String() string {
    switch *t {
    case CATEGORY:
        return "CATEGORY";
    case AREA:
        return "AREA";
    case CAR:
        return "CAR"
    }
    return ""
}