package service

import (
    "fmt"
    . "github.com/Danceiny/dict-service/api"
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
    . "github.com/Danceiny/go.fastjson"
    "github.com/sirupsen/logrus"
)

const (
    AREA_NAME      = "area"
    CAR_NAME       = "car"
    CATEGORY_NAME  = "category"
    COMMUNITY_NAME = "community"
    UNKNOWN_NAME   = "unknown"
)

var (
    EMPTY_ARRAY_JSON     = [0]int{}
    CommonServiceImplCpt *CommonServiceImpl
)

type CommonServiceImpl struct {
    IdFirewall    IdFirewallService
    AreaServ      AreaService
    CategoryServ  CategoryService
    CheServ       CheService
    CommunityServ CommunityService
}

type DictTypeIdStruct struct {
    TypeEnum DictTypeEnum
    Bid      BID
}
type IdChecker struct {
    validIdList        []BID
    invalidIdIndexList []int
    validIdIndexList   []int
}

func (impl *CommonServiceImpl) Get(gid string, p int, c int) *JSONObject {
    fmt.Println(impl)
    return nil
}

func (impl *CommonServiceImpl) extractUnknownGids([]STRING) []*DictTypeIdStruct {
    return nil
}

func (impl *CommonServiceImpl) newIntIdChecker(t DictTypeEnum, originIds []INT) *IdChecker {
    var checker = &IdChecker{
        make([]BID, 0),
        make([]int, 0),
        make([]int, 0),
    }
    var l = len(originIds)
    for j := 0; j < l; j++ {
        if impl.IdFirewall.ValidateId(t, originIds[j]) {
            checker.validIdIndexList = append(checker.validIdIndexList, j)
            checker.validIdList = append(checker.validIdList, originIds[j])
        } else {
            checker.invalidIdIndexList = append(checker.invalidIdIndexList, j)
        }
    }
    return checker
}
func (impl *CommonServiceImpl) newStringIdChecker(t DictTypeEnum, originIds []STRING) *IdChecker {
    var checker = &IdChecker{
        make([]BID, 0),
        make([]int, 0),
        make([]int, 0),
    }
    var l = len(originIds)
    for j := 0; j < l; j++ {
        var bid = STRING(originIds[j])
        if impl.IdFirewall.ValidateId(t, bid) {
            checker.validIdIndexList = append(checker.validIdIndexList, j)
            checker.validIdList = append(checker.validIdList, bid)
        } else {
            checker.invalidIdIndexList = append(checker.invalidIdIndexList, j)
        }
    }
    return checker
}
func (impl *CommonServiceImpl) MultiGet(req *MultiGetQueryReq) *JSONObject {
    // Start 处理参数
    var areaIds = req.Area
    var carIds = req.Car
    var categoryIds = req.Category
    var communityIds = req.Community
    var unknownIds = req.Unknown

    var parentDepth = req.ParentDepth
    var childrenDepth = req.HasChildren
    var onlyId = req.OnlyId
    logrus.Infof("req: %v %v %v %v %v %v %v %v", areaIds, carIds, categoryIds, communityIds, unknownIds, parentDepth, childrenDepth, onlyId)
    // ids为null时将size设为-1，作为一个特殊标记以供后面判断
    var carSize int
    if cap(carIds) == 0 {
        carSize = -1
    } else {
        carSize = len(carIds)
    }
    var communitySize int
    if cap(communityIds) == 0 {
        communitySize = -1
    } else {
        communitySize = len(communityIds)
    }
    // Start unknown 特殊处理
    var unknownSize int
    if cap(unknownIds) == 0 {
        unknownSize = 0
    } else {
        unknownSize = len(unknownIds)
    }
    // indexMarker, unknownList中每个id对应indexMarker数组中的一个元素，该元素是一个长度为2的数组，
    // 该数组中第一个值为DictTypeEnum的ordinal值，第二个值为该unknown-id在该DictTypeEnum所对应的查询结果数组中的索引
    var indexMarker = make([][2]int, unknownSize)
    if unknownSize > 0 {
        var typeIdStructs = impl.extractUnknownGids(unknownIds)
        var l = len(typeIdStructs)
        for i := 0; i < l; i++ {
            var typeIdStruct = typeIdStructs[i];
            var typeEnum = typeIdStruct.TypeEnum
            if CAR == typeEnum {
                carIds = append(carIds, typeIdStruct.Bid.(INT))
                indexMarker[i][0] = int(CAR)
                indexMarker[i][1] = len(carIds) - 1
            } else if COMMUNITY == typeEnum {
                communityIds = append(communityIds, typeIdStruct.Bid.(INT))
                indexMarker[i][0] = int(COMMUNITY)
                indexMarker[i][1] = len(communityIds) - 1
            }
        }
    }
    var ret = &JSONObject{}
    // car && community的查询参数中可能含unknown-id，而后面要恢复到unknown里面去
    var carVOS []*CheVO
    if cap(carIds) == 0 {
        carVOS = make([]*CheVO, 0)
    } else {
        carVOS = make([]*CheVO, len(carIds))
    }
    var communityVOS []*CommunityVO
    if cap(communityIds) == 0 {
        communityVOS = make([]*CommunityVO, 0)
    } else {
        communityVOS = make([]*CommunityVO, len(communityIds))
    }
    // Start 开始查询
    for i := 0; i < 4; i++ {
        if i == 0 && cap(areaIds) != 0 {
            var areaSize = len(areaIds)
            if areaSize > 0 {
                var areaVOS = make([]*AreaVO, areaSize)
                var idChecker = impl.newIntIdChecker(AREA, areaIds);
                var validAreaVOS = impl.AreaServ.BatchQuery(
                    idChecker.validIdList, false,
                    parentDepth, childrenDepth, onlyId)
                var vl = len(validAreaVOS)
                for k := 0; k < vl; k++ {
                    areaVOS[idChecker.validIdIndexList[k]] = validAreaVOS[k]
                }
                for _, k := range idChecker.invalidIdIndexList {
                    areaVOS[k] = nil
                }
                var areas = make([]*JSONObject, len(areaVOS))
                for i, areaVO := range areaVOS {
                    if areaVO != nil {
                        areas[i] = areaVO.ToFlatVO()
                    }
                }
                ret.Put(AREA_NAME, areas)
            } else {
                ret.Put(AREA_NAME, EMPTY_ARRAY_JSON);
            }
        } else if i == 1 && cap(carIds) != 0 {
            var l = len(carIds)
            if l != 0 {
                var idChecker = impl.newIntIdChecker(CAR, carIds);
                var validCheVOS = impl.CheServ.BatchQuery(
                    idChecker.validIdList, false,
                    parentDepth, childrenDepth, onlyId)
                var vl = len(validCheVOS)
                for k := 0; k < vl; k++ {
                    carVOS[idChecker.validIdIndexList[k]] = validCheVOS[k]
                }
                for _, k := range idChecker.invalidIdIndexList {
                    carVOS[k] = nil
                }
            }
            if carSize > 0 {
                var cars = make([]*JSONObject, carSize)
                for i := 0; i < carSize; i++ {
                    if carVOS[i] != nil {
                        cars[i] = carVOS[i].ToFlatVO()
                    }
                }
            } else if carSize == 0 {
                ret.Put(CAR_NAME, EMPTY_ARRAY_JSON)
            }
        } else if i == 2 && cap(categoryIds) != 0 {
            var categorySize = len(categoryIds)
            if categorySize > 0 {
                var categoryVOS = make([]*CategoryVO, categorySize)
                var idChecker = impl.newStringIdChecker(CATEGORY, categoryIds)
                var validCategoryVOS = impl.CategoryServ.BatchQuery(
                    idChecker.validIdList, false, parentDepth, childrenDepth, onlyId)
                var vl = len(validCategoryVOS)
                for k := 0; k < vl; k++ {
                    categoryVOS[idChecker.validIdIndexList[k]] = validCategoryVOS[k]
                }
                for _, k := range idChecker.invalidIdIndexList {
                    categoryVOS[k] = nil
                }
                var categories = make([]*JSONObject, categorySize)
                for i, cate := range categoryVOS {
                    if cate != nil {
                        categories[i] = cate.ToFlatVO()
                    }
                }
                ret.Put(CATEGORY_NAME, categories)
            } else {
                ret.Put(CATEGORY_NAME, EMPTY_ARRAY_JSON)
            }
        } else if i == 3 && cap(communityIds) != 0 {
            if len(communityIds) != 0 {
                var idChecker = impl.newIntIdChecker(COMMUNITY, communityIds);
                var validCommunityVOS = impl.CommunityServ.BatchQuery(
                    idChecker.validIdList, false, 0, 0, onlyId)
                var vl = len(validCommunityVOS)
                for k := 0; k < vl; k++ {
                    communityVOS[idChecker.validIdIndexList[k]] = validCommunityVOS[k]
                }
                for _, k := range idChecker.invalidIdIndexList {
                    communityVOS[k] = nil
                }
            }
            if communitySize > 0 {
                var communities = make([]*JSONObject, communitySize)
                for i := 0; i < communitySize; i++ {
                    if communityVOS[i] != nil {
                        communities[i] = communityVOS[i].ToFlatVO()
                    }

                }
                ret.Put(COMMUNITY_NAME, communities)
            } else if communitySize == 0 {
                ret.Put(COMMUNITY_NAME, EMPTY_ARRAY_JSON)
            }
        }
    }
    // Start query unknownIds
    if unknownSize == 0 {
        ret.Put(UNKNOWN_NAME, EMPTY_ARRAY_JSON)
    } else if unknownSize > 0 && cap(unknownIds) != 0 {
        var jsonArray = NewJSONArray()
        var stillUnknownIndexList = make([]int, 0)
        var stillUnknownIdList = make([]STRING, 0)
        for i := 0; i < unknownSize; i++ {
            if indexMarker[i][0] == CAR.Ordinal() {
                var vo = carVOS[indexMarker[i][1]]
                if vo != nil {
                    jsonArray.Set(i, vo.ToFlatVO())
                }
            } else if indexMarker[i][0] == COMMUNITY.Ordinal() {
                var vo = communityVOS[indexMarker[i][1]]
                if vo != nil {

                }
                jsonArray.Set(i, vo.ToFlatVO())
            } else {
                stillUnknownIndexList = append(stillUnknownIndexList, i)
                stillUnknownIdList = append(stillUnknownIdList, unknownIds[i])
            }
        }
        // 未知类型通过 haojing 接口获取
        if len(stillUnknownIdList) != 0 {
            var stillUnknownVOList = impl.multiGetUnknown(stillUnknownIdList, parentDepth, childrenDepth)
            var l = len(stillUnknownIndexList)
            for i := 0; i < l; i++ {
                jsonArray.Set(stillUnknownIndexList[i], stillUnknownVOList.Get(i))
            }
        }
        ret.Put(UNKNOWN_NAME, jsonArray.Values())
    }
    return ret
}

func (impl *CommonServiceImpl) multiGetUnknown(ids []STRING, p, c int) *JSONArray {
    var ret = NewJSONArrayLimited(len(ids))
    for i, _ := range ids {
        ret.Set(i, nil)
    }
    return ret
}
