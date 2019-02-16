package controller

import (
    "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/service"
    . "github.com/Danceiny/go.utils"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "strconv"
)

var (
    DictControllerCpt *DictHandler
)

func init() {
    DictControllerCpt = &DictHandler{}
}

type AreaController interface {
    GetArea(c *gin.Context)
    UpdateArea(c *gin.Context)
    AddArea(c *gin.Context)
    DeleteArea(c *gin.Context)
    QueryArea(c *gin.Context)
}

type CategoryController interface {
    GetCategory(c *gin.Context)
    UpdateCategory(c *gin.Context)
    AddCategory(c *gin.Context)
    DeleteCategory(c *gin.Context)
    QueryCategory(c *gin.Context)
}

type DictController interface {
    AreaController
    CategoryController
}

type DictHandler struct {
}

func (handler *DictHandler) respond(c *gin.Context, vo FlatVO) {
    if InterfaceHasNilValue(vo) {
        c.JSON(404, Error("Not Found", nil))
    } else {
        c.JSON(200, Success(vo.ToFlatVO()))
    }
}
func (handler *DictHandler) GetArea(c *gin.Context) {
    // 处理参数
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Warning("id is not integer")
    }
    // 调用并响应
    handler.respond(c, AreaServiceImplCpt.GetArea(common.NodeId(id), 0, 0))
}

//
// func (handler *DictController) GetCategory() {
//
// }
