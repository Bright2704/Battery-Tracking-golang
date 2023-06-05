package controller

import (
	"api/Battery-Tracking-go/entity"
	"api/Battery-Tracking-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RelatedNumberController struct {
	RelatedNumberService service.RelatedNumberService
}

func New(relatedNumberservice service.RelatedNumberService) RelatedNumberController {
	return RelatedNumberController{
		RelatedNumberService: relatedNumberservice,
	}
}

func (uc *RelatedNumberController) CreateRelatedNumber(ctx *gin.Context) {
	var user entity.RelatedNumber
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.RelatedNumberService.CreateRelatedNumber(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *RelatedNumberController) GetRelatedNumber(ctx *gin.Context) {
	var username string = ctx.Param("name")
	RelatedNumber, err := uc.RelatedNumberService.GetRelatedNumber(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, RelatedNumber)
}

func (uc *RelatedNumberController) GetAll(ctx *gin.Context) {
	users, err := uc.RelatedNumberService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *RelatedNumberController) UpdateRelatedNumber(ctx *gin.Context) {
	var user entity.RelatedNumber
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}
	err := uc.RelatedNumberService.UpdateRelatedNumber(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *RelatedNumberController) DeletedRelatedNumber(ctx *gin.Context) {
	var username string = ctx.Param("name")
	err := uc.RelatedNumberService.DeleteRelatedNumber(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *RelatedNumberController) RegisterUserRouts(rg *gin.RouterGroup) {
	userroute := rg.Group("/RelatedNumber")
	userroute.POST("/create", uc.CreateRelatedNumber)
	userroute.GET("/get/:name", uc. GetRelatedNumber)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateRelatedNumber)
	userroute.DELETE("/delete/:name", uc.DeletedRelatedNumber)
}
