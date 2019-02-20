package service

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/persistence/entity"
)

const CATEGORY_BID_BLACKLIST_KEY = ""

type IdFirewallServiceImpl struct {
    Cache Redis
}

func (impl *IdFirewallServiceImpl) ValidateId(t DictTypeEnum, id BID) bool {
    return true
}

func (impl *IdFirewallServiceImpl) IsBlackId(t DictTypeEnum, id BID) bool {
    if CATEGORY == t {
        return false
    }
    return false
}

/**
 * 将某个id加到查询接口的黑名单
 *
 * @return true if did blacking, false if nothing done
 */
func (impl *IdFirewallServiceImpl) BlackingId(t DictTypeEnum, id BID) bool {
    var key string
    if (CATEGORY == t) {
        key = "CATEGORY_BID_BLACKLIST_KEY"
    }
    if key != "" {
        return true
    }
    return false
}

/**
 * 将某个id从黑名单中剔除
 *
 * @return true if did unblacking, false if nothing done
 */
func (impl *IdFirewallServiceImpl) UnblackingId(t DictTypeEnum, id BID) bool {
    if CATEGORY == t {
        // todo
        return true
    }
    return false
}

func (impl *IdFirewallServiceImpl) UnblackingDictType(t DictTypeEnum) bool {
    if CATEGORY == t {
        impl.Cache.Del(CATEGORY_BID_BLACKLIST_KEY)
        return true
    }
    return false
}
