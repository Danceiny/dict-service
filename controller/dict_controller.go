package controller

import (
    "github.com/Danceiny/dict-service/api"
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/service"
    . "github.com/Danceiny/go.utils"
    "github.com/gin-gonic/gin"
    "net/http"
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

func (handler *DictHandler) respondToFlat(c *gin.Context, vo FlatVO) {
    if InterfaceHasNilValue(vo) {
        c.JSON(404, Error("Not Found", nil))
    } else {
        c.JSON(200, Success(vo.ToFlatVO()))
    }
}
func (handler *DictHandler) respond(c *gin.Context, vo interface{}) {
    c.JSON(200, Success(vo))
}
func (handler *DictHandler) CommonGet(c *gin.Context) {
    // 处理参数
    var req api.MultiGetQueryReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // 调用并响应
    handler.respond(c, CommonServiceImplCpt.MultiGet(&req))
}
func (handler *DictHandler) GetArea(c *gin.Context) {
    // 处理参数
    var id, parentDepth, childrenDepth int
    var err error
    id, err = strconv.Atoi(c.Param("id"))
    if err != nil {
        ThrowArgException("id is not integer")
    }
    pstr := c.Query("parentDepth")
    if pstr != "" {
        parentDepth, err = strconv.Atoi(pstr)
        if err != nil {
            ThrowArgException("parentDepth is not integer")
        }
    }
    cstr := c.Query("childrenDepth")
    if cstr != "" {
        childrenDepth, err = strconv.Atoi(cstr)
        if err != nil {
            ThrowArgException("parentDepth is not integer")
        }
    }
    // 调用并响应
    handler.respondToFlat(c, AreaServiceImplCpt.GetArea(INT(id), parentDepth, childrenDepth))
}

//
// func (handler *DictController) GetCategory() {
//
// }
