package service

import (
    . "github.com/Danceiny/go.fastjson"
)

type CategoryVO struct {
}

func (*CategoryVO) ToFlatVO() *JSONObject {
    return nil
}
