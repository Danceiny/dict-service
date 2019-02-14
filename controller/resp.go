package controller

type CommonResp struct {
    Code uint8       `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func Success(data interface{}) *CommonResp {
    return &CommonResp{0, "", data}
}

func Error(msg string, data interface{}) *CommonResp {
    return &CommonResp{1, msg, data}
}
