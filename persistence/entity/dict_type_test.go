package entity

import (
    . "github.com/Danceiny/dict-service/common"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestDictTypeEnum_ParseBids(t *testing.T) {
    var js = "[52,177416,177415,177414,374,373,372,371,370,369,368,367,334,333,332,331,330,177417]"
    var bids = AREA.ParseBids([]byte(js))
    assert.Equal(t, INT(52), bids[0])
}
