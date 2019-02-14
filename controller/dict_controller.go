package controller

import (
    "github.com/Danceiny/dict-service/service"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "strconv"
)

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

var areaService = service.AreaServiceImpl{}

func (handler *DictHandler) GetArea(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Warning("id type is not int")
    }
    areaVO := areaService.GetArea(id)
    if areaVO == nil {
        c.JSON(404, Error("NotFound", nil))
    } else {
        c.JSON(200, Success(areaVO.ToFlatVO()))
    }
}

//
// func (handler *DictController) GetCategory() {
//
// }
