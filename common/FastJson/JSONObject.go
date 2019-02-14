package FastJson

import (
    "encoding/json"
    log "github.com/sirupsen/logrus"
)

type JSON interface {
    ToJSON() JsonObject
    Parse(s string)
}

type JsonObject map[string]interface{}

func (jsonObject *JsonObject) Put(k string, v interface{}) {
    (*jsonObject)[k] = v
}
func (jsonObject *JsonObject) PutFluent(k string, v interface{}) *JsonObject {
    (*jsonObject)[k] = v
    return jsonObject
}

func ParseString(s string) JsonObject {
    var o JsonObject
    if err := json.Unmarshal([]byte(s), &o); err != nil {
        log.Warning(err)
        return nil
    } else {
        return o
    }
}

func Parse(bytes []byte) JsonObject {
    var o JsonObject
    if err := json.Unmarshal(bytes, &o); err != nil {
        log.Warning(err)
        return nil
    } else {
        return o
    }
}
func (o *JsonObject) ToJSONString() string {
    return string(o.ToJSON())
}

func (o *JsonObject) ToJSON() []byte {
    if bytes, err := json.Marshal(o); err != nil {
        log.Warning(err)
        return nil
    } else {
        return bytes
    }
}

func ToJSON(v interface{}) []byte {
    if bytes, err := json.Marshal(v); err != nil {
        log.Warning(err)
        return nil
    } else {
        return bytes
    }
}
