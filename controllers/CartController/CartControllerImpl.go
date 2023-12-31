package CartController

import (
	"fmt"
	"net/http"
	"shop/models"
	. "shop/services/CartService"
	"strconv"

	"github.com/gin-gonic/gin"
	"shop/utils/token"
)

type CartControllerImpl struct {
	CartService CartService
}

func NewCartController(CartService CartService) CartController {
	return &CartControllerImpl{
		CartService: CartService,
	}
}

func (c *CartControllerImpl) Save(ctx *gin.Context){
	userId, err := token.GetUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	var Cart models.AddCart

	if err := ctx.ShouldBindJSON(&Cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}
	
	cart, err := c.CartService.Save(userId, &Cart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Status" : true,
		"data"	 : cart,
	})
}

func (c *CartControllerImpl) FindAll(ctx *gin.Context) {

	userId, err := token.GetUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	categories, err := c.CartService.FindAll(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data"	 : categories,
	})
}

func (c *CartControllerImpl) FindById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	category, err := c.CartService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status" : false,
			"message": "data with id " + strconv.Itoa(id) + " not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"data" 	 : category,
	})
}

func (c *CartControllerImpl) Update(ctx *gin.Context) {
	var Cart models.Cart

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&Cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status"  : false,
			"message" : err,
		})
		return
	}

	count, err := c.CartService.Update(id, &Cart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}

	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"message": message,
	})
}

func (c *CartControllerImpl) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status" : false,
			"message": "Invalid request",
		})
		return
	}

	userId, err := token.GetUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err,
		})
		return
	}

	count, err := c.CartService.Delete(userId, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status" : false,
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : true,
		"message": message,
	})	
}